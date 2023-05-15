// Copyright 2022 Google LLC
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
// source: google/cloud/documentai/v1/document_io.proto

package documentaipb

import (
	reflect "reflect"
	sync "sync"

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

// Payload message of raw document content (bytes).
type RawDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Inline document content.
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// An IANA MIME type (RFC6838) indicating the nature and format of the
	// [content][google.cloud.documentai.v1.RawDocument.content].
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *RawDocument) Reset() {
	*x = RawDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawDocument) ProtoMessage() {}

func (x *RawDocument) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawDocument.ProtoReflect.Descriptor instead.
func (*RawDocument) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{0}
}

func (x *RawDocument) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *RawDocument) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

// Specifies a document stored on Cloud Storage.
type GcsDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Cloud Storage object uri.
	GcsUri string `protobuf:"bytes,1,opt,name=gcs_uri,json=gcsUri,proto3" json:"gcs_uri,omitempty"`
	// An IANA MIME type (RFC6838) of the content.
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *GcsDocument) Reset() {
	*x = GcsDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsDocument) ProtoMessage() {}

func (x *GcsDocument) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsDocument.ProtoReflect.Descriptor instead.
func (*GcsDocument) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{1}
}

func (x *GcsDocument) GetGcsUri() string {
	if x != nil {
		return x.GcsUri
	}
	return ""
}

func (x *GcsDocument) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

