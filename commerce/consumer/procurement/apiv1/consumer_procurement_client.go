// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package procurement

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	procurementpb "cloud.google.com/go/commerce/consumer/procurement/apiv1/procurementpb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newConsumerProcurementClientHook clientHook

// ConsumerProcurementCallOptions contains the retry settings for each method of ConsumerProcurementClient.
type ConsumerProcurementCallOptions struct {
	PlaceOrder   []gax.CallOption
	GetOrder     []gax.CallOption
	ListOrders   []gax.CallOption
	GetOperation []gax.CallOption
}

func defaultConsumerProcurementGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudcommerceconsumerprocurement.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudcommerceconsumerprocurement.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudcommerceconsumerprocurement.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConsumerProcurementCallOptions() *ConsumerProcurementCallOptions {
	return &ConsumerProcurementCallOptions{
		PlaceOrder: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetOrder: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListOrders: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetOperation: []gax.CallOption{},
	}
}

func defaultConsumerProcurementRESTCallOptions() *ConsumerProcurementCallOptions {
	return &ConsumerProcurementCallOptions{
		PlaceOrder: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetOrder: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListOrders: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetOperation: []gax.CallOption{},
	}
}

// internalConsumerProcurementClient is an interface that defines the methods available from Cloud Commerce Consumer Procurement API.
type internalConsumerProcurementClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	PlaceOrder(context.Context, *procurementpb.PlaceOrderRequest, ...gax.CallOption) (*PlaceOrderOperation, error)
	PlaceOrderOperation(name string) *PlaceOrderOperation
	GetOrder(context.Context, *procurementpb.GetOrderRequest, ...gax.CallOption) (*procurementpb.Order, error)
	ListOrders(context.Context, *procurementpb.ListOrdersRequest, ...gax.CallOption) *OrderIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// ConsumerProcurementClient is a client for interacting with Cloud Commerce Consumer Procurement API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// ConsumerProcurementService allows customers to make purchases of products
