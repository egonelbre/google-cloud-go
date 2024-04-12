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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/devtools/artifactregistry/v1/vpcsc_config.proto

package artifactregistrypb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// VPCSCPolicy is the VPC SC policy for project and location.
type VPCSCConfig_VPCSCPolicy int32

const (
	// VPCSC_POLICY_UNSPECIFIED - the VPS SC policy is not defined.
	// When VPS SC policy is not defined - the Service will use the default
	// behavior (VPCSC_DENY).
	VPCSCConfig_VPCSC_POLICY_UNSPECIFIED VPCSCConfig_VPCSCPolicy = 0
	// VPCSC_DENY - repository will block the requests to the Upstreams for the
	// Remote Repositories if the resource is in the perimeter.
	VPCSCConfig_DENY VPCSCConfig_VPCSCPolicy = 1
	// VPCSC_ALLOW - repository will allow the requests to the Upstreams for the
	// Remote Repositories if the resource is in the perimeter.
	VPCSCConfig_ALLOW VPCSCConfig_VPCSCPolicy = 2
)

// Enum value maps for VPCSCConfig_VPCSCPolicy.
var (
	VPCSCConfig_VPCSCPolicy_name = map[int32]string{
		0: "VPCSC_POLICY_UNSPECIFIED",
		1: "DENY",
		2: "ALLOW",
	}
	VPCSCConfig_VPCSCPolicy_value = map[string]int32{
		"VPCSC_POLICY_UNSPECIFIED": 0,
		"DENY":                     1,
		"ALLOW":                    2,
	}
)

func (x VPCSCConfig_VPCSCPolicy) Enum() *VPCSCConfig_VPCSCPolicy {
	p := new(VPCSCConfig_VPCSCPolicy)
	*p = x
	return p
}

func (x VPCSCConfig_VPCSCPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VPCSCConfig_VPCSCPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_artifactregistry_v1_vpcsc_config_proto_enumTypes[0].Descriptor()
}

func (VPCSCConfig_VPCSCPolicy) Type() protoreflect.EnumType {
	return &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_enumTypes[0]
}

func (x VPCSCConfig_VPCSCPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VPCSCConfig_VPCSCPolicy.Descriptor instead.
func (VPCSCConfig_VPCSCPolicy) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescGZIP(), []int{0, 0}
}

// The Artifact Registry VPC SC config that apply to a Project.
type VPCSCConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project's VPC SC Config.
	//
	// Always of the form:
	// projects/{projectID}/locations/{location}/vpcscConfig
	//
	// In update request: never set
	// In response: always set
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The project per location VPC SC policy that defines the VPC SC behavior for
	// the Remote Repository (Allow/Deny).
	VpcscPolicy VPCSCConfig_VPCSCPolicy `protobuf:"varint,2,opt,name=vpcsc_policy,json=vpcscPolicy,proto3,enum=google.devtools.artifactregistry.v1.VPCSCConfig_VPCSCPolicy" json:"vpcsc_policy,omitempty"`
}

func (x *VPCSCConfig) Reset() {
	*x = VPCSCConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VPCSCConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VPCSCConfig) ProtoMessage() {}

func (x *VPCSCConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VPCSCConfig.ProtoReflect.Descriptor instead.
func (*VPCSCConfig) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescGZIP(), []int{0}
}

func (x *VPCSCConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VPCSCConfig) GetVpcscPolicy() VPCSCConfig_VPCSCPolicy {
	if x != nil {
		return x.VpcscPolicy
	}
	return VPCSCConfig_VPCSC_POLICY_UNSPECIFIED
}

// Gets the VPC SC config for a project.
type GetVPCSCConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the VPCSCConfig resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetVPCSCConfigRequest) Reset() {
	*x = GetVPCSCConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVPCSCConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVPCSCConfigRequest) ProtoMessage() {}

