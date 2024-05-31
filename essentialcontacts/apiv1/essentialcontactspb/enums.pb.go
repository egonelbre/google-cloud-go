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
// source: google/cloud/essentialcontacts/v1/enums.proto

package essentialcontactspb

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

// The notification categories that an essential contact can be subscribed to.
// Each notification will be categorized by the sender into one of the following
// categories. All contacts that are subscribed to that category will receive
// the notification.
type NotificationCategory int32

const (
	// Notification category is unrecognized or unspecified.
	NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED NotificationCategory = 0
	// All notifications related to the resource, including notifications
	// pertaining to categories added in the future.
	NotificationCategory_ALL NotificationCategory = 2
	// Notifications related to imminent account suspension.
	NotificationCategory_SUSPENSION NotificationCategory = 3
	// Notifications related to security/privacy incidents, notifications, and
	// vulnerabilities.
	NotificationCategory_SECURITY NotificationCategory = 5
	// Notifications related to technical events and issues such as outages,
	// errors, or bugs.
	NotificationCategory_TECHNICAL NotificationCategory = 6
	// Notifications related to billing and payments notifications, price updates,
	// errors, or credits.
	NotificationCategory_BILLING NotificationCategory = 7
	// Notifications related to enforcement actions, regulatory compliance, or
	// government notices.
	NotificationCategory_LEGAL NotificationCategory = 8
	// Notifications related to new versions, product terms updates, or
	// deprecations.
	NotificationCategory_PRODUCT_UPDATES NotificationCategory = 9
	// Child category of TECHNICAL. If assigned, technical incident notifications
	// will go to these contacts instead of TECHNICAL.
	NotificationCategory_TECHNICAL_INCIDENTS NotificationCategory = 10
)

// Enum value maps for NotificationCategory.
var (
	NotificationCategory_name = map[int32]string{
		0:  "NOTIFICATION_CATEGORY_UNSPECIFIED",
		2:  "ALL",
		3:  "SUSPENSION",
		5:  "SECURITY",
		6:  "TECHNICAL",
		7:  "BILLING",
		8:  "LEGAL",
		9:  "PRODUCT_UPDATES",
		10: "TECHNICAL_INCIDENTS",
	}
	NotificationCategory_value = map[string]int32{
		"NOTIFICATION_CATEGORY_UNSPECIFIED": 0,
		"ALL":                               2,
		"SUSPENSION":                        3,
		"SECURITY":                          5,
		"TECHNICAL":                         6,
		"BILLING":                           7,
		"LEGAL":                             8,
		"PRODUCT_UPDATES":                   9,
		"TECHNICAL_INCIDENTS":               10,
	}
)

func (x NotificationCategory) Enum() *NotificationCategory {
	p := new(NotificationCategory)
	*p = x
	return p
}

func (x NotificationCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_essentialcontacts_v1_enums_proto_enumTypes[0].Descriptor()
}

func (NotificationCategory) Type() protoreflect.EnumType {
	return &file_google_cloud_essentialcontacts_v1_enums_proto_enumTypes[0]
}

func (x NotificationCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationCategory.Descriptor instead.
func (NotificationCategory) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_essentialcontacts_v1_enums_proto_rawDescGZIP(), []int{0}
}

// A contact's validation state indicates whether or not it is the correct
// contact to be receiving notifications for a particular resource.
type ValidationState int32

const (
	// The validation state is unknown or unspecified.
	ValidationState_VALIDATION_STATE_UNSPECIFIED ValidationState = 0
	// The contact is marked as valid. This is usually done manually by the
	// contact admin. All new contacts begin in the valid state.
	ValidationState_VALID ValidationState = 1
	// The contact is considered invalid. This may become the state if the
	// contact's email is found to be unreachable.
	ValidationState_INVALID ValidationState = 2
)

// Enum value maps for ValidationState.
var (
	ValidationState_name = map[int32]string{
		0: "VALIDATION_STATE_UNSPECIFIED",
		1: "VALID",
		2: "INVALID",
	}
	ValidationState_value = map[string]int32{
		"VALIDATION_STATE_UNSPECIFIED": 0,
		"VALID":                        1,
		"INVALID":                      2,
	}
)

func (x ValidationState) Enum() *ValidationState {
	p := new(ValidationState)
	*p = x
	return p
}

func (x ValidationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidationState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_essentialcontacts_v1_enums_proto_enumTypes[1].Descriptor()
}

func (ValidationState) Type() protoreflect.EnumType {
	return &file_google_cloud_essentialcontacts_v1_enums_proto_enumTypes[1]
}

func (x ValidationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidationState.Descriptor instead.
func (ValidationState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_essentialcontacts_v1_enums_proto_rawDescGZIP(), []int{1}
}

var File_google_cloud_essentialcontacts_v1_enums_proto protoreflect.FileDescriptor

var file_google_cloud_essentialcontacts_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x73,
	0x73, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2a, 0xb9, 0x01, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x55, 0x53, 0x50, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45, 0x43,
	0x48, 0x4e, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x4c, 0x4c,
	0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x08,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x53, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x49, 0x43,
	0x41, 0x4c, 0x5f, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0a, 0x2a, 0x4b,
	0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x1c, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x42, 0xf9, 0x01, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x53, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x70, 0x62, 0x3b, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x73, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x45, 0x73, 0x73, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x45, 0x73, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_essentialcontacts_v1_enums_proto_rawDescOnce sync.Once
	file_google_cloud_essentialcontacts_v1_enums_proto_rawDescData = file_google_cloud_essentialcontacts_v1_enums_proto_rawDesc
)

func file_google_cloud_essentialcontacts_v1_enums_proto_rawDescGZIP() []byte {
	file_google_cloud_essentialcontacts_v1_enums_proto_rawDescOnce.Do(func() {
		file_google_cloud_essentialcontacts_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_essentialcontacts_v1_enums_proto_rawDescData)
	})
	return file_google_cloud_essentialcontacts_v1_enums_proto_rawDescData
}

var file_google_cloud_essentialcontacts_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_essentialcontacts_v1_enums_proto_goTypes = []interface{}{
	(NotificationCategory)(0), // 0: google.cloud.essentialcontacts.v1.NotificationCategory
	(ValidationState)(0),      // 1: google.cloud.essentialcontacts.v1.ValidationState
}
var file_google_cloud_essentialcontacts_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_essentialcontacts_v1_enums_proto_init() }
func file_google_cloud_essentialcontacts_v1_enums_proto_init() {
	if File_google_cloud_essentialcontacts_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_essentialcontacts_v1_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_essentialcontacts_v1_enums_proto_goTypes,
		DependencyIndexes: file_google_cloud_essentialcontacts_v1_enums_proto_depIdxs,
		EnumInfos:         file_google_cloud_essentialcontacts_v1_enums_proto_enumTypes,
	}.Build()
	File_google_cloud_essentialcontacts_v1_enums_proto = out.File
	file_google_cloud_essentialcontacts_v1_enums_proto_rawDesc = nil
	file_google_cloud_essentialcontacts_v1_enums_proto_goTypes = nil
	file_google_cloud_essentialcontacts_v1_enums_proto_depIdxs = nil
}
