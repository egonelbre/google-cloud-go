// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/logging/v2/log_entry.proto

package loggingpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	_type "google.golang.org/genproto/googleapis/logging/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An individual entry in a log.
type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the log to which this log entry belongs:
	//
	//	"projects/[PROJECT_ID]/logs/[LOG_ID]"
	//	"organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//	"billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//	"folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	// A project number may be used in place of PROJECT_ID. The project number is
	// translated to its corresponding PROJECT_ID internally and the `log_name`
	// field will contain PROJECT_ID in queries and exports.
	//
	// `[LOG_ID]` must be URL-encoded within `log_name`. Example:
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	//
	// `[LOG_ID]` must be less than 512 characters long and can only include the
	// following characters: upper and lower case alphanumeric characters,
	// forward-slash, underscore, hyphen, and period.
	//
	// For backward compatibility, if `log_name` begins with a forward-slash, such
	// as `/projects/...`, then the log entry is ingested as usual, but the
	// forward-slash is removed. Listing the log entry will not show the leading
	// slash and filtering for a log name with a leading slash will never return
	// any results.
	LogName string `protobuf:"bytes,12,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// Required. The monitored resource that produced this log entry.
	//
	// Example: a log entry that reports a database error would be associated with
	// the monitored resource designating the particular database that reported
	// the error.
	Resource *monitoredres.MonitoredResource `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	// The log entry payload, which can be one of multiple types.
	//
	// Types that are assignable to Payload:
	//
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_JsonPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
	// Optional. The time the event described by the log entry occurred. This time
	// is used to compute the log entry's age and to enforce the logs retention
	// period. If this field is omitted in a new log entry, then Logging assigns
	// it the current time. Timestamps have nanosecond accuracy, but trailing
	// zeros in the fractional seconds might be omitted when the timestamp is
	// displayed.
	//
	// Incoming log entries must have timestamps that don't exceed the
	// [logs retention
	// period](https://cloud.google.com/logging/quotas#logs_retention_periods) in
	// the past, and that don't exceed 24 hours in the future. Log entries outside
	// those time boundaries aren't ingested by Logging.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Output only. The time the log entry was received by Logging.
	ReceiveTimestamp *timestamppb.Timestamp `protobuf:"bytes,24,opt,name=receive_timestamp,json=receiveTimestamp,proto3" json:"receive_timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity _type.LogSeverity `protobuf:"varint,10,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// Optional. A unique identifier for the log entry. If you provide a value,
	// then Logging considers other log entries in the same project, with the same
	// `timestamp`, and with the same `insert_id` to be duplicates which are
	// removed in a single query result. However, there are no guarantees of
	// de-duplication in the export of logs.
	//
	// If the `insert_id` is omitted when writing a log entry, the Logging API
	// assigns its own unique identifier in this field.
	//
	// In queries, the `insert_id` is also used to order log entries that have
	// the same `log_name` and `timestamp` values.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// Optional. Information about the HTTP request associated with this log
	// entry, if applicable.
	HttpRequest *_type.HttpRequest `protobuf:"bytes,7,opt,name=http_request,json=httpRequest,proto3" json:"http_request,omitempty"`
	// Optional. A map of key, value pairs that provides additional information
	// about the log entry. The labels can be user-defined or system-defined.
	//
	// User-defined labels are arbitrary key, value pairs that you can use to
	// classify logs.
	//
	// System-defined labels are defined by GCP services for platform logs.
	// They have two components - a service namespace component and the
	// attribute name. For example: `compute.googleapis.com/resource_name`.
	//
	// Cloud Logging truncates label keys that exceed 512 B and label
	// values that exceed 64 KB upon their associated log entry being
	// written. The truncation is indicated by an ellipsis at the
	// end of the character string.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Information about an operation associated with the log entry, if
	// applicable.
	Operation *LogEntryOperation `protobuf:"bytes,15,opt,name=operation,proto3" json:"operation,omitempty"`
	// Optional. The REST resource name of the trace being written to
	// [Cloud Trace](https://cloud.google.com/trace) in
	// association with this log entry. For example, if your trace data is stored
	// in the Cloud project "my-trace-project" and if the service that is creating
	// the log entry receives a trace header that includes the trace ID "12345",
	// then the service should use "projects/my-tracing-project/traces/12345".
	//
	// The `trace` field provides the link between logs and traces. By using
	// this field, you can navigate from a log entry to a trace.
	Trace string `protobuf:"bytes,22,opt,name=trace,proto3" json:"trace,omitempty"`
	// Optional. The ID of the [Cloud Trace](https://cloud.google.com/trace) span
	// associated with the current operation in which the log is being written.
	// For example, if a span has the REST resource name of
	// "projects/some-project/traces/some-trace/spans/some-span-id", then the
	// `span_id` field is "some-span-id".
	//
	// A
	// [Span](https://cloud.google.com/trace/docs/reference/v2/rest/v2/projects.traces/batchWrite#Span)
	// represents a single operation within a trace. Whereas a trace may involve
	// multiple different microservices running on multiple different machines,
	// a span generally corresponds to a single logical operation being performed
	// in a single instance of a microservice on one specific machine. Spans
	// are the nodes within the tree that is a trace.
	//
	// Applications that are [instrumented for
	// tracing](https://cloud.google.com/trace/docs/setup) will generally assign a
	// new, unique span ID on each incoming request. It is also common to create
	// and record additional spans corresponding to internal processing elements
	// as well as issuing requests to dependencies.
	//
	// The span ID is expected to be a 16-character, hexadecimal encoding of an
	// 8-byte array and should not be zero. It should be unique within the trace
	// and should, ideally, be generated in a manner that is uniformly random.
	//
	// Example values:
	//
	//   - `000000000000004a`
	//   - `7a2190356c3fc94b`
	//   - `0000f00300090021`
	//   - `d39223e101960076`
	SpanId string `protobuf:"bytes,27,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Optional. The sampling decision of the trace associated with the log entry.
	//
	// True means that the trace resource name in the `trace` field was sampled
	// for storage in a trace backend. False means that the trace was not sampled
	// for storage when this log entry was written, or the sampling decision was
	// unknown at the time. A non-sampled `trace` value is still useful as a
	// request correlation identifier. The default is False.
	TraceSampled bool `protobuf:"varint,30,opt,name=trace_sampled,json=traceSampled,proto3" json:"trace_sampled,omitempty"`
	// Optional. Source code location information associated with the log entry,
	// if any.
	SourceLocation *LogEntrySourceLocation `protobuf:"bytes,23,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	// Optional. Information indicating this LogEntry is part of a sequence of
	// multiple log entries split from a single LogEntry.
	Split *LogSplit `protobuf:"bytes,35,opt,name=split,proto3" json:"split,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_logging_v2_log_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_logging_v2_log_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_google_logging_v2_log_entry_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetLogName() string {
	if x != nil {
		return x.LogName
	}
	return ""
}

func (x *LogEntry) GetResource() *monitoredres.MonitoredResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *LogEntry) GetProtoPayload() *anypb.Any {
	if x, ok := x.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (x *LogEntry) GetTextPayload() string {
	if x, ok := x.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (x *LogEntry) GetJsonPayload() *structpb.Struct {
	if x, ok := x.GetPayload().(*LogEntry_JsonPayload); ok {
		return x.JsonPayload
	}
	return nil
}

func (x *LogEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogEntry) GetReceiveTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceiveTimestamp
	}
	return nil
}

func (x *LogEntry) GetSeverity() _type.LogSeverity {
	if x != nil {
		return x.Severity
	}
	return _type.LogSeverity(0)
}

func (x *LogEntry) GetInsertId() string {
	if x != nil {
		return x.InsertId
	}
	return ""
}

func (x *LogEntry) GetHttpRequest() *_type.HttpRequest {
	if x != nil {
		return x.HttpRequest
	}
	return nil
}

func (x *LogEntry) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LogEntry) GetOperation() *LogEntryOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *LogEntry) GetTrace() string {
	if x != nil {
		return x.Trace
	}
	return ""
}

func (x *LogEntry) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *LogEntry) GetTraceSampled() bool {
	if x != nil {
		return x.TraceSampled
	}
	return false
}

func (x *LogEntry) GetSourceLocation() *LogEntrySourceLocation {
	if x != nil {
		return x.SourceLocation
	}
	return nil
}

func (x *LogEntry) GetSplit() *LogSplit {
	if x != nil {
		return x.Split
	}
	return nil
}

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	// The log entry payload, represented as a protocol buffer. Some Google
	// Cloud Platform services use this field for their log entry payloads.
	//
	// The following protocol buffer types are supported; user-defined types
	// are not supported:
	//
	//	"type.googleapis.com/google.cloud.audit.AuditLog"
	//	"type.googleapis.com/google.appengine.logging.v1.RequestLog"
	ProtoPayload *anypb.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,proto3,oneof"`
}

