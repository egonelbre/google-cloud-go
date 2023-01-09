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

package accessapproval_test

import (
	"context"

	accessapproval "cloud.google.com/go/accessapproval/apiv1"
	accessapprovalpb "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	"google.golang.org/api/iterator"
)

func ExampleNewClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListApprovalRequests() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.ListApprovalRequestsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#ListApprovalRequestsMessage.
	}
	it := c.ListApprovalRequests(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetApprovalRequest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.GetApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#GetApprovalRequestMessage.
	}
	resp, err := c.GetApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ApproveApprovalRequest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.ApproveApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#ApproveApprovalRequestMessage.
	}
	resp, err := c.ApproveApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DismissApprovalRequest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.DismissApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#DismissApprovalRequestMessage.
	}
	resp, err := c.DismissApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_InvalidateApprovalRequest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.InvalidateApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#InvalidateApprovalRequestMessage.
	}
	resp, err := c.InvalidateApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetAccessApprovalSettings() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.GetAccessApprovalSettingsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#GetAccessApprovalSettingsMessage.
	}
	resp, err := c.GetAccessApprovalSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateAccessApprovalSettings() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.UpdateAccessApprovalSettingsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#UpdateAccessApprovalSettingsMessage.
	}
	resp, err := c.UpdateAccessApprovalSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteAccessApprovalSettings() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.DeleteAccessApprovalSettingsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#DeleteAccessApprovalSettingsMessage.
	}
	err = c.DeleteAccessApprovalSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetAccessApprovalServiceAccount() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.GetAccessApprovalServiceAccountMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/accessapproval/apiv1/accessapprovalpb#GetAccessApprovalServiceAccountMessage.
	}
	resp, err := c.GetAccessApprovalServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
