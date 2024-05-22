// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: einride/sensitive/example/v1/example.proto

package examplev1

import (
	_ "go.einride.tech/protobuf-sensitive/gen/einride/sensitive/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Example message for testing and validation of sensitive fields.
type ExampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensitiveField    string                           `protobuf:"bytes,1,opt,name=sensitive_field,json=sensitiveField,proto3" json:"sensitive_field,omitempty"`
	NonSensitiveField string                           `protobuf:"bytes,2,opt,name=non_sensitive_field,json=nonSensitiveField,proto3" json:"non_sensitive_field,omitempty"`
	NonStringField    int64                            `protobuf:"varint,3,opt,name=non_string_field,json=nonStringField,proto3" json:"non_string_field,omitempty"`
	Nested            *ExampleMessage_Nested           `protobuf:"bytes,4,opt,name=nested,proto3" json:"nested,omitempty"`
	RepeatedNested    []*ExampleMessage_Nested         `protobuf:"bytes,5,rep,name=repeated_nested,json=repeatedNested,proto3" json:"repeated_nested,omitempty"`
	MapNested         map[int64]*ExampleMessage_Nested `protobuf:"bytes,6,rep,name=map_nested,json=mapNested,proto3" json:"map_nested,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExampleMessage) Reset() {
	*x = ExampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_sensitive_example_v1_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage) ProtoMessage() {}

func (x *ExampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_einride_sensitive_example_v1_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage.ProtoReflect.Descriptor instead.
func (*ExampleMessage) Descriptor() ([]byte, []int) {
	return file_einride_sensitive_example_v1_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleMessage) GetSensitiveField() string {
	if x != nil {
		return x.SensitiveField
	}
	return ""
}

func (x *ExampleMessage) GetNonSensitiveField() string {
	if x != nil {
		return x.NonSensitiveField
	}
	return ""
}

func (x *ExampleMessage) GetNonStringField() int64 {
	if x != nil {
		return x.NonStringField
	}
	return 0
}

func (x *ExampleMessage) GetNested() *ExampleMessage_Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *ExampleMessage) GetRepeatedNested() []*ExampleMessage_Nested {
	if x != nil {
		return x.RepeatedNested
	}
	return nil
}

func (x *ExampleMessage) GetMapNested() map[int64]*ExampleMessage_Nested {
	if x != nil {
		return x.MapNested
	}
	return nil
}

// Nested example message.
type ExampleMessage_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensitiveField        string `protobuf:"bytes,1,opt,name=sensitive_field,json=sensitiveField,proto3" json:"sensitive_field,omitempty"`
	NonSensitiveField     string `protobuf:"bytes,2,opt,name=non_sensitive_field,json=nonSensitiveField,proto3" json:"non_sensitive_field,omitempty"`
	NonStringField        int64  `protobuf:"varint,3,opt,name=non_string_field,json=nonStringField,proto3" json:"non_string_field,omitempty"`
	AnotherSensitiveField string `protobuf:"bytes,4,opt,name=another_sensitive_field,json=anotherSensitiveField,proto3" json:"another_sensitive_field,omitempty"`
}

func (x *ExampleMessage_Nested) Reset() {
	*x = ExampleMessage_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_sensitive_example_v1_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleMessage_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMessage_Nested) ProtoMessage() {}

func (x *ExampleMessage_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_einride_sensitive_example_v1_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMessage_Nested.ProtoReflect.Descriptor instead.
func (*ExampleMessage_Nested) Descriptor() ([]byte, []int) {
	return file_einride_sensitive_example_v1_example_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ExampleMessage_Nested) GetSensitiveField() string {
	if x != nil {
		return x.SensitiveField
	}
	return ""
}

func (x *ExampleMessage_Nested) GetNonSensitiveField() string {
	if x != nil {
		return x.NonSensitiveField
	}
	return ""
}

func (x *ExampleMessage_Nested) GetNonStringField() int64 {
	if x != nil {
		return x.NonStringField
	}
	return 0
}

func (x *ExampleMessage_Nested) GetAnotherSensitiveField() string {
	if x != nil {
		return x.AnotherSensitiveField
	}
	return ""
}

var File_einride_sensitive_example_v1_example_proto protoreflect.FileDescriptor

var file_einride_sensitive_example_v1_example_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe8, 0x05, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05,
	0xd0, 0xd7, 0xcb, 0x29, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x6e, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x4b, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x5c, 0x0a, 0x0f,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x5a, 0x0a, 0x0a, 0x6d, 0x61,
	0x70, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x71, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x49, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xd1, 0x01, 0x0a, 0x06, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xd0,
	0xd7, 0xcb, 0x29, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6e, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x6e, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3d,
	0x0a, 0x17, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x05, 0xd0, 0xd7, 0xcb, 0x29, 0x01, 0x52, 0x15, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x92, 0x02,
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4d, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x45, 0x53, 0x45, 0xaa, 0x02, 0x1c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1c, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x5c, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x28, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c,
	0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1f, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_sensitive_example_v1_example_proto_rawDescOnce sync.Once
	file_einride_sensitive_example_v1_example_proto_rawDescData = file_einride_sensitive_example_v1_example_proto_rawDesc
)

func file_einride_sensitive_example_v1_example_proto_rawDescGZIP() []byte {
	file_einride_sensitive_example_v1_example_proto_rawDescOnce.Do(func() {
		file_einride_sensitive_example_v1_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_sensitive_example_v1_example_proto_rawDescData)
	})
	return file_einride_sensitive_example_v1_example_proto_rawDescData
}

var file_einride_sensitive_example_v1_example_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_sensitive_example_v1_example_proto_goTypes = []interface{}{
	(*ExampleMessage)(nil),        // 0: einride.sensitive.example.v1.ExampleMessage
	nil,                           // 1: einride.sensitive.example.v1.ExampleMessage.MapNestedEntry
	(*ExampleMessage_Nested)(nil), // 2: einride.sensitive.example.v1.ExampleMessage.Nested
}
var file_einride_sensitive_example_v1_example_proto_depIdxs = []int32{
	2, // 0: einride.sensitive.example.v1.ExampleMessage.nested:type_name -> einride.sensitive.example.v1.ExampleMessage.Nested
	2, // 1: einride.sensitive.example.v1.ExampleMessage.repeated_nested:type_name -> einride.sensitive.example.v1.ExampleMessage.Nested
	1, // 2: einride.sensitive.example.v1.ExampleMessage.map_nested:type_name -> einride.sensitive.example.v1.ExampleMessage.MapNestedEntry
	2, // 3: einride.sensitive.example.v1.ExampleMessage.MapNestedEntry.value:type_name -> einride.sensitive.example.v1.ExampleMessage.Nested
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_einride_sensitive_example_v1_example_proto_init() }
func file_einride_sensitive_example_v1_example_proto_init() {
	if File_einride_sensitive_example_v1_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_sensitive_example_v1_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleMessage); i {
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
		file_einride_sensitive_example_v1_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleMessage_Nested); i {
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
			RawDescriptor: file_einride_sensitive_example_v1_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_sensitive_example_v1_example_proto_goTypes,
		DependencyIndexes: file_einride_sensitive_example_v1_example_proto_depIdxs,
		MessageInfos:      file_einride_sensitive_example_v1_example_proto_msgTypes,
	}.Build()
	File_einride_sensitive_example_v1_example_proto = out.File
	file_einride_sensitive_example_v1_example_proto_rawDesc = nil
	file_einride_sensitive_example_v1_example_proto_goTypes = nil
	file_einride_sensitive_example_v1_example_proto_depIdxs = nil
}