type LogEntry_TextPayload struct {
	// The log entry payload, represented as a Unicode string (UTF-8).
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

type LogEntry_JsonPayload struct {
	// The log entry payload, represented as a structure that is
	// expressed as a JSON object.
	JsonPayload *structpb.Struct `protobuf:"bytes,6,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}

func (*LogEntry_TextPayload) isLogEntry_Payload() {}

func (*LogEntry_JsonPayload) isLogEntry_Payload() {}

// Additional information about a potentially long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. An arbitrary operation identifier. Log entries with the same
	// identifier are assumed to be part of the same operation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. An arbitrary producer identifier. The combination of `id` and
	// `producer` must be globally unique. Examples for `producer`:
	// `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	Producer string `protobuf:"bytes,2,opt,name=producer,proto3" json:"producer,omitempty"`
	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `protobuf:"varint,3,opt,name=first,proto3" json:"first,omitempty"`
	// Optional. Set this to True if this is the last log entry in the operation.
	Last bool `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *LogEntryOperation) Reset() {
	*x = LogEntryOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_logging_v2_log_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntryOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntryOperation) ProtoMessage() {}

func (x *LogEntryOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_logging_v2_log_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntryOperation.ProtoReflect.Descriptor instead.
func (*LogEntryOperation) Descriptor() ([]byte, []int) {
	return file_google_logging_v2_log_entry_proto_rawDescGZIP(), []int{1}
}

func (x *LogEntryOperation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LogEntryOperation) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *LogEntryOperation) GetFirst() bool {
	if x != nil {
		return x.First
	}
	return false
}

func (x *LogEntryOperation) GetLast() bool {
	if x != nil {
		return x.Last
	}
	return false
}

// Additional information about the source code location that produced the log
// entry.
type LogEntrySourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Source file name. Depending on the runtime environment, this
	// might be a simple name or a fully-qualified name.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Optional. Line within the source file. 1-based; 0 indicates no line number
	// available.
	Line int64 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Optional. Human-readable name of the function or method being invoked, with
	// optional context such as the class or package name. This information may be
	// used in contexts such as the logs viewer, where a file and line number are
	// less meaningful. The format can vary by language. For example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	// (Python).
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
}

func (x *LogEntrySourceLocation) Reset() {
	*x = LogEntrySourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_logging_v2_log_entry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntrySourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntrySourceLocation) ProtoMessage() {}

func (x *LogEntrySourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_logging_v2_log_entry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntrySourceLocation.ProtoReflect.Descriptor instead.
func (*LogEntrySourceLocation) Descriptor() ([]byte, []int) {
	return file_google_logging_v2_log_entry_proto_rawDescGZIP(), []int{2}
}

func (x *LogEntrySourceLocation) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *LogEntrySourceLocation) GetLine() int64 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *LogEntrySourceLocation) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

// Additional information used to correlate multiple log entries. Used when a
// single LogEntry would exceed the Google Cloud Logging size limit and is
// split across multiple log entries.
type LogSplit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A globally unique identifier for all log entries in a sequence of split log
	// entries. All log entries with the same |LogSplit.uid| are assumed to be
	// part of the same sequence of split log entries.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// The index of this LogEntry in the sequence of split log entries. Log
	// entries are given |index| values 0, 1, ..., n-1 for a sequence of n log
	// entries.
	Index int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// The total number of log entries that the original LogEntry was split into.
	TotalSplits int32 `protobuf:"varint,3,opt,name=total_splits,json=totalSplits,proto3" json:"total_splits,omitempty"`
}

func (x *LogSplit) Reset() {
	*x = LogSplit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_logging_v2_log_entry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSplit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSplit) ProtoMessage() {}

func (x *LogSplit) ProtoReflect() protoreflect.Message {
	mi := &file_google_logging_v2_log_entry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSplit.ProtoReflect.Descriptor instead.
func (*LogSplit) Descriptor() ([]byte, []int) {
	return file_google_logging_v2_log_entry_proto_rawDescGZIP(), []int{3}
}

func (x *LogSplit) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *LogSplit) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LogSplit) GetTotalSplits() int32 {
	if x != nil {
		return x.TotalSplits
	}
	return 0
}

var File_google_logging_v2_log_entry_proto protoreflect.FileDescriptor

var file_google_logging_v2_log_entry_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcf, 0x09, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e,
	0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x3c, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3d,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4c, 0x0a,
	0x11, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x41, 0x0a, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x20,
	0x0a, 0x09, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x64,
	0x12, 0x48, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x68,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x47, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x0f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x18, 0x23,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0xbd, 0x01, 0xea, 0x41, 0xb9, 0x01, 0x0a,
	0x1a, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x67, 0x12, 0x1d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x67, 0x7d, 0x12, 0x27, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x67, 0x7d, 0x12, 0x1b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x67, 0x7d,
	0x12, 0x2c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x67, 0x7d, 0x1a, 0x08,
	0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x7d, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x61, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x6c, 0x61,
	0x73, 0x74, 0x22, 0x6b, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1f,
	0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x55, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x42, 0xb3, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x42, 0x0d, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x17, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_logging_v2_log_entry_proto_rawDescOnce sync.Once
	file_google_logging_v2_log_entry_proto_rawDescData = file_google_logging_v2_log_entry_proto_rawDesc
)

func file_google_logging_v2_log_entry_proto_rawDescGZIP() []byte {
	file_google_logging_v2_log_entry_proto_rawDescOnce.Do(func() {
		file_google_logging_v2_log_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_logging_v2_log_entry_proto_rawDescData)
	})
	return file_google_logging_v2_log_entry_proto_rawDescData
}

var file_google_logging_v2_log_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_logging_v2_log_entry_proto_goTypes = []interface{}{
	(*LogEntry)(nil),                       // 0: google.logging.v2.LogEntry
	(*LogEntryOperation)(nil),              // 1: google.logging.v2.LogEntryOperation
	(*LogEntrySourceLocation)(nil),         // 2: google.logging.v2.LogEntrySourceLocation
	(*LogSplit)(nil),                       // 3: google.logging.v2.LogSplit
	nil,                                    // 4: google.logging.v2.LogEntry.LabelsEntry
	(*monitoredres.MonitoredResource)(nil), // 5: google.api.MonitoredResource
	(*anypb.Any)(nil),                      // 6: google.protobuf.Any
	(*structpb.Struct)(nil),                // 7: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),          // 8: google.protobuf.Timestamp
	(_type.LogSeverity)(0),                 // 9: google.logging.type.LogSeverity
	(*_type.HttpRequest)(nil),              // 10: google.logging.type.HttpRequest
}
var file_google_logging_v2_log_entry_proto_depIdxs = []int32{
	5,  // 0: google.logging.v2.LogEntry.resource:type_name -> google.api.MonitoredResource
	6,  // 1: google.logging.v2.LogEntry.proto_payload:type_name -> google.protobuf.Any
	7,  // 2: google.logging.v2.LogEntry.json_payload:type_name -> google.protobuf.Struct
	8,  // 3: google.logging.v2.LogEntry.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 4: google.logging.v2.LogEntry.receive_timestamp:type_name -> google.protobuf.Timestamp
	9,  // 5: google.logging.v2.LogEntry.severity:type_name -> google.logging.type.LogSeverity
	10, // 6: google.logging.v2.LogEntry.http_request:type_name -> google.logging.type.HttpRequest
	4,  // 7: google.logging.v2.LogEntry.labels:type_name -> google.logging.v2.LogEntry.LabelsEntry
	1,  // 8: google.logging.v2.LogEntry.operation:type_name -> google.logging.v2.LogEntryOperation
	2,  // 9: google.logging.v2.LogEntry.source_location:type_name -> google.logging.v2.LogEntrySourceLocation
	3,  // 10: google.logging.v2.LogEntry.split:type_name -> google.logging.v2.LogSplit
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_logging_v2_log_entry_proto_init() }
func file_google_logging_v2_log_entry_proto_init() {
	if File_google_logging_v2_log_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_logging_v2_log_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_logging_v2_log_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntryOperation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_logging_v2_log_entry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntrySourceLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_logging_v2_log_entry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSplit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_logging_v2_log_entry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_JsonPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_logging_v2_log_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_logging_v2_log_entry_proto_goTypes,
		DependencyIndexes: file_google_logging_v2_log_entry_proto_depIdxs,
		MessageInfos:      file_google_logging_v2_log_entry_proto_msgTypes,
	}.Build()
	File_google_logging_v2_log_entry_proto = out.File
	file_google_logging_v2_log_entry_proto_rawDesc = nil
	file_google_logging_v2_log_entry_proto_goTypes = nil
	file_google_logging_v2_log_entry_proto_depIdxs = nil
}
