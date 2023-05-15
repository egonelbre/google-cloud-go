// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/cloud/datalabeling/v1beta1/annotation_spec_set.proto

package datalabelingpb

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

// An AnnotationSpecSet is a collection of label definitions. For example, in
// image classification tasks, you define a set of possible labels for images as
// an AnnotationSpecSet. An AnnotationSpecSet is immutable upon creation.
type AnnotationSpecSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The AnnotationSpecSet resource name in the following format:
	//
	// "projects/<var>{project_id}</var>/annotationSpecSets/<var>{annotation_spec_set_id}</var>"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name for AnnotationSpecSet that you define when you
	// create it. Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the annotation specification set.
	// The description can be up to 10,000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required. The array of AnnotationSpecs that you define when you create the
	// AnnotationSpecSet. These are the possible labels for the labeling task.
	AnnotationSpecs []*AnnotationSpec `protobuf:"bytes,4,rep,name=annotation_specs,json=annotationSpecs,proto3" json:"annotation_specs,omitempty"`
	// Output only. The names of any related resources that are blocking changes
	// to the annotation spec set.
	BlockingResources []string `protobuf:"bytes,5,rep,name=blocking_resources,json=blockingResources,proto3" json:"blocking_resources,omitempty"`
}

func (x *AnnotationSpecSet) Reset() {
	*x = AnnotationSpecSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationSpecSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationSpecSet) ProtoMessage() {}

func (x *AnnotationSpecSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationSpecSet.ProtoReflect.Descriptor instead.
func (*AnnotationSpecSet) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationSpecSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnnotationSpecSet) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AnnotationSpecSet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AnnotationSpecSet) GetAnnotationSpecs() []*AnnotationSpec {
	if x != nil {
		return x.AnnotationSpecs
	}
	return nil
}

func (x *AnnotationSpecSet) GetBlockingResources() []string {
	if x != nil {
		return x.BlockingResources
	}
	return nil
}

// Container of information related to one possible annotation that can be used
// in a labeling task. For example, an image classification task where images
// are labeled as `dog` or `cat` must reference an AnnotationSpec for `dog` and
// an AnnotationSpec for `cat`.
type AnnotationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The display name of the AnnotationSpec. Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the annotation specification.
	// The description can be up to 10,000 characters long.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AnnotationSpec) Reset() {
	*x = AnnotationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationSpec) ProtoMessage() {}

func (x *AnnotationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationSpec.ProtoReflect.Descriptor instead.
func (*AnnotationSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescGZIP(), []int{1}
}

func (x *AnnotationSpec) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AnnotationSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto protoreflect.FileDescriptor

var file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x11,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x10, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x6f, 0xea, 0x41, 0x6c, 0x0a, 0x2d, 0x64, 0x61,
	0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65, 0x74, 0x12, 0x3b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65,
	0x74, 0x73, 0x2f, 0x7b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x7d, 0x22, 0x55, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xe3, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x49, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescOnce sync.Once
	file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescData = file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDesc
)

func file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescGZIP() []byte {
	file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescOnce.Do(func() {
		file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescData)
	})
	return file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDescData
}

var file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_goTypes = []interface{}{
	(*AnnotationSpecSet)(nil), // 0: google.cloud.datalabeling.v1beta1.AnnotationSpecSet
	(*AnnotationSpec)(nil),    // 1: google.cloud.datalabeling.v1beta1.AnnotationSpec
}
var file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_depIdxs = []int32{
	1, // 0: google.cloud.datalabeling.v1beta1.AnnotationSpecSet.annotation_specs:type_name -> google.cloud.datalabeling.v1beta1.AnnotationSpec
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_init() }
func file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_init() {
	if File_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationSpecSet); i {
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
		file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationSpec); i {
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
			RawDescriptor: file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_goTypes,
		DependencyIndexes: file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_depIdxs,
		MessageInfos:      file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_msgTypes,
	}.Build()
	File_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto = out.File
	file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_rawDesc = nil
	file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_goTypes = nil
	file_google_cloud_datalabeling_v1beta1_annotation_spec_set_proto_depIdxs = nil
}