func (x *GetVPCSCConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVPCSCConfigRequest.ProtoReflect.Descriptor instead.
func (*GetVPCSCConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescGZIP(), []int{1}
}

func (x *GetVPCSCConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Sets the VPCSC config of the project.
type UpdateVPCSCConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project config.
	VpcscConfig *VPCSCConfig `protobuf:"bytes,1,opt,name=vpcsc_config,json=vpcscConfig,proto3" json:"vpcsc_config,omitempty"`
	// Field mask to support partial updates.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateVPCSCConfigRequest) Reset() {
	*x = UpdateVPCSCConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVPCSCConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVPCSCConfigRequest) ProtoMessage() {}

func (x *UpdateVPCSCConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVPCSCConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateVPCSCConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateVPCSCConfigRequest) GetVpcscConfig() *VPCSCConfig {
	if x != nil {
		return x.VpcscConfig
	}
	return nil
}

func (x *UpdateVPCSCConfigRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_google_devtools_artifactregistry_v1_vpcsc_config_proto protoreflect.FileDescriptor

var file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x70, 0x63, 0x73, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0b,
	0x56, 0x50, 0x43, 0x53, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x5f, 0x0a, 0x0c, 0x76, 0x70, 0x63, 0x73, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x50, 0x43, 0x53,
	0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x50, 0x43, 0x53, 0x43, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x0b, 0x76, 0x70, 0x63, 0x73, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x40, 0x0a, 0x0b, 0x56, 0x50, 0x43, 0x53, 0x43, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x1c, 0x0a, 0x18, 0x56, 0x50, 0x43, 0x53, 0x43, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x10, 0x02, 0x3a, 0x65, 0xea, 0x41, 0x62, 0x0a, 0x2b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x70, 0x63, 0x73, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x33, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x70,
	0x63, 0x73, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x60, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x56, 0x50, 0x43, 0x53, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x70, 0x63, 0x73, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x50, 0x43, 0x53, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x0c, 0x76, 0x70, 0x63, 0x73,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x50, 0x43, 0x53, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0b, 0x76, 0x70, 0x63, 0x73, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0xfb, 0x01, 0x0a, 0x27, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x56, 0x50, 0x43, 0x53, 0x43, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x70, 0x62, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x70, 0x62, 0xaa, 0x02, 0x20, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescOnce sync.Once
	file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescData = file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDesc
)

func file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescGZIP() []byte {
	file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescOnce.Do(func() {
		file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescData)
	})
	return file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDescData
}

var file_google_devtools_artifactregistry_v1_vpcsc_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_devtools_artifactregistry_v1_vpcsc_config_proto_goTypes = []interface{}{
	(VPCSCConfig_VPCSCPolicy)(0),     // 0: google.devtools.artifactregistry.v1.VPCSCConfig.VPCSCPolicy
	(*VPCSCConfig)(nil),              // 1: google.devtools.artifactregistry.v1.VPCSCConfig
	(*GetVPCSCConfigRequest)(nil),    // 2: google.devtools.artifactregistry.v1.GetVPCSCConfigRequest
	(*UpdateVPCSCConfigRequest)(nil), // 3: google.devtools.artifactregistry.v1.UpdateVPCSCConfigRequest
	(*fieldmaskpb.FieldMask)(nil),    // 4: google.protobuf.FieldMask
}
var file_google_devtools_artifactregistry_v1_vpcsc_config_proto_depIdxs = []int32{
	0, // 0: google.devtools.artifactregistry.v1.VPCSCConfig.vpcsc_policy:type_name -> google.devtools.artifactregistry.v1.VPCSCConfig.VPCSCPolicy
	1, // 1: google.devtools.artifactregistry.v1.UpdateVPCSCConfigRequest.vpcsc_config:type_name -> google.devtools.artifactregistry.v1.VPCSCConfig
	4, // 2: google.devtools.artifactregistry.v1.UpdateVPCSCConfigRequest.update_mask:type_name -> google.protobuf.FieldMask
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_devtools_artifactregistry_v1_vpcsc_config_proto_init() }
func file_google_devtools_artifactregistry_v1_vpcsc_config_proto_init() {
	if File_google_devtools_artifactregistry_v1_vpcsc_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VPCSCConfig); i {
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
		file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVPCSCConfigRequest); i {
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
		file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVPCSCConfigRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_artifactregistry_v1_vpcsc_config_proto_goTypes,
		DependencyIndexes: file_google_devtools_artifactregistry_v1_vpcsc_config_proto_depIdxs,
		EnumInfos:         file_google_devtools_artifactregistry_v1_vpcsc_config_proto_enumTypes,
		MessageInfos:      file_google_devtools_artifactregistry_v1_vpcsc_config_proto_msgTypes,
	}.Build()
	File_google_devtools_artifactregistry_v1_vpcsc_config_proto = out.File
	file_google_devtools_artifactregistry_v1_vpcsc_config_proto_rawDesc = nil
	file_google_devtools_artifactregistry_v1_vpcsc_config_proto_goTypes = nil
	file_google_devtools_artifactregistry_v1_vpcsc_config_proto_depIdxs = nil
}
