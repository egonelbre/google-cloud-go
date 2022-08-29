// Copyright 2022 Google LLC
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

package domains_test

import (
	"context"

	domains "cloud.google.com/go/domains/apiv1beta1"
	"google.golang.org/api/iterator"
	domainspb "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRESTClient() {
	ctx := context.Background()
	c, err := domains.NewRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_SearchDomains() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.SearchDomainsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#SearchDomainsRequest.
	}
	resp, err := c.SearchDomains(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RetrieveRegisterParameters() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.RetrieveRegisterParametersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#RetrieveRegisterParametersRequest.
	}
	resp, err := c.RetrieveRegisterParameters(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RegisterDomain() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.RegisterDomainRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#RegisterDomainRequest.
	}
	op, err := c.RegisterDomain(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RetrieveTransferParameters() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.RetrieveTransferParametersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#RetrieveTransferParametersRequest.
	}
	resp, err := c.RetrieveTransferParameters(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_TransferDomain() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.TransferDomainRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#TransferDomainRequest.
	}
	op, err := c.TransferDomain(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListRegistrations() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.ListRegistrationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#ListRegistrationsRequest.
	}
	it := c.ListRegistrations(ctx, req)
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

func ExampleClient_GetRegistration() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.GetRegistrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#GetRegistrationRequest.
	}
	resp, err := c.GetRegistration(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateRegistration() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.UpdateRegistrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#UpdateRegistrationRequest.
	}
	op, err := c.UpdateRegistration(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ConfigureManagementSettings() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.ConfigureManagementSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#ConfigureManagementSettingsRequest.
	}
	op, err := c.ConfigureManagementSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ConfigureDnsSettings() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.ConfigureDnsSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#ConfigureDnsSettingsRequest.
	}
	op, err := c.ConfigureDnsSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ConfigureContactSettings() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.ConfigureContactSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#ConfigureContactSettingsRequest.
	}
	op, err := c.ConfigureContactSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ExportRegistration() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.ExportRegistrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#ExportRegistrationRequest.
	}
	op, err := c.ExportRegistration(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteRegistration() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.DeleteRegistrationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#DeleteRegistrationRequest.
	}
	op, err := c.DeleteRegistration(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_RetrieveAuthorizationCode() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.RetrieveAuthorizationCodeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#RetrieveAuthorizationCodeRequest.
	}
	resp, err := c.RetrieveAuthorizationCode(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ResetAuthorizationCode() {
	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &domainspb.ResetAuthorizationCodeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/domains/v1beta1#ResetAuthorizationCodeRequest.
	}
	resp, err := c.ResetAuthorizationCode(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
