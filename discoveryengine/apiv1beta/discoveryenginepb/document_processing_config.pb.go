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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/discoveryengine/v1beta/document_processing_config.proto

package discoveryenginepb

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

// A singleton resource of
// [DataStore][google.cloud.discoveryengine.v1beta.DataStore]. It's empty when
// [DataStore][google.cloud.discoveryengine.v1beta.DataStore] is created, which
// defaults to digital parser. The first call to
// [DataStoreService.UpdateDocumentProcessingConfig][] method will initialize
// the config.
type DocumentProcessingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full resource name of the Document Processing Config.
	// Format:
	// `projects/*/locations/*/collections/*/dataStores/*/documentProcessingConfig`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Configurations for default Document parser.
	// If not specified, we will configure it as default DigitalParsingConfig, and
	// the default parsing config will be applied to all file types for Document
	// parsing.
	DefaultParsingConfig *DocumentProcessingConfig_ParsingConfig `protobuf:"bytes,4,opt,name=default_parsing_config,json=defaultParsingConfig,proto3" json:"default_parsing_config,omitempty"`
	// Map from file type to override the default parsing configuration based on
	// the file type. Supported keys:
	// * `pdf`: Override parsing config for PDF files, either digital parsing, ocr
	// parsing or layout parsing is supported.
	// * `html`: Override parsing config for HTML files, only digital parsing and
	// or layout parsing are supported.
	// * `docx`: Override parsing config for DOCX files, only digital parsing and
	// or layout parsing are supported.
	ParsingConfigOverrides map[string]*DocumentProcessingConfig_ParsingConfig `protobuf:"bytes,5,rep,name=parsing_config_overrides,json=parsingConfigOverrides,proto3" json:"parsing_config_overrides,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DocumentProcessingConfig) Reset() {
	*x = DocumentProcessingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentProcessingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentProcessingConfig) ProtoMessage() {}

func (x *DocumentProcessingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentProcessingConfig.ProtoReflect.Descriptor instead.
func (*DocumentProcessingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentProcessingConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DocumentProcessingConfig) GetDefaultParsingConfig() *DocumentProcessingConfig_ParsingConfig {
	if x != nil {
		return x.DefaultParsingConfig
	}
	return nil
}

func (x *DocumentProcessingConfig) GetParsingConfigOverrides() map[string]*DocumentProcessingConfig_ParsingConfig {
	if x != nil {
		return x.ParsingConfigOverrides
	}
	return nil
}

// Related configurations applied to a specific type of document parser.
type DocumentProcessingConfig_ParsingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configs for document processing types.
	//
	// Types that are assignable to TypeDedicatedConfig:
	//
	//	*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_
	//	*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_
	TypeDedicatedConfig isDocumentProcessingConfig_ParsingConfig_TypeDedicatedConfig `protobuf_oneof:"type_dedicated_config"`
}

func (x *DocumentProcessingConfig_ParsingConfig) Reset() {
	*x = DocumentProcessingConfig_ParsingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentProcessingConfig_ParsingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentProcessingConfig_ParsingConfig) ProtoMessage() {}

func (x *DocumentProcessingConfig_ParsingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentProcessingConfig_ParsingConfig.ProtoReflect.Descriptor instead.
func (*DocumentProcessingConfig_ParsingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescGZIP(), []int{0, 0}
}

func (m *DocumentProcessingConfig_ParsingConfig) GetTypeDedicatedConfig() isDocumentProcessingConfig_ParsingConfig_TypeDedicatedConfig {
	if m != nil {
		return m.TypeDedicatedConfig
	}
	return nil
}

func (x *DocumentProcessingConfig_ParsingConfig) GetDigitalParsingConfig() *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig {
	if x, ok := x.GetTypeDedicatedConfig().(*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_); ok {
		return x.DigitalParsingConfig
	}
	return nil
}

func (x *DocumentProcessingConfig_ParsingConfig) GetOcrParsingConfig() *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig {
	if x, ok := x.GetTypeDedicatedConfig().(*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_); ok {
		return x.OcrParsingConfig
	}
	return nil
}

type isDocumentProcessingConfig_ParsingConfig_TypeDedicatedConfig interface {
	isDocumentProcessingConfig_ParsingConfig_TypeDedicatedConfig()
}

type DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_ struct {
	// Configurations applied to digital parser.
	DigitalParsingConfig *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig `protobuf:"bytes,1,opt,name=digital_parsing_config,json=digitalParsingConfig,proto3,oneof"`
}

type DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_ struct {
	// Configurations applied to OCR parser. Currently it only applies to
	// PDFs.
	OcrParsingConfig *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig `protobuf:"bytes,2,opt,name=ocr_parsing_config,json=ocrParsingConfig,proto3,oneof"`
}

func (*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_) isDocumentProcessingConfig_ParsingConfig_TypeDedicatedConfig() {
}

func (*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_) isDocumentProcessingConfig_ParsingConfig_TypeDedicatedConfig() {
}

// The digital parsing configurations for documents.
type DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) Reset() {
	*x = DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) ProtoMessage() {}

func (x *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig.ProtoReflect.Descriptor instead.
func (*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescGZIP(), []int{0, 0, 0}
}

// The OCR parsing configurations for documents.
type DocumentProcessingConfig_ParsingConfig_OcrParsingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [DEPRECATED] This field is deprecated. To use the additional enhanced
	// document elements processing, please switch to `layout_parsing_config`.
	//
	// Deprecated: Marked as deprecated in google/cloud/discoveryengine/v1beta/document_processing_config.proto.
	EnhancedDocumentElements []string `protobuf:"bytes,1,rep,name=enhanced_document_elements,json=enhancedDocumentElements,proto3" json:"enhanced_document_elements,omitempty"`
	// If true, will use native text instead of OCR text on pages containing
	// native text.
	UseNativeText bool `protobuf:"varint,2,opt,name=use_native_text,json=useNativeText,proto3" json:"use_native_text,omitempty"`
}

func (x *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) Reset() {
	*x = DocumentProcessingConfig_ParsingConfig_OcrParsingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) ProtoMessage() {}

func (x *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentProcessingConfig_ParsingConfig_OcrParsingConfig.ProtoReflect.Descriptor instead.
func (*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescGZIP(), []int{0, 0, 1}
}

// Deprecated: Marked as deprecated in google/cloud/discoveryengine/v1beta/document_processing_config.proto.
func (x *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) GetEnhancedDocumentElements() []string {
	if x != nil {
		return x.EnhancedDocumentElements
	}
	return nil
}

func (x *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig) GetUseNativeText() bool {
	if x != nil {
		return x.UseNativeText
	}
	return false
}

var File_google_cloud_discoveryengine_v1beta_document_processing_config_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x09, 0x0a, 0x18, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x16, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50,
	0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x93, 0x01, 0x0a,
	0x18, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x59, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50,
	0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x70, 0x61, 0x72, 0x73,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x73, 0x1a, 0xe8, 0x03, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x98, 0x01, 0x0a, 0x16, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x60, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x14, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x8c, 0x01, 0x0a, 0x12, 0x6f, 0x63, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x73,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x63, 0x72, 0x50, 0x61, 0x72,
	0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x10, 0x6f, 0x63,
	0x72, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x16,
	0x0a, 0x14, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x7c, 0x0a, 0x10, 0x4f, 0x63, 0x72, 0x50, 0x61, 0x72,
	0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a, 0x1a, 0x65, 0x6e,
	0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x18, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x75, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x42, 0x17, 0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x96, 0x01,
	0x0a, 0x1b, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x61, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61,
	0x72, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x8a, 0x02, 0xea, 0x41, 0x86, 0x02, 0x0a, 0x37, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x71, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0xa4, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42,
	0x1d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70,
	0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45,
	0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x23, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_goTypes = []interface{}{
	(*DocumentProcessingConfig)(nil),               // 0: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig
	(*DocumentProcessingConfig_ParsingConfig)(nil), // 1: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig
	nil, // 2: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfigOverridesEntry
	(*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig)(nil), // 3: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.DigitalParsingConfig
	(*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig)(nil),     // 4: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig
}
var file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_depIdxs = []int32{
	1, // 0: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.default_parsing_config:type_name -> google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig
	2, // 1: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.parsing_config_overrides:type_name -> google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfigOverridesEntry
	3, // 2: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.digital_parsing_config:type_name -> google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.DigitalParsingConfig
	4, // 3: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.ocr_parsing_config:type_name -> google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig
	1, // 4: google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfigOverridesEntry.value:type_name -> google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_init() }
func file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_document_processing_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentProcessingConfig); i {
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
		file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentProcessingConfig_ParsingConfig); i {
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
		file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig); i {
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
		file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig); i {
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
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig_)(nil),
		(*DocumentProcessingConfig_ParsingConfig_OcrParsingConfig_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_document_processing_config_proto = out.File
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_document_processing_config_proto_depIdxs = nil
}
