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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/aiplatform/v1/feature_online_store_service.proto

package aiplatformpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Format of the data in the Feature View.
type FeatureViewDataFormat int32

const (
	// Not set. Will be treated as the KeyValue format.
	FeatureViewDataFormat_FEATURE_VIEW_DATA_FORMAT_UNSPECIFIED FeatureViewDataFormat = 0
	// Return response data in key-value format.
	FeatureViewDataFormat_KEY_VALUE FeatureViewDataFormat = 1
	// Return response data in proto Struct format.
	FeatureViewDataFormat_PROTO_STRUCT FeatureViewDataFormat = 2
)

// Enum value maps for FeatureViewDataFormat.
var (
	FeatureViewDataFormat_name = map[int32]string{
		0: "FEATURE_VIEW_DATA_FORMAT_UNSPECIFIED",
		1: "KEY_VALUE",
		2: "PROTO_STRUCT",
	}
	FeatureViewDataFormat_value = map[string]int32{
		"FEATURE_VIEW_DATA_FORMAT_UNSPECIFIED": 0,
		"KEY_VALUE":                            1,
		"PROTO_STRUCT":                         2,
	}
)

func (x FeatureViewDataFormat) Enum() *FeatureViewDataFormat {
	p := new(FeatureViewDataFormat)
	*p = x
	return p
}

func (x FeatureViewDataFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureViewDataFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_enumTypes[0].Descriptor()
}

func (FeatureViewDataFormat) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_enumTypes[0]
}

func (x FeatureViewDataFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureViewDataFormat.Descriptor instead.
func (FeatureViewDataFormat) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP(), []int{0}
}

// Lookup key for a feature view.
type FeatureViewDataKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to KeyOneof:
	//
	//	*FeatureViewDataKey_Key
	KeyOneof isFeatureViewDataKey_KeyOneof `protobuf_oneof:"key_oneof"`
}

func (x *FeatureViewDataKey) Reset() {
	*x = FeatureViewDataKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureViewDataKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureViewDataKey) ProtoMessage() {}

func (x *FeatureViewDataKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureViewDataKey.ProtoReflect.Descriptor instead.
func (*FeatureViewDataKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP(), []int{0}
}

func (m *FeatureViewDataKey) GetKeyOneof() isFeatureViewDataKey_KeyOneof {
	if m != nil {
		return m.KeyOneof
	}
	return nil
}

func (x *FeatureViewDataKey) GetKey() string {
	if x, ok := x.GetKeyOneof().(*FeatureViewDataKey_Key); ok {
		return x.Key
	}
	return ""
}

type isFeatureViewDataKey_KeyOneof interface {
	isFeatureViewDataKey_KeyOneof()
}

type FeatureViewDataKey_Key struct {
	// String key to use for lookup.
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*FeatureViewDataKey_Key) isFeatureViewDataKey_KeyOneof() {}

// Request message for
// [FeatureOnlineStoreService.FetchFeatureValues][google.cloud.aiplatform.v1.FeatureOnlineStoreService.FetchFeatureValues].
// All the features under the requested feature view will be returned.
type FetchFeatureValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. FeatureView resource format
	// `projects/{project}/locations/{location}/featureOnlineStores/{featureOnlineStore}/featureViews/{featureView}`
	FeatureView string `protobuf:"bytes,1,opt,name=feature_view,json=featureView,proto3" json:"feature_view,omitempty"`
	// Optional. The request key to fetch feature values for.
	DataKey *FeatureViewDataKey `protobuf:"bytes,6,opt,name=data_key,json=dataKey,proto3" json:"data_key,omitempty"`
	// Optional. Response data format. If not set,
	// [FeatureViewDataFormat.KEY_VALUE][google.cloud.aiplatform.v1.FeatureViewDataFormat.KEY_VALUE]
	// will be used.
	DataFormat FeatureViewDataFormat `protobuf:"varint,7,opt,name=data_format,json=dataFormat,proto3,enum=google.cloud.aiplatform.v1.FeatureViewDataFormat" json:"data_format,omitempty"`
}

