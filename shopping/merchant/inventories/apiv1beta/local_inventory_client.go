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

package inventories

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	inventoriespb "cloud.google.com/go/shopping/merchant/inventories/apiv1beta/inventoriespb"
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

var newLocalInventoryClientHook clientHook

// LocalInventoryCallOptions contains the retry settings for each method of LocalInventoryClient.
type LocalInventoryCallOptions struct {
	ListLocalInventories []gax.CallOption
	InsertLocalInventory []gax.CallOption
	DeleteLocalInventory []gax.CallOption
}

func defaultLocalInventoryGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("merchantapi.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("merchantapi.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultLocalInventoryCallOptions() *LocalInventoryCallOptions {
	return &LocalInventoryCallOptions{
		ListLocalInventories: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		InsertLocalInventory: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteLocalInventory: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultLocalInventoryRESTCallOptions() *LocalInventoryCallOptions {
	return &LocalInventoryCallOptions{
		ListLocalInventories: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		InsertLocalInventory: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteLocalInventory: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalLocalInventoryClient is an interface that defines the methods available from Merchant API.
type internalLocalInventoryClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListLocalInventories(context.Context, *inventoriespb.ListLocalInventoriesRequest, ...gax.CallOption) *LocalInventoryIterator
	InsertLocalInventory(context.Context, *inventoriespb.InsertLocalInventoryRequest, ...gax.CallOption) (*inventoriespb.LocalInventory, error)
	DeleteLocalInventory(context.Context, *inventoriespb.DeleteLocalInventoryRequest, ...gax.CallOption) error
}

// LocalInventoryClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage local inventory for products
type LocalInventoryClient struct {
	// The internal transport-dependent client.
	internalClient internalLocalInventoryClient

	// The call options for this service.
	CallOptions *LocalInventoryCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *LocalInventoryClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *LocalInventoryClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *LocalInventoryClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListLocalInventories lists the LocalInventory resources for the given product in your merchant
// account. The response might contain fewer items than specified by
// pageSize. If pageToken was returned in previous request, it can be used
// to obtain additional results.
//
// LocalInventory resources are listed per product for a given account.
func (c *LocalInventoryClient) ListLocalInventories(ctx context.Context, req *inventoriespb.ListLocalInventoriesRequest, opts ...gax.CallOption) *LocalInventoryIterator {
	return c.internalClient.ListLocalInventories(ctx, req, opts...)
}

// InsertLocalInventory inserts a LocalInventory resource to a product in your merchant
// account.
//
// Replaces the full LocalInventory resource if an entry with the same
// [storeCode][google.shopping.merchant.inventories.v1beta.LocalInventory.store_code]
// already exists for the product.
//
// It might take up to 30 minutes for the new or updated LocalInventory
// resource to appear in products.
func (c *LocalInventoryClient) InsertLocalInventory(ctx context.Context, req *inventoriespb.InsertLocalInventoryRequest, opts ...gax.CallOption) (*inventoriespb.LocalInventory, error) {
	return c.internalClient.InsertLocalInventory(ctx, req, opts...)
}

// DeleteLocalInventory deletes the specified LocalInventory from the given product in your
// merchant account. It might take a up to an hour for the
// LocalInventory to be deleted from the specific product.
// Once you have received a successful delete response, wait for that
// period before attempting a delete again.
func (c *LocalInventoryClient) DeleteLocalInventory(ctx context.Context, req *inventoriespb.DeleteLocalInventoryRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteLocalInventory(ctx, req, opts...)
}

// localInventoryGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type localInventoryGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing LocalInventoryClient
	CallOptions **LocalInventoryCallOptions

	// The gRPC API client.
	localInventoryClient inventoriespb.LocalInventoryServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewLocalInventoryClient creates a new local inventory service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage local inventory for products
func NewLocalInventoryClient(ctx context.Context, opts ...option.ClientOption) (*LocalInventoryClient, error) {
	clientOpts := defaultLocalInventoryGRPCClientOptions()
	if newLocalInventoryClientHook != nil {
		hookOpts, err := newLocalInventoryClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := LocalInventoryClient{CallOptions: defaultLocalInventoryCallOptions()}

	c := &localInventoryGRPCClient{
		connPool:             connPool,
		localInventoryClient: inventoriespb.NewLocalInventoryServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *localInventoryGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *localInventoryGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *localInventoryGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type localInventoryRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing LocalInventoryClient
	CallOptions **LocalInventoryCallOptions
}

// NewLocalInventoryRESTClient creates a new local inventory service rest client.
//
// Service to manage local inventory for products
func NewLocalInventoryRESTClient(ctx context.Context, opts ...option.ClientOption) (*LocalInventoryClient, error) {
	clientOpts := append(defaultLocalInventoryRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultLocalInventoryRESTCallOptions()
	c := &localInventoryRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &LocalInventoryClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultLocalInventoryRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://merchantapi.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://merchantapi.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *localInventoryRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *localInventoryRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *localInventoryRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *localInventoryGRPCClient) ListLocalInventories(ctx context.Context, req *inventoriespb.ListLocalInventoriesRequest, opts ...gax.CallOption) *LocalInventoryIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocalInventories[0:len((*c.CallOptions).ListLocalInventories):len((*c.CallOptions).ListLocalInventories)], opts...)
	it := &LocalInventoryIterator{}
	req = proto.Clone(req).(*inventoriespb.ListLocalInventoriesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*inventoriespb.LocalInventory, string, error) {
		resp := &inventoriespb.ListLocalInventoriesResponse{}
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
			resp, err = c.localInventoryClient.ListLocalInventories(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocalInventories(), resp.GetNextPageToken(), nil
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

func (c *localInventoryGRPCClient) InsertLocalInventory(ctx context.Context, req *inventoriespb.InsertLocalInventoryRequest, opts ...gax.CallOption) (*inventoriespb.LocalInventory, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).InsertLocalInventory[0:len((*c.CallOptions).InsertLocalInventory):len((*c.CallOptions).InsertLocalInventory)], opts...)
	var resp *inventoriespb.LocalInventory
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.localInventoryClient.InsertLocalInventory(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *localInventoryGRPCClient) DeleteLocalInventory(ctx context.Context, req *inventoriespb.DeleteLocalInventoryRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteLocalInventory[0:len((*c.CallOptions).DeleteLocalInventory):len((*c.CallOptions).DeleteLocalInventory)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.localInventoryClient.DeleteLocalInventory(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListLocalInventories lists the LocalInventory resources for the given product in your merchant
// account. The response might contain fewer items than specified by
// pageSize. If pageToken was returned in previous request, it can be used
// to obtain additional results.
//
// LocalInventory resources are listed per product for a given account.
func (c *localInventoryRESTClient) ListLocalInventories(ctx context.Context, req *inventoriespb.ListLocalInventoriesRequest, opts ...gax.CallOption) *LocalInventoryIterator {
	it := &LocalInventoryIterator{}
	req = proto.Clone(req).(*inventoriespb.ListLocalInventoriesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*inventoriespb.LocalInventory, string, error) {
		resp := &inventoriespb.ListLocalInventoriesResponse{}
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
		baseUrl.Path += fmt.Sprintf("/inventories/v1beta/%v/localInventories", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
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
		return resp.GetLocalInventories(), resp.GetNextPageToken(), nil
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

// InsertLocalInventory inserts a LocalInventory resource to a product in your merchant
// account.
//
// Replaces the full LocalInventory resource if an entry with the same
// [storeCode][google.shopping.merchant.inventories.v1beta.LocalInventory.store_code]
// already exists for the product.
//
// It might take up to 30 minutes for the new or updated LocalInventory
// resource to appear in products.
func (c *localInventoryRESTClient) InsertLocalInventory(ctx context.Context, req *inventoriespb.InsertLocalInventoryRequest, opts ...gax.CallOption) (*inventoriespb.LocalInventory, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetLocalInventory()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/inventories/v1beta/%v/localInventories:insert", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).InsertLocalInventory[0:len((*c.CallOptions).InsertLocalInventory):len((*c.CallOptions).InsertLocalInventory)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &inventoriespb.LocalInventory{}
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
	return resp, nil
}

// DeleteLocalInventory deletes the specified LocalInventory from the given product in your
// merchant account. It might take a up to an hour for the
// LocalInventory to be deleted from the specific product.
// Once you have received a successful delete response, wait for that
// period before attempting a delete again.
func (c *localInventoryRESTClient) DeleteLocalInventory(ctx context.Context, req *inventoriespb.DeleteLocalInventoryRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/inventories/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}
