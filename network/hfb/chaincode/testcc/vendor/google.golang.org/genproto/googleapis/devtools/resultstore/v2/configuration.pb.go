// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/devtools/resultstore/v2/configuration.proto

package resultstore

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

// Represents a configuration within an Invocation associated with one or more
// ConfiguredTargets. It captures the environment and other settings that
// were used.
type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The format of this Configuration resource name must be:
	// invocations/${INVOCATION_ID}/configs/${url_encode(CONFIG_ID)}
	// The configuration ID of "default" should be preferred for the default
	// configuration in a single-config invocation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the Configuration. They must match
	// the resource name after proper encoding.
	Id *Configuration_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status for this configuration.
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// Attributes that apply only to this configuration.
	ConfigurationAttributes *ConfigurationAttributes `protobuf:"bytes,5,opt,name=configuration_attributes,json=configurationAttributes,proto3" json:"configuration_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty"`
	// A human-readable name for Configuration for UIs.
	// It is recommended that this name be unique.
	// If omitted, UIs should default to configuration_id.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Configuration) GetId() *Configuration_Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Configuration) GetStatusAttributes() *StatusAttributes {
	if x != nil {
		return x.StatusAttributes
	}
	return nil
}

func (x *Configuration) GetConfigurationAttributes() *ConfigurationAttributes {
	if x != nil {
		return x.ConfigurationAttributes
	}
	return nil
}

func (x *Configuration) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Configuration) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// Attributes that apply only to the configuration.
type ConfigurationAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of cpu. (e.g. "x86", "powerpc")
	Cpu string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *ConfigurationAttributes) Reset() {
	*x = ConfigurationAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationAttributes) ProtoMessage() {}

func (x *ConfigurationAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationAttributes.ProtoReflect.Descriptor instead.
func (*ConfigurationAttributes) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigurationAttributes) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

// The resource ID components that identify the Configuration.
type Configuration_Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Configuration ID.
	ConfigurationId string `protobuf:"bytes,2,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
}

func (x *Configuration_Id) Reset() {
	*x = Configuration_Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration_Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration_Id) ProtoMessage() {}

func (x *Configuration_Id) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration_Id.ProtoReflect.Descriptor instead.
func (*Configuration_Id) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_configuration_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Configuration_Id) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

func (x *Configuration_Id) GetConfigurationId() string {
	if x != nil {
		return x.ConfigurationId
	}
	return ""
}

var File_google_devtools_resultstore_v2_configuration_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_configuration_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x04, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x5d, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x17,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x54, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e,
	0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a,
	0x28, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x69, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x7d, 0x22, 0x2b, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70,
	0x75, 0x42, 0x85, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_devtools_resultstore_v2_configuration_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_configuration_proto_rawDescData = file_google_devtools_resultstore_v2_configuration_proto_rawDesc
)

func file_google_devtools_resultstore_v2_configuration_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_configuration_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_configuration_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_configuration_proto_rawDescData
}

var file_google_devtools_resultstore_v2_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_devtools_resultstore_v2_configuration_proto_goTypes = []interface{}{
	(*Configuration)(nil),           // 0: google.devtools.resultstore.v2.Configuration
	(*ConfigurationAttributes)(nil), // 1: google.devtools.resultstore.v2.ConfigurationAttributes
	(*Configuration_Id)(nil),        // 2: google.devtools.resultstore.v2.Configuration.Id
	(*StatusAttributes)(nil),        // 3: google.devtools.resultstore.v2.StatusAttributes
	(*Property)(nil),                // 4: google.devtools.resultstore.v2.Property
}
var file_google_devtools_resultstore_v2_configuration_proto_depIdxs = []int32{
	2, // 0: google.devtools.resultstore.v2.Configuration.id:type_name -> google.devtools.resultstore.v2.Configuration.Id
	3, // 1: google.devtools.resultstore.v2.Configuration.status_attributes:type_name -> google.devtools.resultstore.v2.StatusAttributes
	1, // 2: google.devtools.resultstore.v2.Configuration.configuration_attributes:type_name -> google.devtools.resultstore.v2.ConfigurationAttributes
	4, // 3: google.devtools.resultstore.v2.Configuration.properties:type_name -> google.devtools.resultstore.v2.Property
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_configuration_proto_init() }
func file_google_devtools_resultstore_v2_configuration_proto_init() {
	if File_google_devtools_resultstore_v2_configuration_proto != nil {
		return
	}
	file_google_devtools_resultstore_v2_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_google_devtools_resultstore_v2_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurationAttributes); i {
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
		file_google_devtools_resultstore_v2_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration_Id); i {
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
			RawDescriptor: file_google_devtools_resultstore_v2_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_configuration_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_configuration_proto_depIdxs,
		MessageInfos:      file_google_devtools_resultstore_v2_configuration_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_configuration_proto = out.File
	file_google_devtools_resultstore_v2_configuration_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_configuration_proto_goTypes = nil
	file_google_devtools_resultstore_v2_configuration_proto_depIdxs = nil
}
