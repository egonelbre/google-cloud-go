// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/cloud/eventarc/v1/google_channel_config.proto

package eventarcpb

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

// A GoogleChannelConfig is a resource that stores the custom settings
// respected by Eventarc first-party triggers in the matching region.
// Once configured, first-party event data will be protected
// using the specified custom managed encryption key instead of Google-managed
// encryption keys.
type GoogleChannelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the config. Must be in the format of,
	// `projects/{project}/locations/{location}/googleChannelConfig`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The last-modified time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	// encrypt/decrypt their event data.
	//
	// It must match the pattern
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName string `protobuf:"bytes,7,opt,name=crypto_key_name,json=cryptoKeyName,proto3" json:"crypto_key_name,omitempty"`
}

func (x *GoogleChannelConfig) Reset() {
	*x = GoogleChannelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_eventarc_v1_google_channel_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleChannelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleChannelConfig) ProtoMessage() {}

func (x *GoogleChannelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_eventarc_v1_google_channel_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleChannelConfig.ProtoReflect.Descriptor instead.
func (*GoogleChannelConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleChannelConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleChannelConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *GoogleChannelConfig) GetCryptoKeyName() string {
	if x != nil {
		return x.CryptoKeyName
	}
	return ""
}

var File_google_cloud_eventarc_v1_google_channel_config_proto protoreflect.FileDescriptor

var file_google_cloud_eventarc_v1_google_channel_config_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02,
	0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x51, 0x0a, 0x0f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x01, 0xfa, 0x41,
	0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x3a, 0x82, 0x01, 0xea, 0x41, 0x7f, 0x0a, 0x2b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x61, 0x72, 0x63, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x32, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xc3, 0x02, 0xea, 0x41, 0x78, 0x0a, 0x21,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65,
	0x79, 0x12, 0x53, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x52, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72,
	0x63, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x70, 0x62, 0x3b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x70, 0x62, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61,
	0x72, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescOnce sync.Once
	file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescData = file_google_cloud_eventarc_v1_google_channel_config_proto_rawDesc
)

func file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescGZIP() []byte {
	file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescData)
	})
	return file_google_cloud_eventarc_v1_google_channel_config_proto_rawDescData
}

var file_google_cloud_eventarc_v1_google_channel_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_eventarc_v1_google_channel_config_proto_goTypes = []interface{}{
	(*GoogleChannelConfig)(nil),   // 0: google.cloud.eventarc.v1.GoogleChannelConfig
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_google_cloud_eventarc_v1_google_channel_config_proto_depIdxs = []int32{
	1, // 0: google.cloud.eventarc.v1.GoogleChannelConfig.update_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_eventarc_v1_google_channel_config_proto_init() }
func file_google_cloud_eventarc_v1_google_channel_config_proto_init() {
	if File_google_cloud_eventarc_v1_google_channel_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_eventarc_v1_google_channel_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleChannelConfig); i {
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
			RawDescriptor: file_google_cloud_eventarc_v1_google_channel_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_eventarc_v1_google_channel_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_eventarc_v1_google_channel_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_eventarc_v1_google_channel_config_proto_msgTypes,
	}.Build()
	File_google_cloud_eventarc_v1_google_channel_config_proto = out.File
	file_google_cloud_eventarc_v1_google_channel_config_proto_rawDesc = nil
	file_google_cloud_eventarc_v1_google_channel_config_proto_goTypes = nil
	file_google_cloud_eventarc_v1_google_channel_config_proto_depIdxs = nil
}
