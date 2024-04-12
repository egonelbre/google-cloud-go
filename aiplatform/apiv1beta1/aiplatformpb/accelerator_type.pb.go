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
// source: google/cloud/aiplatform/v1beta1/accelerator_type.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a hardware accelerator type.
type AcceleratorType int32

const (
	// Unspecified accelerator type, which means no accelerator.
	AcceleratorType_ACCELERATOR_TYPE_UNSPECIFIED AcceleratorType = 0
	// Nvidia Tesla K80 GPU.
	AcceleratorType_NVIDIA_TESLA_K80 AcceleratorType = 1
	// Nvidia Tesla P100 GPU.
	AcceleratorType_NVIDIA_TESLA_P100 AcceleratorType = 2
	// Nvidia Tesla V100 GPU.
	AcceleratorType_NVIDIA_TESLA_V100 AcceleratorType = 3
	// Nvidia Tesla P4 GPU.
	AcceleratorType_NVIDIA_TESLA_P4 AcceleratorType = 4
	// Nvidia Tesla T4 GPU.
	AcceleratorType_NVIDIA_TESLA_T4 AcceleratorType = 5
	// Nvidia Tesla A100 GPU.
	AcceleratorType_NVIDIA_TESLA_A100 AcceleratorType = 8
	// Nvidia A100 80GB GPU.
	AcceleratorType_NVIDIA_A100_80GB AcceleratorType = 9
	// Nvidia L4 GPU.
	AcceleratorType_NVIDIA_L4 AcceleratorType = 11
	// Nvidia H100 80Gb GPU.
	AcceleratorType_NVIDIA_H100_80GB AcceleratorType = 13
	// TPU v2.
	AcceleratorType_TPU_V2 AcceleratorType = 6
	// TPU v3.
	AcceleratorType_TPU_V3 AcceleratorType = 7
	// TPU v4.
	AcceleratorType_TPU_V4_POD AcceleratorType = 10
	// TPU v5.
	AcceleratorType_TPU_V5_LITEPOD AcceleratorType = 12
)

// Enum value maps for AcceleratorType.
var (
	AcceleratorType_name = map[int32]string{
		0:  "ACCELERATOR_TYPE_UNSPECIFIED",
		1:  "NVIDIA_TESLA_K80",
		2:  "NVIDIA_TESLA_P100",
		3:  "NVIDIA_TESLA_V100",
		4:  "NVIDIA_TESLA_P4",
		5:  "NVIDIA_TESLA_T4",
		8:  "NVIDIA_TESLA_A100",
		9:  "NVIDIA_A100_80GB",
		11: "NVIDIA_L4",
		13: "NVIDIA_H100_80GB",
		6:  "TPU_V2",
		7:  "TPU_V3",
		10: "TPU_V4_POD",
		12: "TPU_V5_LITEPOD",
	}
	AcceleratorType_value = map[string]int32{
		"ACCELERATOR_TYPE_UNSPECIFIED": 0,
		"NVIDIA_TESLA_K80":             1,
		"NVIDIA_TESLA_P100":            2,
		"NVIDIA_TESLA_V100":            3,
		"NVIDIA_TESLA_P4":              4,
		"NVIDIA_TESLA_T4":              5,
		"NVIDIA_TESLA_A100":            8,
		"NVIDIA_A100_80GB":             9,
		"NVIDIA_L4":                    11,
		"NVIDIA_H100_80GB":             13,
		"TPU_V2":                       6,
		"TPU_V3":                       7,
		"TPU_V4_POD":                   10,
		"TPU_V5_LITEPOD":               12,
	}
)

func (x AcceleratorType) Enum() *AcceleratorType {
	p := new(AcceleratorType)
	*p = x
	return p
}

func (x AcceleratorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcceleratorType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_enumTypes[0].Descriptor()
}

func (AcceleratorType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_enumTypes[0]
}

func (x AcceleratorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcceleratorType.Descriptor instead.
func (AcceleratorType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_aiplatform_v1beta1_accelerator_type_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2a, 0xaf, 0x02, 0x0a, 0x0f, 0x41, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x1c, 0x41, 0x43, 0x43, 0x45, 0x4c, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x45, 0x53, 0x4c, 0x41, 0x5f,
	0x4b, 0x38, 0x30, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x5f,
	0x54, 0x45, 0x53, 0x4c, 0x41, 0x5f, 0x50, 0x31, 0x30, 0x30, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x45, 0x53, 0x4c, 0x41, 0x5f, 0x56, 0x31, 0x30,
	0x30, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x45,
	0x53, 0x4c, 0x41, 0x5f, 0x50, 0x34, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x56, 0x49, 0x44,
	0x49, 0x41, 0x5f, 0x54, 0x45, 0x53, 0x4c, 0x41, 0x5f, 0x54, 0x34, 0x10, 0x05, 0x12, 0x15, 0x0a,
	0x11, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x45, 0x53, 0x4c, 0x41, 0x5f, 0x41, 0x31,
	0x30, 0x30, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x5f, 0x41,
	0x31, 0x30, 0x30, 0x5f, 0x38, 0x30, 0x47, 0x42, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x56,
	0x49, 0x44, 0x49, 0x41, 0x5f, 0x4c, 0x34, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x56, 0x49,
	0x44, 0x49, 0x41, 0x5f, 0x48, 0x31, 0x30, 0x30, 0x5f, 0x38, 0x30, 0x47, 0x42, 0x10, 0x0d, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x50, 0x55, 0x5f, 0x56, 0x32, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x50, 0x55, 0x5f, 0x56, 0x33, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x50, 0x55, 0x5f, 0x56,
	0x34, 0x5f, 0x50, 0x4f, 0x44, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x50, 0x55, 0x5f, 0x56,
	0x35, 0x5f, 0x4c, 0x49, 0x54, 0x45, 0x50, 0x4f, 0x44, 0x10, 0x0c, 0x42, 0xeb, 0x01, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x14, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_goTypes = []interface{}{
	(AcceleratorType)(0), // 0: google.cloud.aiplatform.v1beta1.AcceleratorType
}
var file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_init() }
func file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_accelerator_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_enumTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_accelerator_type_proto = out.File
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_depIdxs = nil
}
