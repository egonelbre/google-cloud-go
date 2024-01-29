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

package dataflow

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newSnapshotsV1Beta3ClientHook clientHook

// SnapshotsV1Beta3CallOptions contains the retry settings for each method of SnapshotsV1Beta3Client.
type SnapshotsV1Beta3CallOptions struct {
	GetSnapshot    []gax.CallOption
	DeleteSnapshot []gax.CallOption
	ListSnapshots  []gax.CallOption
}

func defaultSnapshotsV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSnapshotsV1Beta3CallOptions() *SnapshotsV1Beta3CallOptions {
	return &SnapshotsV1Beta3CallOptions{
		GetSnapshot: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteSnapshot: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListSnapshots: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultSnapshotsV1Beta3RESTCallOptions() *SnapshotsV1Beta3CallOptions {
	return &SnapshotsV1Beta3CallOptions{
		GetSnapshot: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteSnapshot: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListSnapshots: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalSnapshotsV1Beta3Client is an interface that defines the methods available from Dataflow API.
type internalSnapshotsV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetSnapshot(context.Context, *dataflowpb.GetSnapshotRequest, ...gax.CallOption) (*dataflowpb.Snapshot, error)
	DeleteSnapshot(context.Context, *dataflowpb.DeleteSnapshotRequest, ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error)
	ListSnapshots(context.Context, *dataflowpb.ListSnapshotsRequest, ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error)
}

// SnapshotsV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides methods to manage snapshots of Google Cloud Dataflow jobs.
type SnapshotsV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalSnapshotsV1Beta3Client

	// The call options for this service.
	CallOptions *SnapshotsV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SnapshotsV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SnapshotsV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SnapshotsV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetSnapshot gets information about a snapshot.
func (c *SnapshotsV1Beta3Client) GetSnapshot(ctx context.Context, req *dataflowpb.GetSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	return c.internalClient.GetSnapshot(ctx, req, opts...)
}

// DeleteSnapshot deletes a snapshot.
func (c *SnapshotsV1Beta3Client) DeleteSnapshot(ctx context.Context, req *dataflowpb.DeleteSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error) {
	return c.internalClient.DeleteSnapshot(ctx, req, opts...)
}

// ListSnapshots lists snapshots.
func (c *SnapshotsV1Beta3Client) ListSnapshots(ctx context.Context, req *dataflowpb.ListSnapshotsRequest, opts ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error) {
	return c.internalClient.ListSnapshots(ctx, req, opts...)
}

// snapshotsV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type snapshotsV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing SnapshotsV1Beta3Client
	CallOptions **SnapshotsV1Beta3CallOptions

	// The gRPC API client.
	snapshotsV1Beta3Client dataflowpb.SnapshotsV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewSnapshotsV1Beta3Client creates a new snapshots v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides methods to manage snapshots of Google Cloud Dataflow jobs.
func NewSnapshotsV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*SnapshotsV1Beta3Client, error) {
	clientOpts := defaultSnapshotsV1Beta3GRPCClientOptions()
	if newSnapshotsV1Beta3ClientHook != nil {
		hookOpts, err := newSnapshotsV1Beta3ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SnapshotsV1Beta3Client{CallOptions: defaultSnapshotsV1Beta3CallOptions()}

	c := &snapshotsV1Beta3GRPCClient{
		connPool:               connPool,
		snapshotsV1Beta3Client: dataflowpb.NewSnapshotsV1Beta3Client(connPool),
		CallOptions:            &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *snapshotsV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *snapshotsV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *snapshotsV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type snapshotsV1Beta3RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing SnapshotsV1Beta3Client
	CallOptions **SnapshotsV1Beta3CallOptions
}

// NewSnapshotsV1Beta3RESTClient creates a new snapshots v1 beta3 rest client.
//
// Provides methods to manage snapshots of Google Cloud Dataflow jobs.
func NewSnapshotsV1Beta3RESTClient(ctx context.Context, opts ...option.ClientOption) (*SnapshotsV1Beta3Client, error) {
	clientOpts := append(defaultSnapshotsV1Beta3RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultSnapshotsV1Beta3RESTCallOptions()
	c := &snapshotsV1Beta3RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &SnapshotsV1Beta3Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultSnapshotsV1Beta3RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataflow.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://dataflow.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *snapshotsV1Beta3RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *snapshotsV1Beta3RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *snapshotsV1Beta3RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *snapshotsV1Beta3GRPCClient) GetSnapshot(ctx context.Context, req *dataflowpb.GetSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "snapshot_id", url.QueryEscape(req.GetSnapshotId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSnapshot[0:len((*c.CallOptions).GetSnapshot):len((*c.CallOptions).GetSnapshot)], opts...)
	var resp *dataflowpb.Snapshot
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snapshotsV1Beta3Client.GetSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *snapshotsV1Beta3GRPCClient) DeleteSnapshot(ctx context.Context, req *dataflowpb.DeleteSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "snapshot_id", url.QueryEscape(req.GetSnapshotId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteSnapshot[0:len((*c.CallOptions).DeleteSnapshot):len((*c.CallOptions).DeleteSnapshot)], opts...)
	var resp *dataflowpb.DeleteSnapshotResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snapshotsV1Beta3Client.DeleteSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *snapshotsV1Beta3GRPCClient) ListSnapshots(ctx context.Context, req *dataflowpb.ListSnapshotsRequest, opts ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListSnapshots[0:len((*c.CallOptions).ListSnapshots):len((*c.CallOptions).ListSnapshots)], opts...)
	var resp *dataflowpb.ListSnapshotsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snapshotsV1Beta3Client.ListSnapshots(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSnapshot gets information about a snapshot.
func (c *snapshotsV1Beta3RESTClient) GetSnapshot(ctx context.Context, req *dataflowpb.GetSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/snapshots/%v", req.GetProjectId(), req.GetLocation(), req.GetSnapshotId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "snapshot_id", url.QueryEscape(req.GetSnapshotId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetSnapshot[0:len((*c.CallOptions).GetSnapshot):len((*c.CallOptions).GetSnapshot)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.Snapshot{}
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

// DeleteSnapshot deletes a snapshot.
func (c *snapshotsV1Beta3RESTClient) DeleteSnapshot(ctx context.Context, req *dataflowpb.DeleteSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/snapshots/%v", req.GetProjectId(), req.GetLocation(), req.GetSnapshotId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "snapshot_id", url.QueryEscape(req.GetSnapshotId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).DeleteSnapshot[0:len((*c.CallOptions).DeleteSnapshot):len((*c.CallOptions).DeleteSnapshot)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.DeleteSnapshotResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
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

// ListSnapshots lists snapshots.
func (c *snapshotsV1Beta3RESTClient) ListSnapshots(ctx context.Context, req *dataflowpb.ListSnapshotsRequest, opts ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v/snapshots", req.GetProjectId(), req.GetLocation(), req.GetJobId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ListSnapshots[0:len((*c.CallOptions).ListSnapshots):len((*c.CallOptions).ListSnapshots)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.ListSnapshotsResponse{}
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