func (x *FetchFeatureValuesRequest) Reset() {
	*x = FetchFeatureValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeatureValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeatureValuesRequest) ProtoMessage() {}

func (x *FetchFeatureValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeatureValuesRequest.ProtoReflect.Descriptor instead.
func (*FetchFeatureValuesRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP(), []int{1}
}

func (x *FetchFeatureValuesRequest) GetFeatureView() string {
	if x != nil {
		return x.FeatureView
	}
	return ""
}

func (x *FetchFeatureValuesRequest) GetDataKey() *FeatureViewDataKey {
	if x != nil {
		return x.DataKey
	}
	return nil
}

func (x *FetchFeatureValuesRequest) GetDataFormat() FeatureViewDataFormat {
	if x != nil {
		return x.DataFormat
	}
	return FeatureViewDataFormat_FEATURE_VIEW_DATA_FORMAT_UNSPECIFIED
}

// Response message for
// [FeatureOnlineStoreService.FetchFeatureValues][google.cloud.aiplatform.v1.FeatureOnlineStoreService.FetchFeatureValues]
type FetchFeatureValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Format:
	//
	//	*FetchFeatureValuesResponse_KeyValues
	//	*FetchFeatureValuesResponse_ProtoStruct
	Format isFetchFeatureValuesResponse_Format `protobuf_oneof:"format"`
}

func (x *FetchFeatureValuesResponse) Reset() {
	*x = FetchFeatureValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeatureValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeatureValuesResponse) ProtoMessage() {}

func (x *FetchFeatureValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeatureValuesResponse.ProtoReflect.Descriptor instead.
func (*FetchFeatureValuesResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP(), []int{2}
}

func (m *FetchFeatureValuesResponse) GetFormat() isFetchFeatureValuesResponse_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (x *FetchFeatureValuesResponse) GetKeyValues() *FetchFeatureValuesResponse_FeatureNameValuePairList {
	if x, ok := x.GetFormat().(*FetchFeatureValuesResponse_KeyValues); ok {
		return x.KeyValues
	}
	return nil
}

func (x *FetchFeatureValuesResponse) GetProtoStruct() *structpb.Struct {
	if x, ok := x.GetFormat().(*FetchFeatureValuesResponse_ProtoStruct); ok {
		return x.ProtoStruct
	}
	return nil
}

type isFetchFeatureValuesResponse_Format interface {
	isFetchFeatureValuesResponse_Format()
}

type FetchFeatureValuesResponse_KeyValues struct {
	// Feature values in KeyValue format.
	KeyValues *FetchFeatureValuesResponse_FeatureNameValuePairList `protobuf:"bytes,3,opt,name=key_values,json=keyValues,proto3,oneof"`
}

type FetchFeatureValuesResponse_ProtoStruct struct {
	// Feature values in proto Struct format.
	ProtoStruct *structpb.Struct `protobuf:"bytes,2,opt,name=proto_struct,json=protoStruct,proto3,oneof"`
}

func (*FetchFeatureValuesResponse_KeyValues) isFetchFeatureValuesResponse_Format() {}

func (*FetchFeatureValuesResponse_ProtoStruct) isFetchFeatureValuesResponse_Format() {}

// Response structure in the format of key (feature name) and (feature) value
// pair.
type FetchFeatureValuesResponse_FeatureNameValuePairList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of feature names and values.
	Features []*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList) Reset() {
	*x = FetchFeatureValuesResponse_FeatureNameValuePairList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeatureValuesResponse_FeatureNameValuePairList) ProtoMessage() {}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeatureValuesResponse_FeatureNameValuePairList.ProtoReflect.Descriptor instead.
func (*FetchFeatureValuesResponse_FeatureNameValuePairList) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList) GetFeatures() []*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair {
	if x != nil {
		return x.Features
	}
	return nil
}