// served by the Cloud Commerce platform.
//
// When purchases are made, the
// ConsumerProcurementService
// programs the appropriate backends, including both Google’s own
// infrastructure, as well as third-party systems, and to enable billing setup
// for charging for the procured item.
type ConsumerProcurementClient struct {
	// The internal transport-dependent client.
	internalClient internalConsumerProcurementClient

	// The call options for this service.
	CallOptions *ConsumerProcurementCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConsumerProcurementClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConsumerProcurementClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ConsumerProcurementClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// PlaceOrder creates a new Order.
//
// This API only supports GCP spend-based committed use
// discounts specified by GCP documentation.
//
// The returned long-running operation is in-progress until the backend
// completes the creation of the resource. Once completed, the order is
// in
// OrderState.ORDER_STATE_ACTIVE.
// In case of failure, the order resource will be removed.
func (c *ConsumerProcurementClient) PlaceOrder(ctx context.Context, req *procurementpb.PlaceOrderRequest, opts ...gax.CallOption) (*PlaceOrderOperation, error) {
	return c.internalClient.PlaceOrder(ctx, req, opts...)
}

// PlaceOrderOperation returns a new PlaceOrderOperation from a given name.
// The name must be that of a previously created PlaceOrderOperation, possibly from a different process.
func (c *ConsumerProcurementClient) PlaceOrderOperation(name string) *PlaceOrderOperation {
	return c.internalClient.PlaceOrderOperation(name)
}

// GetOrder returns the requested
// Order resource.
func (c *ConsumerProcurementClient) GetOrder(ctx context.Context, req *procurementpb.GetOrderRequest, opts ...gax.CallOption) (*procurementpb.Order, error) {
	return c.internalClient.GetOrder(ctx, req, opts...)
}

// ListOrders lists Order
// resources that the user has access to, within the scope of the parent
// resource.
func (c *ConsumerProcurementClient) ListOrders(ctx context.Context, req *procurementpb.ListOrdersRequest, opts ...gax.CallOption) *OrderIterator {
	return c.internalClient.ListOrders(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *ConsumerProcurementClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// consumerProcurementGRPCClient is a client for interacting with Cloud Commerce Consumer Procurement API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type consumerProcurementGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ConsumerProcurementClient
	CallOptions **ConsumerProcurementCallOptions

	// The gRPC API client.
	consumerProcurementClient procurementpb.ConsumerProcurementServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewConsumerProcurementClient creates a new consumer procurement service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// ConsumerProcurementService allows customers to make purchases of products
// served by the Cloud Commerce platform.
//
// When purchases are made, the
// ConsumerProcurementService
// programs the appropriate backends, including both Google’s own
// infrastructure, as well as third-party systems, and to enable billing setup
// for charging for the procured item.
func NewConsumerProcurementClient(ctx context.Context, opts ...option.ClientOption) (*ConsumerProcurementClient, error) {
	clientOpts := defaultConsumerProcurementGRPCClientOptions()
	if newConsumerProcurementClientHook != nil {
		hookOpts, err := newConsumerProcurementClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ConsumerProcurementClient{CallOptions: defaultConsumerProcurementCallOptions()}

	c := &consumerProcurementGRPCClient{
		connPool:                  connPool,
		consumerProcurementClient: procurementpb.NewConsumerProcurementServiceClient(connPool),
		CallOptions:               &client.CallOptions,
		operationsClient:          longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *consumerProcurementGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *consumerProcurementGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *consumerProcurementGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type consumerProcurementRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ConsumerProcurementClient
	CallOptions **ConsumerProcurementCallOptions
}

// NewConsumerProcurementRESTClient creates a new consumer procurement service rest client.
//
// ConsumerProcurementService allows customers to make purchases of products
// served by the Cloud Commerce platform.
//
// When purchases are made, the
// ConsumerProcurementService
// programs the appropriate backends, including both Google’s own
// infrastructure, as well as third-party systems, and to enable billing setup
// for charging for the procured item.
func NewConsumerProcurementRESTClient(ctx context.Context, opts ...option.ClientOption) (*ConsumerProcurementClient, error) {
	clientOpts := append(defaultConsumerProcurementRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultConsumerProcurementRESTCallOptions()
	c := &consumerProcurementRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &ConsumerProcurementClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultConsumerProcurementRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudcommerceconsumerprocurement.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudcommerceconsumerprocurement.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudcommerceconsumerprocurement.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *consumerProcurementRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *consumerProcurementRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *consumerProcurementRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *consumerProcurementGRPCClient) PlaceOrder(ctx context.Context, req *procurementpb.PlaceOrderRequest, opts ...gax.CallOption) (*PlaceOrderOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PlaceOrder[0:len((*c.CallOptions).PlaceOrder):len((*c.CallOptions).PlaceOrder)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.consumerProcurementClient.PlaceOrder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &PlaceOrderOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *consumerProcurementGRPCClient) GetOrder(ctx context.Context, req *procurementpb.GetOrderRequest, opts ...gax.CallOption) (*procurementpb.Order, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOrder[0:len((*c.CallOptions).GetOrder):len((*c.CallOptions).GetOrder)], opts...)
	var resp *procurementpb.Order
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.consumerProcurementClient.GetOrder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *consumerProcurementGRPCClient) ListOrders(ctx context.Context, req *procurementpb.ListOrdersRequest, opts ...gax.CallOption) *OrderIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOrders[0:len((*c.CallOptions).ListOrders):len((*c.CallOptions).ListOrders)], opts...)
	it := &OrderIterator{}
	req = proto.Clone(req).(*procurementpb.ListOrdersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*procurementpb.Order, string, error) {
		resp := &procurementpb.ListOrdersResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.consumerProcurementClient.ListOrders(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOrders(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *consumerProcurementGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PlaceOrder creates a new Order.
//
// This API only supports GCP spend-based committed use
// discounts specified by GCP documentation.
//
// The returned long-running operation is in-progress until the backend
// completes the creation of the resource. Once completed, the order is
// in
// OrderState.ORDER_STATE_ACTIVE.
// In case of failure, the order resource will be removed.
func (c *consumerProcurementRESTClient) PlaceOrder(ctx context.Context, req *procurementpb.PlaceOrderRequest, opts ...gax.CallOption) (*PlaceOrderOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/orders:place", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &PlaceOrderOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetOrder returns the requested
// Order resource.
func (c *consumerProcurementRESTClient) GetOrder(ctx context.Context, req *procurementpb.GetOrderRequest, opts ...gax.CallOption) (*procurementpb.Order, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOrder[0:len((*c.CallOptions).GetOrder):len((*c.CallOptions).GetOrder)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &procurementpb.Order{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListOrders lists Order
// resources that the user has access to, within the scope of the parent
// resource.
func (c *consumerProcurementRESTClient) ListOrders(ctx context.Context, req *procurementpb.ListOrdersRequest, opts ...gax.CallOption) *OrderIterator {
	it := &OrderIterator{}
	req = proto.Clone(req).(*procurementpb.ListOrdersRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*procurementpb.Order, string, error) {
		resp := &procurementpb.ListOrdersResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/orders", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetOrders(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *consumerProcurementRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// PlaceOrderOperation returns a new PlaceOrderOperation from a given name.
// The name must be that of a previously created PlaceOrderOperation, possibly from a different process.
func (c *consumerProcurementGRPCClient) PlaceOrderOperation(name string) *PlaceOrderOperation {
	return &PlaceOrderOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// PlaceOrderOperation returns a new PlaceOrderOperation from a given name.
// The name must be that of a previously created PlaceOrderOperation, possibly from a different process.
func (c *consumerProcurementRESTClient) PlaceOrderOperation(name string) *PlaceOrderOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &PlaceOrderOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
