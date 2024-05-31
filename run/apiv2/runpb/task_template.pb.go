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
// source: google/cloud/run/v2/task_template.proto

package runpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TaskTemplate describes the data a task should have when created
// from a template.
type TaskTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Holds the single container that defines the unit of execution for this
	// task.
	Containers []*Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	// A list of Volumes to make available to containers.
	Volumes []*Volume `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty"`
	// Types that are assignable to Retries:
	//
	//	*TaskTemplate_MaxRetries
	Retries isTaskTemplate_Retries `protobuf_oneof:"retries"`
	// Max allowed time duration the Task may be active before the system will
	// actively try to mark it failed and kill associated containers. This applies
	// per attempt of a task, meaning each retry can run for the full timeout.
	// Defaults to 600 seconds.
	Timeout *durationpb.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Email address of the IAM service account associated with the Task of a
	// Job. The service account represents the identity of the
	// running task, and determines what permissions the task has. If
	// not provided, the task will use the project's default service account.
	ServiceAccount string `protobuf:"bytes,5,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// The execution environment being used to host this Task.
	ExecutionEnvironment ExecutionEnvironment `protobuf:"varint,6,opt,name=execution_environment,json=executionEnvironment,proto3,enum=google.cloud.run.v2.ExecutionEnvironment" json:"execution_environment,omitempty"`
	// A reference to a customer managed encryption key (CMEK) to use to encrypt
	// this container image. For more information, go to
	// https://cloud.google.com/run/docs/securing/using-cmek
	EncryptionKey string `protobuf:"bytes,7,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
	// VPC Access configuration to use for this Task. For more information,
	// visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
	VpcAccess *VpcAccess `protobuf:"bytes,8,opt,name=vpc_access,json=vpcAccess,proto3" json:"vpc_access,omitempty"`
}

func (x *TaskTemplate) Reset() {
	*x = TaskTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v2_task_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTemplate) ProtoMessage() {}

func (x *TaskTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_task_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTemplate.ProtoReflect.Descriptor instead.
func (*TaskTemplate) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_task_template_proto_rawDescGZIP(), []int{0}
}

func (x *TaskTemplate) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *TaskTemplate) GetVolumes() []*Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (m *TaskTemplate) GetRetries() isTaskTemplate_Retries {
	if m != nil {
		return m.Retries
	}
	return nil
}

func (x *TaskTemplate) GetMaxRetries() int32 {
	if x, ok := x.GetRetries().(*TaskTemplate_MaxRetries); ok {
		return x.MaxRetries
	}
	return 0
}

func (x *TaskTemplate) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *TaskTemplate) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (x *TaskTemplate) GetExecutionEnvironment() ExecutionEnvironment {
	if x != nil {
		return x.ExecutionEnvironment
	}
	return ExecutionEnvironment_EXECUTION_ENVIRONMENT_UNSPECIFIED
}

func (x *TaskTemplate) GetEncryptionKey() string {
	if x != nil {
		return x.EncryptionKey
	}
	return ""
}

func (x *TaskTemplate) GetVpcAccess() *VpcAccess {
	if x != nil {
		return x.VpcAccess
	}
	return nil
}

type isTaskTemplate_Retries interface {
	isTaskTemplate_Retries()
}

type TaskTemplate_MaxRetries struct {
	// Number of retries allowed per Task, before marking this Task failed.
	// Defaults to 3.
	MaxRetries int32 `protobuf:"varint,3,opt,name=max_retries,json=maxRetries,proto3,oneof"`
}

func (*TaskTemplate_MaxRetries) isTaskTemplate_Retries() {}

var File_google_cloud_run_v2_task_template_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_task_template_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x6b,
	0x38, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x03, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x5e, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x14, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x4d, 0x0a, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79,
	0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12,
	0x3d, 0x0a, 0x0a, 0x76, 0x70, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x70, 0x63, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x09, 0x76, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x09,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x59, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x32, 0x42, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x75, 0x6e, 0x70, 0x62, 0x3b, 0x72,
	0x75, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_task_template_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_task_template_proto_rawDescData = file_google_cloud_run_v2_task_template_proto_rawDesc
)

func file_google_cloud_run_v2_task_template_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_task_template_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_task_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_task_template_proto_rawDescData)
	})
	return file_google_cloud_run_v2_task_template_proto_rawDescData
}

var file_google_cloud_run_v2_task_template_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_run_v2_task_template_proto_goTypes = []interface{}{
	(*TaskTemplate)(nil),        // 0: google.cloud.run.v2.TaskTemplate
	(*Container)(nil),           // 1: google.cloud.run.v2.Container
	(*Volume)(nil),              // 2: google.cloud.run.v2.Volume
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(ExecutionEnvironment)(0),   // 4: google.cloud.run.v2.ExecutionEnvironment
	(*VpcAccess)(nil),           // 5: google.cloud.run.v2.VpcAccess
}
var file_google_cloud_run_v2_task_template_proto_depIdxs = []int32{
	1, // 0: google.cloud.run.v2.TaskTemplate.containers:type_name -> google.cloud.run.v2.Container
	2, // 1: google.cloud.run.v2.TaskTemplate.volumes:type_name -> google.cloud.run.v2.Volume
	3, // 2: google.cloud.run.v2.TaskTemplate.timeout:type_name -> google.protobuf.Duration
	4, // 3: google.cloud.run.v2.TaskTemplate.execution_environment:type_name -> google.cloud.run.v2.ExecutionEnvironment
	5, // 4: google.cloud.run.v2.TaskTemplate.vpc_access:type_name -> google.cloud.run.v2.VpcAccess
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_task_template_proto_init() }
func file_google_cloud_run_v2_task_template_proto_init() {
	if File_google_cloud_run_v2_task_template_proto != nil {
		return
	}
	file_google_cloud_run_v2_k8s_min_proto_init()
	file_google_cloud_run_v2_vendor_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v2_task_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskTemplate); i {
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
	file_google_cloud_run_v2_task_template_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TaskTemplate_MaxRetries)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_run_v2_task_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_task_template_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_task_template_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v2_task_template_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_task_template_proto = out.File
	file_google_cloud_run_v2_task_template_proto_rawDesc = nil
	file_google_cloud_run_v2_task_template_proto_goTypes = nil
	file_google_cloud_run_v2_task_template_proto_depIdxs = nil
}
