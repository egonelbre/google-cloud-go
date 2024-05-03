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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/apphub/v1/service_project_attachment.proto

package apphubpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServiceProjectAttachment state.
type ServiceProjectAttachment_State int32

const (
	// Unspecified state.
	ServiceProjectAttachment_STATE_UNSPECIFIED ServiceProjectAttachment_State = 0
	// The ServiceProjectAttachment is being created.
	ServiceProjectAttachment_CREATING ServiceProjectAttachment_State = 1
	// The ServiceProjectAttachment is ready.
	// This means Services and Workloads under the corresponding
	// ServiceProjectAttachment is ready for registration.
	ServiceProjectAttachment_ACTIVE ServiceProjectAttachment_State = 2
	// The ServiceProjectAttachment is being deleted.
	ServiceProjectAttachment_DELETING ServiceProjectAttachment_State = 3
)

// Enum value maps for ServiceProjectAttachment_State.
var (
	ServiceProjectAttachment_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
	}
	ServiceProjectAttachment_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"ACTIVE":            2,
		"DELETING":          3,
	}
)

func (x ServiceProjectAttachment_State) Enum() *ServiceProjectAttachment_State {
	p := new(ServiceProjectAttachment_State)
	*p = x
	return p
}

func (x ServiceProjectAttachment_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceProjectAttachment_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apphub_v1_service_project_attachment_proto_enumTypes[0].Descriptor()
}

func (ServiceProjectAttachment_State) Type() protoreflect.EnumType {
	return &file_google_cloud_apphub_v1_service_project_attachment_proto_enumTypes[0]
}

func (x ServiceProjectAttachment_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceProjectAttachment_State.Descriptor instead.
func (ServiceProjectAttachment_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescGZIP(), []int{0, 0}
}

// ServiceProjectAttachment represents an attachment from a service project to a
// host project. Service projects contain the underlying cloud
// infrastructure resources, and expose these resources to the host project
// through a ServiceProjectAttachment. With the attachments, the host project
// can provide an aggregated view of resources across all service projects.
type ServiceProjectAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of a ServiceProjectAttachment. Format:
	// "projects/{host-project-id}/locations/global/serviceProjectAttachments/{service-project-id}."
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Immutable. Service project name in the format: "projects/abc" or
	// "projects/123". As input, project name with either project id or number are
	// accepted. As output, this field will contain project number.
	ServiceProject string `protobuf:"bytes,2,opt,name=service_project,json=serviceProject,proto3" json:"service_project,omitempty"`
	// Output only. Create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. A globally unique identifier (in UUID4 format) for the
	// `ServiceProjectAttachment`.
	Uid string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. ServiceProjectAttachment state.
	State ServiceProjectAttachment_State `protobuf:"varint,5,opt,name=state,proto3,enum=google.cloud.apphub.v1.ServiceProjectAttachment_State" json:"state,omitempty"`
}

func (x *ServiceProjectAttachment) Reset() {
	*x = ServiceProjectAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_service_project_attachment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProjectAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProjectAttachment) ProtoMessage() {}

func (x *ServiceProjectAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_service_project_attachment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceProjectAttachment.ProtoReflect.Descriptor instead.
func (*ServiceProjectAttachment) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceProjectAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceProjectAttachment) GetServiceProject() string {
	if x != nil {
		return x.ServiceProject
	}
	return ""
}

func (x *ServiceProjectAttachment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ServiceProjectAttachment) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ServiceProjectAttachment) GetState() ServiceProjectAttachment_State {
	if x != nil {
		return x.State
	}
	return ServiceProjectAttachment_STATE_UNSPECIFIED
}

var File_google_cloud_apphub_v1_service_project_attachment_proto protoreflect.FileDescriptor

var file_google_cloud_apphub_v1_service_project_attachment_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x04, 0x0a, 0x18,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x5f, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xe2, 0x8c, 0xcf, 0xd7, 0x08, 0x02, 0x08, 0x01, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x51, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x3a, 0xc9,
	0x01, 0xea, 0x41, 0xc5, 0x01, 0x0a, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x2a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x32, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xbf, 0x01, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x1d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x68, 0x75, 0x62, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x70, 0x62, 0xaa, 0x02,
	0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70,
	0x70, 0x48, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x48, 0x75, 0x62, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x41, 0x70, 0x70, 0x48, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescOnce sync.Once
	file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescData = file_google_cloud_apphub_v1_service_project_attachment_proto_rawDesc
)

func file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescGZIP() []byte {
	file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescOnce.Do(func() {
		file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescData)
	})
	return file_google_cloud_apphub_v1_service_project_attachment_proto_rawDescData
}

var file_google_cloud_apphub_v1_service_project_attachment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_apphub_v1_service_project_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_apphub_v1_service_project_attachment_proto_goTypes = []interface{}{
	(ServiceProjectAttachment_State)(0), // 0: google.cloud.apphub.v1.ServiceProjectAttachment.State
	(*ServiceProjectAttachment)(nil),    // 1: google.cloud.apphub.v1.ServiceProjectAttachment
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
}
var file_google_cloud_apphub_v1_service_project_attachment_proto_depIdxs = []int32{
	2, // 0: google.cloud.apphub.v1.ServiceProjectAttachment.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: google.cloud.apphub.v1.ServiceProjectAttachment.state:type_name -> google.cloud.apphub.v1.ServiceProjectAttachment.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_apphub_v1_service_project_attachment_proto_init() }
func file_google_cloud_apphub_v1_service_project_attachment_proto_init() {
	if File_google_cloud_apphub_v1_service_project_attachment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apphub_v1_service_project_attachment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceProjectAttachment); i {
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
			RawDescriptor: file_google_cloud_apphub_v1_service_project_attachment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apphub_v1_service_project_attachment_proto_goTypes,
		DependencyIndexes: file_google_cloud_apphub_v1_service_project_attachment_proto_depIdxs,
		EnumInfos:         file_google_cloud_apphub_v1_service_project_attachment_proto_enumTypes,
		MessageInfos:      file_google_cloud_apphub_v1_service_project_attachment_proto_msgTypes,
	}.Build()
	File_google_cloud_apphub_v1_service_project_attachment_proto = out.File
	file_google_cloud_apphub_v1_service_project_attachment_proto_rawDesc = nil
	file_google_cloud_apphub_v1_service_project_attachment_proto_goTypes = nil
	file_google_cloud_apphub_v1_service_project_attachment_proto_depIdxs = nil
}
