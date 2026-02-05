# Spanner Sequence Number Race Condition Fix

## Problem

The Spanner Go client was producing "out-of-order seqno" errors:

```
spanner: code = "InvalidArgument", desc = "Previously received a different request with this seqno. seqno=1"
```

This error occurs when Spanner receives two different SQL statements with the same sequence number for the same transaction.

## Root Cause

The bug was a race condition in the `query()` function in `transaction.go`.

### The Flow

1. `prepareExecuteSQL()` is called which:
   - Gets the transaction selector via `t.acquire(ctx)`
   - Atomically increments `sequenceNumber` and assigns it to `req.Seqno`
   - Sets `req.Transaction` to the current transaction selector

2. `query()` returns a `RowIterator` with a streaming callback

3. When `RowIterator.Next()` is called later, the streaming callback executes:
   - It called `t.getTransactionSelector()` **again**
   - This **overwrote** `req.Transaction` with a potentially different value

### The Race

If the transaction state changed between steps 1 and 3 (e.g., due to an abort/retry), `getTransactionSelector()` would return a `Begin` selector (to start a new transaction) instead of the original active transaction ID.

This meant:
- The request had `Seqno=1` from the original transaction
- But the transaction selector was for a **new** transaction (Begin selector)
- When multiple statements hit this race, Spanner saw multiple requests with `Seqno=1` for what it thought was the same new transaction, and rejected them

### Why the Client-Side Assertion Didn't Catch It

The `lastSentSequenceNumber` assertion tracks sequence numbers per `txReadOnly` instance. When a transaction is retried via `ResetForRetry()`, a **new** `ReadWriteStmtBasedTransaction` is created with its own `txReadOnly` and its own `lastSentSequenceNumber` starting at 0.

The race condition caused the **old** request (with seqno from the old transaction) to be sent with the **new** transaction's Begin selector, but the client-side check couldn't detect this because:
1. Each transaction instance has its own counter
2. The mismatch was between seqno and transaction selector, not between seqnos

## The Fix

In both `query()` and `ReadWithOptions()`:

1. For the **first call** (when `resumeToken` is nil), use the original transaction selector from `prepareExecuteSQL()`/`acquire()`
2. Only call `t.getTransactionSelector()` again for **resume operations** (when `resumeToken` is not nil)

This ensures the sequence number and transaction selector are always consistent - they were both determined at the same point in time during `prepareExecuteSQL()`.

### Code Changes

**query() function:**
```go
// Only update the transaction selector for resume operations.
// For the first call (resumeToken is nil), keep the original transaction selector
// from prepareExecuteSQL to avoid race conditions where the transaction state
// changes between prepareExecuteSQL and this callback execution.
if len(resumeToken) > 0 {
    req.Transaction = t.getTransactionSelector()
}
```

**ReadWithOptions() function:**
```go
// Capture the original transaction selector for use in the streaming callback.
originalTS := ts
// ... in callback:
currentTS := originalTS
if len(resumeToken) > 0 {
    currentTS = t.getTransactionSelector()
}
```

### Sequence Number Ordering Check

A check exists to detect concurrent operations on the same transaction:

```go
// Check for out-of-order sequence numbers before sending.
// Note: Use < instead of <= to allow retries of the same request (when seqno == lastSent).
// The streaming callback can be called multiple times for the same request due to
// internal retry logic in resumableStreamDecoder.
if req.Seqno > 0 && len(resumeToken) == 0 {
    for {
        lastSent := atomic.LoadInt64(&t.lastSentSequenceNumber)
        if req.Seqno < lastSent {
            panic(fmt.Sprintf("spanner: out-of-order sequence number detected in query: ..."))
        }
        if atomic.CompareAndSwapInt64(&t.lastSentSequenceNumber, lastSent, req.Seqno) {
            break
        }
    }
}
```

**Important**: The condition uses `<` instead of `<=` to allow retries of the same request:
- The `resumableStreamDecoder` can call the RPC callback multiple times with `resumeToken = nil` when retrying failed connections (e.g., UNAVAILABLE errors)
- Each retry uses the same `req` object with the same `Seqno`
- Using `<=` would panic on legitimate retries where `seqno == lastSent`
- Using `<` still catches true out-of-order issues where an older request tries to send after a newer one

Note: A "Begin selector with seqno > 0" check was initially added but removed because it's valid for the first statement in a transaction to have seqno=1 with a Begin selector (inline begin transaction).

## Files Changed

- `transaction.go`:
  - `query()` function (~line 720): Keep original `req.Transaction` for first call
  - `ReadWithOptions()` function (~line 358): Capture `originalTS` and use it for first call

## Testing

The fix was verified by:
1. Adding the panic check first, which confirmed the race condition was occurring
2. Applying the fix to preserve the original transaction selector
3. The panic should no longer trigger with the fix in place
