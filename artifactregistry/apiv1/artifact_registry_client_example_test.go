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

package artifactregistry_test

import (
	"context"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	"google.golang.org/api/iterator"
	artifactregistrypb "google.golang.org/genproto/googleapis/devtools/artifactregistry/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListDockerImages() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListDockerImagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ListDockerImagesRequest.
	}
	it := c.ListDockerImages(ctx, req)
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

func ExampleClient_GetDockerImage() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetDockerImageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetDockerImageRequest.
	}
	resp, err := c.GetDockerImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ImportAptArtifacts() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ImportAptArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ImportAptArtifactsRequest.
	}
	op, err := c.ImportAptArtifacts(ctx, req)
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

func ExampleClient_ImportYumArtifacts() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ImportYumArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ImportYumArtifactsRequest.
	}
	op, err := c.ImportYumArtifacts(ctx, req)
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

func ExampleClient_ListRepositories() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListRepositoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ListRepositoriesRequest.
	}
	it := c.ListRepositories(ctx, req)
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

func ExampleClient_GetRepository() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetRepositoryRequest.
	}
	resp, err := c.GetRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateRepository() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.CreateRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#CreateRepositoryRequest.
	}
	op, err := c.CreateRepository(ctx, req)
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

func ExampleClient_UpdateRepository() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#UpdateRepositoryRequest.
	}
	resp, err := c.UpdateRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteRepository() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#DeleteRepositoryRequest.
	}
	op, err := c.DeleteRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListPackages() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListPackagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ListPackagesRequest.
	}
	it := c.ListPackages(ctx, req)
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

func ExampleClient_GetPackage() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetPackageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetPackageRequest.
	}
	resp, err := c.GetPackage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeletePackage() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeletePackageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#DeletePackageRequest.
	}
	op, err := c.DeletePackage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListVersions() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListVersionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ListVersionsRequest.
	}
	it := c.ListVersions(ctx, req)
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

func ExampleClient_GetVersion() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetVersionRequest.
	}
	resp, err := c.GetVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteVersion() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#DeleteVersionRequest.
	}
	op, err := c.DeleteVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListFiles() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListFilesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ListFilesRequest.
	}
	it := c.ListFiles(ctx, req)
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

func ExampleClient_GetFile() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetFileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetFileRequest.
	}
	resp, err := c.GetFile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListTags() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListTagsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#ListTagsRequest.
	}
	it := c.ListTags(ctx, req)
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

func ExampleClient_GetTag() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetTagRequest.
	}
	resp, err := c.GetTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateTag() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.CreateTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#CreateTagRequest.
	}
	resp, err := c.CreateTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTag() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#UpdateTagRequest.
	}
	resp, err := c.UpdateTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteTag() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#DeleteTagRequest.
	}
	err = c.DeleteTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetProjectSettings() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetProjectSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#GetProjectSettingsRequest.
	}
	resp, err := c.GetProjectSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateProjectSettings() {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateProjectSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/artifactregistry/v1#UpdateProjectSettingsRequest.
	}
	resp, err := c.UpdateProjectSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
