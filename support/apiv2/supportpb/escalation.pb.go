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
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/cloud/support/v2/escalation.proto

package supportpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An enum detailing the possible reasons a case may be escalated.
type Escalation_Reason int32

const (
	// The escalation reason is in an unknown state or has not been specified.
	Escalation_REASON_UNSPECIFIED Escalation_Reason = 0
	// The case is taking too long to resolve.
	Escalation_RESOLUTION_TIME Escalation_Reason = 1
	// The support agent does not have the expertise required to successfully
	// resolve the issue.
	Escalation_TECHNICAL_EXPERTISE Escalation_Reason = 2
	// The issue is having a significant business impact.
	Escalation_BUSINESS_IMPACT Escalation_Reason = 3
)

// Enum value maps for Escalation_Reason.
var (
	Escalation_Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "RESOLUTION_TIME",
		2: "TECHNICAL_EXPERTISE",
		3: "BUSINESS_IMPACT",
	}
	Escalation_Reason_value = map[string]int32{
		"REASON_UNSPECIFIED":  0,
		"RESOLUTION_TIME":     1,
		"TECHNICAL_EXPERTISE": 2,
		"BUSINESS_IMPACT":     3,
	}
)

func (x Escalation_Reason) Enum() *Escalation_Reason {
	p := new(Escalation_Reason)
	*p = x
	return p
}

func (x Escalation_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Escalation_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_support_v2_escalation_proto_enumTypes[0].Descriptor()
}

func (Escalation_Reason) Type() protoreflect.EnumType {
	return &file_google_cloud_support_v2_escalation_proto_enumTypes[0]
}

func (x Escalation_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Escalation_Reason.Descriptor instead.
func (Escalation_Reason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_escalation_proto_rawDescGZIP(), []int{0, 0}
}

// An escalation of a support case.
type Escalation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The reason why the Case is being escalated.
	Reason Escalation_Reason `protobuf:"varint,4,opt,name=reason,proto3,enum=google.cloud.support.v2.Escalation_Reason" json:"reason,omitempty"`
	// Required. A free text description to accompany the `reason` field above.
	// Provides additional context on why the case is being escalated.
	Justification string `protobuf:"bytes,5,opt,name=justification,proto3" json:"justification,omitempty"`
}

func (x *Escalation) Reset() {
	*x = Escalation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_support_v2_escalation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Escalation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Escalation) ProtoMessage() {}

func (x *Escalation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2_escalation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Escalation.ProtoReflect.Descriptor instead.
func (*Escalation) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_escalation_proto_rawDescGZIP(), []int{0}
}

func (x *Escalation) GetReason() Escalation_Reason {
	if x != nil {
		return x.Reason
	}
	return Escalation_REASON_UNSPECIFIED
}

func (x *Escalation) GetJustification() string {
	if x != nil {
		return x.Justification
	}
	return ""
}

var File_google_cloud_support_v2_escalation_proto protoreflect.FileDescriptor

var file_google_cloud_support_v2_escalation_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x0a, 0x45, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x73,
	0x63, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0d,
	0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x63, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x53,
	0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x50, 0x45,
	0x52, 0x54, 0x49, 0x53, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x55, 0x53, 0x49, 0x4e,
	0x45, 0x53, 0x53, 0x5f, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x10, 0x03, 0x42, 0xb8, 0x01, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x45, 0x73,
	0x63, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x3b, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_support_v2_escalation_proto_rawDescOnce sync.Once
	file_google_cloud_support_v2_escalation_proto_rawDescData = file_google_cloud_support_v2_escalation_proto_rawDesc
)

func file_google_cloud_support_v2_escalation_proto_rawDescGZIP() []byte {
	file_google_cloud_support_v2_escalation_proto_rawDescOnce.Do(func() {
		file_google_cloud_support_v2_escalation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_support_v2_escalation_proto_rawDescData)
	})
	return file_google_cloud_support_v2_escalation_proto_rawDescData
}

var file_google_cloud_support_v2_escalation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_support_v2_escalation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_support_v2_escalation_proto_goTypes = []interface{}{
	(Escalation_Reason)(0), // 0: google.cloud.support.v2.Escalation.Reason
	(*Escalation)(nil),     // 1: google.cloud.support.v2.Escalation
}
var file_google_cloud_support_v2_escalation_proto_depIdxs = []int32{
	0, // 0: google.cloud.support.v2.Escalation.reason:type_name -> google.cloud.support.v2.Escalation.Reason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_support_v2_escalation_proto_init() }
func file_google_cloud_support_v2_escalation_proto_init() {
	if File_google_cloud_support_v2_escalation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_support_v2_escalation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Escalation); i {
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
			RawDescriptor: file_google_cloud_support_v2_escalation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_support_v2_escalation_proto_goTypes,
		DependencyIndexes: file_google_cloud_support_v2_escalation_proto_depIdxs,
		EnumInfos:         file_google_cloud_support_v2_escalation_proto_enumTypes,
		MessageInfos:      file_google_cloud_support_v2_escalation_proto_msgTypes,
	}.Build()
	File_google_cloud_support_v2_escalation_proto = out.File
	file_google_cloud_support_v2_escalation_proto_rawDesc = nil
	file_google_cloud_support_v2_escalation_proto_goTypes = nil
	file_google_cloud_support_v2_escalation_proto_depIdxs = nil
}