// Specifies a set of documents on Cloud Storage.
type GcsDocuments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of documents.
	Documents []*GcsDocument `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *GcsDocuments) Reset() {
	*x = GcsDocuments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsDocuments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsDocuments) ProtoMessage() {}

func (x *GcsDocuments) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsDocuments.ProtoReflect.Descriptor instead.
func (*GcsDocuments) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{2}
}

func (x *GcsDocuments) GetDocuments() []*GcsDocument {
	if x != nil {
		return x.Documents
	}
	return nil
}

// Specifies all documents on Cloud Storage with a common prefix.
type GcsPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URI prefix.
	GcsUriPrefix string `protobuf:"bytes,1,opt,name=gcs_uri_prefix,json=gcsUriPrefix,proto3" json:"gcs_uri_prefix,omitempty"`
}

func (x *GcsPrefix) Reset() {
	*x = GcsPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsPrefix) ProtoMessage() {}

func (x *GcsPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsPrefix.ProtoReflect.Descriptor instead.
func (*GcsPrefix) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{3}
}

func (x *GcsPrefix) GetGcsUriPrefix() string {
	if x != nil {
		return x.GcsUriPrefix
	}
	return ""
}

// The common config to specify a set of documents used as input.
type BatchDocumentsInputConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source.
	//
	// Types that are assignable to Source:
	//	*BatchDocumentsInputConfig_GcsPrefix
	//	*BatchDocumentsInputConfig_GcsDocuments
	Source isBatchDocumentsInputConfig_Source `protobuf_oneof:"source"`
}

func (x *BatchDocumentsInputConfig) Reset() {
	*x = BatchDocumentsInputConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDocumentsInputConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDocumentsInputConfig) ProtoMessage() {}

func (x *BatchDocumentsInputConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDocumentsInputConfig.ProtoReflect.Descriptor instead.
func (*BatchDocumentsInputConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{4}
}

func (m *BatchDocumentsInputConfig) GetSource() isBatchDocumentsInputConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *BatchDocumentsInputConfig) GetGcsPrefix() *GcsPrefix {
	if x, ok := x.GetSource().(*BatchDocumentsInputConfig_GcsPrefix); ok {
		return x.GcsPrefix
	}
	return nil
}

func (x *BatchDocumentsInputConfig) GetGcsDocuments() *GcsDocuments {
	if x, ok := x.GetSource().(*BatchDocumentsInputConfig_GcsDocuments); ok {
		return x.GcsDocuments
	}
	return nil
}

type isBatchDocumentsInputConfig_Source interface {
	isBatchDocumentsInputConfig_Source()
}

type BatchDocumentsInputConfig_GcsPrefix struct {
	// The set of documents that match the specified Cloud Storage `gcs_prefix`.
	GcsPrefix *GcsPrefix `protobuf:"bytes,1,opt,name=gcs_prefix,json=gcsPrefix,proto3,oneof"`
}

type BatchDocumentsInputConfig_GcsDocuments struct {
	// The set of documents individually specified on Cloud Storage.
	GcsDocuments *GcsDocuments `protobuf:"bytes,2,opt,name=gcs_documents,json=gcsDocuments,proto3,oneof"`
}

func (*BatchDocumentsInputConfig_GcsPrefix) isBatchDocumentsInputConfig_Source() {}

func (*BatchDocumentsInputConfig_GcsDocuments) isBatchDocumentsInputConfig_Source() {}

// Config that controls the output of documents. All documents will be written
// as a JSON file.
type DocumentOutputConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The destination of the results.
	//
	// Types that are assignable to Destination:
	//	*DocumentOutputConfig_GcsOutputConfig_
	Destination isDocumentOutputConfig_Destination `protobuf_oneof:"destination"`
}

func (x *DocumentOutputConfig) Reset() {
	*x = DocumentOutputConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentOutputConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentOutputConfig) ProtoMessage() {}

func (x *DocumentOutputConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentOutputConfig.ProtoReflect.Descriptor instead.
func (*DocumentOutputConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{5}
}

func (m *DocumentOutputConfig) GetDestination() isDocumentOutputConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *DocumentOutputConfig) GetGcsOutputConfig() *DocumentOutputConfig_GcsOutputConfig {
	if x, ok := x.GetDestination().(*DocumentOutputConfig_GcsOutputConfig_); ok {
		return x.GcsOutputConfig
	}
	return nil
}

type isDocumentOutputConfig_Destination interface {
	isDocumentOutputConfig_Destination()
}

type DocumentOutputConfig_GcsOutputConfig_ struct {
	// Output config to write the results to Cloud Storage.
	GcsOutputConfig *DocumentOutputConfig_GcsOutputConfig `protobuf:"bytes,1,opt,name=gcs_output_config,json=gcsOutputConfig,proto3,oneof"`
}

func (*DocumentOutputConfig_GcsOutputConfig_) isDocumentOutputConfig_Destination() {}

// The configuration used when outputting documents.
type DocumentOutputConfig_GcsOutputConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Cloud Storage uri (a directory) of the output.
	GcsUri string `protobuf:"bytes,1,opt,name=gcs_uri,json=gcsUri,proto3" json:"gcs_uri,omitempty"`
	// Specifies which fields to include in the output documents.
	// Only supports top level document and pages field so it must be in the
	// form of `{document_field_name}` or `pages.{page_field_name}`.
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Specifies the sharding config for the output document.
	ShardingConfig *DocumentOutputConfig_GcsOutputConfig_ShardingConfig `protobuf:"bytes,3,opt,name=sharding_config,json=shardingConfig,proto3" json:"sharding_config,omitempty"`
}

func (x *DocumentOutputConfig_GcsOutputConfig) Reset() {
	*x = DocumentOutputConfig_GcsOutputConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentOutputConfig_GcsOutputConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentOutputConfig_GcsOutputConfig) ProtoMessage() {}

func (x *DocumentOutputConfig_GcsOutputConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentOutputConfig_GcsOutputConfig.ProtoReflect.Descriptor instead.
func (*DocumentOutputConfig_GcsOutputConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{5, 0}
}

func (x *DocumentOutputConfig_GcsOutputConfig) GetGcsUri() string {
	if x != nil {
		return x.GcsUri
	}
	return ""
}

func (x *DocumentOutputConfig_GcsOutputConfig) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *DocumentOutputConfig_GcsOutputConfig) GetShardingConfig() *DocumentOutputConfig_GcsOutputConfig_ShardingConfig {
	if x != nil {
		return x.ShardingConfig
	}
	return nil
}

// The sharding config for the output document.
type DocumentOutputConfig_GcsOutputConfig_ShardingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of pages per shard.
	PagesPerShard int32 `protobuf:"varint,1,opt,name=pages_per_shard,json=pagesPerShard,proto3" json:"pages_per_shard,omitempty"`
	// The number of overlapping pages between consecutive shards.
	PagesOverlap int32 `protobuf:"varint,2,opt,name=pages_overlap,json=pagesOverlap,proto3" json:"pages_overlap,omitempty"`
}

func (x *DocumentOutputConfig_GcsOutputConfig_ShardingConfig) Reset() {
	*x = DocumentOutputConfig_GcsOutputConfig_ShardingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentOutputConfig_GcsOutputConfig_ShardingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentOutputConfig_GcsOutputConfig_ShardingConfig) ProtoMessage() {}

func (x *DocumentOutputConfig_GcsOutputConfig_ShardingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_document_io_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentOutputConfig_GcsOutputConfig_ShardingConfig.ProtoReflect.Descriptor instead.
func (*DocumentOutputConfig_GcsOutputConfig_ShardingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP(), []int{5, 0, 0}
}

func (x *DocumentOutputConfig_GcsOutputConfig_ShardingConfig) GetPagesPerShard() int32 {
	if x != nil {
		return x.PagesPerShard
	}
	return 0
}

func (x *DocumentOutputConfig_GcsOutputConfig_ShardingConfig) GetPagesOverlap() int32 {
	if x != nil {
		return x.PagesOverlap
	}
	return 0
}

var File_google_cloud_documentai_v1_document_io_proto protoreflect.FileDescriptor

var file_google_cloud_documentai_v1_document_io_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0b,
	0x52, 0x61, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x43, 0x0a, 0x0b, 0x47, 0x63, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x63, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x47, 0x63, 0x73, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x31,
	0x0a, 0x09, 0x47, 0x63, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x24, 0x0a, 0x0e, 0x67,
	0x63, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x63, 0x73, 0x55, 0x72, 0x69, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x46, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x63, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x48, 0x00, 0x52, 0x09, 0x67, 0x63,
	0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x4f, 0x0a, 0x0d, 0x67, 0x63, 0x73, 0x5f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x67, 0x63, 0x73, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0xd6, 0x03, 0x0a, 0x14, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6e, 0x0a, 0x11, 0x67,
	0x63, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x63, 0x73, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x63, 0x73, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xbe, 0x02, 0x0a, 0x0f,
	0x47, 0x63, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x17, 0x0a, 0x07, 0x67, 0x63, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x78, 0x0a, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47,
	0x63, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x5d, 0x0a,
	0x0e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x26, 0x0a, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x5f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x70, 0x42, 0x0d, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xcd, 0x01, 0x0a, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0f,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x69, 0x70, 0x62, 0x3b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x70,
	0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_documentai_v1_document_io_proto_rawDescOnce sync.Once
	file_google_cloud_documentai_v1_document_io_proto_rawDescData = file_google_cloud_documentai_v1_document_io_proto_rawDesc
)

func file_google_cloud_documentai_v1_document_io_proto_rawDescGZIP() []byte {
	file_google_cloud_documentai_v1_document_io_proto_rawDescOnce.Do(func() {
		file_google_cloud_documentai_v1_document_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_documentai_v1_document_io_proto_rawDescData)
	})
	return file_google_cloud_documentai_v1_document_io_proto_rawDescData
}

var file_google_cloud_documentai_v1_document_io_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_documentai_v1_document_io_proto_goTypes = []interface{}{
	(*RawDocument)(nil),                                         // 0: google.cloud.documentai.v1.RawDocument
	(*GcsDocument)(nil),                                         // 1: google.cloud.documentai.v1.GcsDocument
	(*GcsDocuments)(nil),                                        // 2: google.cloud.documentai.v1.GcsDocuments
	(*GcsPrefix)(nil),                                           // 3: google.cloud.documentai.v1.GcsPrefix
	(*BatchDocumentsInputConfig)(nil),                           // 4: google.cloud.documentai.v1.BatchDocumentsInputConfig
	(*DocumentOutputConfig)(nil),                                // 5: google.cloud.documentai.v1.DocumentOutputConfig
	(*DocumentOutputConfig_GcsOutputConfig)(nil),                // 6: google.cloud.documentai.v1.DocumentOutputConfig.GcsOutputConfig
	(*DocumentOutputConfig_GcsOutputConfig_ShardingConfig)(nil), // 7: google.cloud.documentai.v1.DocumentOutputConfig.GcsOutputConfig.ShardingConfig
	(*fieldmaskpb.FieldMask)(nil),                               // 8: google.protobuf.FieldMask
}
var file_google_cloud_documentai_v1_document_io_proto_depIdxs = []int32{
	1, // 0: google.cloud.documentai.v1.GcsDocuments.documents:type_name -> google.cloud.documentai.v1.GcsDocument
	3, // 1: google.cloud.documentai.v1.BatchDocumentsInputConfig.gcs_prefix:type_name -> google.cloud.documentai.v1.GcsPrefix
	2, // 2: google.cloud.documentai.v1.BatchDocumentsInputConfig.gcs_documents:type_name -> google.cloud.documentai.v1.GcsDocuments
	6, // 3: google.cloud.documentai.v1.DocumentOutputConfig.gcs_output_config:type_name -> google.cloud.documentai.v1.DocumentOutputConfig.GcsOutputConfig
	8, // 4: google.cloud.documentai.v1.DocumentOutputConfig.GcsOutputConfig.field_mask:type_name -> google.protobuf.FieldMask
	7, // 5: google.cloud.documentai.v1.DocumentOutputConfig.GcsOutputConfig.sharding_config:type_name -> google.cloud.documentai.v1.DocumentOutputConfig.GcsOutputConfig.ShardingConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_documentai_v1_document_io_proto_init() }
func file_google_cloud_documentai_v1_document_io_proto_init() {
	if File_google_cloud_documentai_v1_document_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawDocument); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsDocument); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsDocuments); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsPrefix); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDocumentsInputConfig); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentOutputConfig); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentOutputConfig_GcsOutputConfig); i {
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
		file_google_cloud_documentai_v1_document_io_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentOutputConfig_GcsOutputConfig_ShardingConfig); i {
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
	file_google_cloud_documentai_v1_document_io_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*BatchDocumentsInputConfig_GcsPrefix)(nil),
		(*BatchDocumentsInputConfig_GcsDocuments)(nil),
	}
	file_google_cloud_documentai_v1_document_io_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*DocumentOutputConfig_GcsOutputConfig_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_documentai_v1_document_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_documentai_v1_document_io_proto_goTypes,
		DependencyIndexes: file_google_cloud_documentai_v1_document_io_proto_depIdxs,
		MessageInfos:      file_google_cloud_documentai_v1_document_io_proto_msgTypes,
	}.Build()
	File_google_cloud_documentai_v1_document_io_proto = out.File
	file_google_cloud_documentai_v1_document_io_proto_rawDesc = nil
	file_google_cloud_documentai_v1_document_io_proto_goTypes = nil
	file_google_cloud_documentai_v1_document_io_proto_depIdxs = nil
}
