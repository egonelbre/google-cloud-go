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
// source: google/cloud/notebooks/v1/diagnostic_config.proto

package notebookspb

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

// Defines flags that are used to run the diagnostic tool
type DiagnosticConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User Cloud Storage bucket location (REQUIRED).
	// Must be formatted with path prefix (`gs://$GCS_BUCKET`).
	//
	// Permissions:
	// User Managed Notebooks:
	//   - storage.buckets.writer: Must be given to the project's service account
	//     attached to VM.
	//
	// Google Managed Notebooks:
	//   - storage.buckets.writer: Must be given to the project's service account or
	//     user credentials attached to VM depending on authentication mode.
	//
	// Cloud Storage bucket Log file will be written to
	// `gs://$GCS_BUCKET/$RELATIVE_PATH/$VM_DATE_$TIME.tar.gz`
	GcsBucket string `protobuf:"bytes,1,opt,name=gcs_bucket,json=gcsBucket,proto3" json:"gcs_bucket,omitempty"`
	// Optional. Defines the relative storage path in the Cloud Storage bucket
	// where the diagnostic logs will be written: Default path will be the root
	// directory of the Cloud Storage bucket
	// (`gs://$GCS_BUCKET/$DATE_$TIME.tar.gz`)
	// Example of full path where Log file will be written:
	// `gs://$GCS_BUCKET/$RELATIVE_PATH/`
	RelativePath string `protobuf:"bytes,2,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	// Optional. Enables flag to repair service for instance
	RepairFlagEnabled bool `protobuf:"varint,3,opt,name=repair_flag_enabled,json=repairFlagEnabled,proto3" json:"repair_flag_enabled,omitempty"`
	// Optional. Enables flag to capture packets from the instance for 30 seconds
	PacketCaptureFlagEnabled bool `protobuf:"varint,4,opt,name=packet_capture_flag_enabled,json=packetCaptureFlagEnabled,proto3" json:"packet_capture_flag_enabled,omitempty"`
	// Optional. Enables flag to copy all `/home/jupyter` folder contents
	CopyHomeFilesFlagEnabled bool `protobuf:"varint,5,opt,name=copy_home_files_flag_enabled,json=copyHomeFilesFlagEnabled,proto3" json:"copy_home_files_flag_enabled,omitempty"`
}

func (x *DiagnosticConfig) Reset() {
	*x = DiagnosticConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_notebooks_v1_diagnostic_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiagnosticConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiagnosticConfig) ProtoMessage() {}

func (x *DiagnosticConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_notebooks_v1_diagnostic_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiagnosticConfig.ProtoReflect.Descriptor instead.
func (*DiagnosticConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescGZIP(), []int{0}
}

func (x *DiagnosticConfig) GetGcsBucket() string {
	if x != nil {
		return x.GcsBucket
	}
	return ""
}

func (x *DiagnosticConfig) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

func (x *DiagnosticConfig) GetRepairFlagEnabled() bool {
	if x != nil {
		return x.RepairFlagEnabled
	}
	return false
}

func (x *DiagnosticConfig) GetPacketCaptureFlagEnabled() bool {
	if x != nil {
		return x.PacketCaptureFlagEnabled
	}
	return false
}

func (x *DiagnosticConfig) GetCopyHomeFilesFlagEnabled() bool {
	if x != nil {
		return x.CopyHomeFilesFlagEnabled
	}
	return false
}

var File_google_cloud_notebooks_v1_diagnostic_config_proto protoreflect.FileDescriptor

var file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x02, 0x0a, 0x10, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x67,
	0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x33, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x66, 0x6c, 0x61,
	0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x11, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x46, 0x6c, 0x61, 0x67,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x1b, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x18, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x1c, 0x63,
	0x6f, 0x70, 0x79, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x18, 0x63, 0x6f, 0x70, 0x79, 0x48, 0x6f, 0x6d, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x42, 0x75, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x15, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62, 0x3b, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescOnce sync.Once
	file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescData = file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDesc
)

func file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescGZIP() []byte {
	file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescData)
	})
	return file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDescData
}

var file_google_cloud_notebooks_v1_diagnostic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_notebooks_v1_diagnostic_config_proto_goTypes = []interface{}{
	(*DiagnosticConfig)(nil), // 0: google.cloud.notebooks.v1.DiagnosticConfig
}
var file_google_cloud_notebooks_v1_diagnostic_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_notebooks_v1_diagnostic_config_proto_init() }
func file_google_cloud_notebooks_v1_diagnostic_config_proto_init() {
	if File_google_cloud_notebooks_v1_diagnostic_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_notebooks_v1_diagnostic_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiagnosticConfig); i {
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
			RawDescriptor: file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_notebooks_v1_diagnostic_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_notebooks_v1_diagnostic_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_notebooks_v1_diagnostic_config_proto_msgTypes,
	}.Build()
	File_google_cloud_notebooks_v1_diagnostic_config_proto = out.File
	file_google_cloud_notebooks_v1_diagnostic_config_proto_rawDesc = nil
	file_google_cloud_notebooks_v1_diagnostic_config_proto_goTypes = nil
	file_google_cloud_notebooks_v1_diagnostic_config_proto_depIdxs = nil
}