// Feature name & value pair.
type FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Value
	Data isFetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Data `protobuf_oneof:"data"`
	// Feature short name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) Reset() {
	*x = FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) ProtoMessage() {}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair.ProtoReflect.Descriptor instead.
func (*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (m *FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) GetData() isFetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) GetValue() *FeatureValue {
	if x, ok := x.GetData().(*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Value); ok {
		return x.Value
	}
	return nil
}

func (x *FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type isFetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Data interface {
	isFetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Data()
}

type FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Value struct {
	// Feature value.
	Value *FeatureValue `protobuf:"bytes,2,opt,name=value,proto3,oneof"`
}

func (*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Value) isFetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Data() {
}

var File_google_cloud_aiplatform_v1_feature_online_store_service_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x12, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x22, 0x96, 0x02, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x50, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x0b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x4e, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x4b, 0x65, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65,
	0x79, 0x12, 0x57, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xec, 0x03, 0x0a, 0x1a, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0a, 0x6b, 0x65, 0x79,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x09, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x93, 0x02, 0x0a, 0x18, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61,
	0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x74, 0x0a, 0x14, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x08, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2a, 0x62, 0x0a, 0x15, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x28, 0x0a, 0x24, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4b, 0x45, 0x59, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x5f, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x10, 0x02, 0x32, 0xf8, 0x02,
	0x0a, 0x19, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b, 0x02, 0x0a, 0x12,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x85, 0x01, 0xda, 0x41, 0x16, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x2c, 0x20, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x66, 0x3a, 0x01, 0x2a, 0x22, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a,
	0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69,
	0x65, 0x77, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x4d, 0xca, 0x41, 0x19, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xdc, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x1e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62,
	0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescData = file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_feature_online_store_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_aiplatform_v1_feature_online_store_service_proto_goTypes = []interface{}{
	(FeatureViewDataFormat)(0),                                                       // 0: google.cloud.aiplatform.v1.FeatureViewDataFormat
	(*FeatureViewDataKey)(nil),                                                       // 1: google.cloud.aiplatform.v1.FeatureViewDataKey
	(*FetchFeatureValuesRequest)(nil),                                                // 2: google.cloud.aiplatform.v1.FetchFeatureValuesRequest
	(*FetchFeatureValuesResponse)(nil),                                               // 3: google.cloud.aiplatform.v1.FetchFeatureValuesResponse
	(*FetchFeatureValuesResponse_FeatureNameValuePairList)(nil),                      // 4: google.cloud.aiplatform.v1.FetchFeatureValuesResponse.FeatureNameValuePairList
	(*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair)(nil), // 5: google.cloud.aiplatform.v1.FetchFeatureValuesResponse.FeatureNameValuePairList.FeatureNameValuePair
	(*structpb.Struct)(nil),                                                          // 6: google.protobuf.Struct
	(*FeatureValue)(nil),                                                             // 7: google.cloud.aiplatform.v1.FeatureValue
}
var file_google_cloud_aiplatform_v1_feature_online_store_service_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1.FetchFeatureValuesRequest.data_key:type_name -> google.cloud.aiplatform.v1.FeatureViewDataKey
	0, // 1: google.cloud.aiplatform.v1.FetchFeatureValuesRequest.data_format:type_name -> google.cloud.aiplatform.v1.FeatureViewDataFormat
	4, // 2: google.cloud.aiplatform.v1.FetchFeatureValuesResponse.key_values:type_name -> google.cloud.aiplatform.v1.FetchFeatureValuesResponse.FeatureNameValuePairList
	6, // 3: google.cloud.aiplatform.v1.FetchFeatureValuesResponse.proto_struct:type_name -> google.protobuf.Struct
	5, // 4: google.cloud.aiplatform.v1.FetchFeatureValuesResponse.FeatureNameValuePairList.features:type_name -> google.cloud.aiplatform.v1.FetchFeatureValuesResponse.FeatureNameValuePairList.FeatureNameValuePair
	7, // 5: google.cloud.aiplatform.v1.FetchFeatureValuesResponse.FeatureNameValuePairList.FeatureNameValuePair.value:type_name -> google.cloud.aiplatform.v1.FeatureValue
	2, // 6: google.cloud.aiplatform.v1.FeatureOnlineStoreService.FetchFeatureValues:input_type -> google.cloud.aiplatform.v1.FetchFeatureValuesRequest
	3, // 7: google.cloud.aiplatform.v1.FeatureOnlineStoreService.FetchFeatureValues:output_type -> google.cloud.aiplatform.v1.FetchFeatureValuesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_feature_online_store_service_proto_init() }
func file_google_cloud_aiplatform_v1_feature_online_store_service_proto_init() {
	if File_google_cloud_aiplatform_v1_feature_online_store_service_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_featurestore_online_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureViewDataKey); i {
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
		file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeatureValuesRequest); i {
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
		file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeatureValuesResponse); i {
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
		file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeatureValuesResponse_FeatureNameValuePairList); i {
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
		file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair); i {
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
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FeatureViewDataKey_Key)(nil),
	}
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FetchFeatureValuesResponse_KeyValues)(nil),
		(*FetchFeatureValuesResponse_ProtoStruct)(nil),
	}
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*FetchFeatureValuesResponse_FeatureNameValuePairList_FeatureNameValuePair_Value)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_feature_online_store_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_feature_online_store_service_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_feature_online_store_service_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_feature_online_store_service_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_feature_online_store_service_proto = out.File
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_feature_online_store_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeatureOnlineStoreServiceClient is the client API for FeatureOnlineStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeatureOnlineStoreServiceClient interface {
	// Fetch feature values under a FeatureView.
	FetchFeatureValues(ctx context.Context, in *FetchFeatureValuesRequest, opts ...grpc.CallOption) (*FetchFeatureValuesResponse, error)
}

type featureOnlineStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureOnlineStoreServiceClient(cc grpc.ClientConnInterface) FeatureOnlineStoreServiceClient {
	return &featureOnlineStoreServiceClient{cc}
}

func (c *featureOnlineStoreServiceClient) FetchFeatureValues(ctx context.Context, in *FetchFeatureValuesRequest, opts ...grpc.CallOption) (*FetchFeatureValuesResponse, error) {
	out := new(FetchFeatureValuesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.FeatureOnlineStoreService/FetchFeatureValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureOnlineStoreServiceServer is the server API for FeatureOnlineStoreService service.
type FeatureOnlineStoreServiceServer interface {
	// Fetch feature values under a FeatureView.
	FetchFeatureValues(context.Context, *FetchFeatureValuesRequest) (*FetchFeatureValuesResponse, error)
}

// UnimplementedFeatureOnlineStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeatureOnlineStoreServiceServer struct {
}

func (*UnimplementedFeatureOnlineStoreServiceServer) FetchFeatureValues(context.Context, *FetchFeatureValuesRequest) (*FetchFeatureValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchFeatureValues not implemented")
}

func RegisterFeatureOnlineStoreServiceServer(s *grpc.Server, srv FeatureOnlineStoreServiceServer) {
	s.RegisterService(&_FeatureOnlineStoreService_serviceDesc, srv)
}

func _FeatureOnlineStoreService_FetchFeatureValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchFeatureValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureOnlineStoreServiceServer).FetchFeatureValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.FeatureOnlineStoreService/FetchFeatureValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureOnlineStoreServiceServer).FetchFeatureValues(ctx, req.(*FetchFeatureValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeatureOnlineStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1.FeatureOnlineStoreService",
	HandlerType: (*FeatureOnlineStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchFeatureValues",
			Handler:    _FeatureOnlineStoreService_FetchFeatureValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1/feature_online_store_service.proto",
}
