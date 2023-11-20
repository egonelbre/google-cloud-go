// Copyright 2023 Google LLC
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

package workstations

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	workstationspb "cloud.google.com/go/workstations/apiv1beta/workstationspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// CreateWorkstationClusterOperation manages a long-running operation from CreateWorkstationCluster.
type CreateWorkstationClusterOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateWorkstationClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationCluster, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationCluster
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateWorkstationClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationCluster, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationCluster
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateWorkstationClusterOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateWorkstationClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateWorkstationClusterOperation) Name() string {
	return op.lro.Name()
}

// CreateWorkstationConfigOperation manages a long-running operation from CreateWorkstationConfig.
type CreateWorkstationConfigOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateWorkstationConfigOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationConfig, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationConfig
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateWorkstationConfigOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationConfig, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationConfig
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateWorkstationConfigOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateWorkstationConfigOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateWorkstationConfigOperation) Name() string {
	return op.lro.Name()
}

// CreateWorkstationOperation manages a long-running operation from CreateWorkstation.
type CreateWorkstationOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateWorkstationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateWorkstationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateWorkstationOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateWorkstationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateWorkstationOperation) Name() string {
	return op.lro.Name()
}

// DeleteWorkstationClusterOperation manages a long-running operation from DeleteWorkstationCluster.
type DeleteWorkstationClusterOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteWorkstationClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationCluster, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationCluster
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteWorkstationClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationCluster, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationCluster
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteWorkstationClusterOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteWorkstationClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteWorkstationClusterOperation) Name() string {
	return op.lro.Name()
}

// DeleteWorkstationConfigOperation manages a long-running operation from DeleteWorkstationConfig.
type DeleteWorkstationConfigOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteWorkstationConfigOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationConfig, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationConfig
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteWorkstationConfigOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationConfig, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationConfig
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteWorkstationConfigOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteWorkstationConfigOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteWorkstationConfigOperation) Name() string {
	return op.lro.Name()
}

// DeleteWorkstationOperation manages a long-running operation from DeleteWorkstation.
type DeleteWorkstationOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteWorkstationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteWorkstationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteWorkstationOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteWorkstationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteWorkstationOperation) Name() string {
	return op.lro.Name()
}

// StartWorkstationOperation manages a long-running operation from StartWorkstation.
type StartWorkstationOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *StartWorkstationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *StartWorkstationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *StartWorkstationOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *StartWorkstationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *StartWorkstationOperation) Name() string {
	return op.lro.Name()
}

// StopWorkstationOperation manages a long-running operation from StopWorkstation.
type StopWorkstationOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *StopWorkstationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *StopWorkstationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *StopWorkstationOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *StopWorkstationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *StopWorkstationOperation) Name() string {
	return op.lro.Name()
}

// UpdateWorkstationClusterOperation manages a long-running operation from UpdateWorkstationCluster.
type UpdateWorkstationClusterOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateWorkstationClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationCluster, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationCluster
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateWorkstationClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationCluster, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationCluster
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateWorkstationClusterOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateWorkstationClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateWorkstationClusterOperation) Name() string {
	return op.lro.Name()
}

// UpdateWorkstationConfigOperation manages a long-running operation from UpdateWorkstationConfig.
type UpdateWorkstationConfigOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateWorkstationConfigOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationConfig, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationConfig
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateWorkstationConfigOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.WorkstationConfig, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.WorkstationConfig
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateWorkstationConfigOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateWorkstationConfigOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateWorkstationConfigOperation) Name() string {
	return op.lro.Name()
}

// UpdateWorkstationOperation manages a long-running operation from UpdateWorkstation.
type UpdateWorkstationOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateWorkstationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateWorkstationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workstationspb.Workstation, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp workstationspb.Workstation
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateWorkstationOperation) Metadata() (*workstationspb.OperationMetadata, error) {
	var meta workstationspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateWorkstationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateWorkstationOperation) Name() string {
	return op.lro.Name()
}

// OperationIterator manages a stream of *longrunningpb.Operation.
type OperationIterator struct {
	items    []*longrunningpb.Operation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*longrunningpb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*longrunningpb.Operation, error) {
	var item *longrunningpb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkstationClusterIterator manages a stream of *workstationspb.WorkstationCluster.
type WorkstationClusterIterator struct {
	items    []*workstationspb.WorkstationCluster
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*workstationspb.WorkstationCluster, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WorkstationClusterIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkstationClusterIterator) Next() (*workstationspb.WorkstationCluster, error) {
	var item *workstationspb.WorkstationCluster
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkstationClusterIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkstationClusterIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkstationConfigIterator manages a stream of *workstationspb.WorkstationConfig.
type WorkstationConfigIterator struct {
	items    []*workstationspb.WorkstationConfig
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*workstationspb.WorkstationConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WorkstationConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkstationConfigIterator) Next() (*workstationspb.WorkstationConfig, error) {
	var item *workstationspb.WorkstationConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkstationConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkstationConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkstationIterator manages a stream of *workstationspb.Workstation.
type WorkstationIterator struct {
	items    []*workstationspb.Workstation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*workstationspb.Workstation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WorkstationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkstationIterator) Next() (*workstationspb.Workstation, error) {
	var item *workstationspb.Workstation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkstationIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkstationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
