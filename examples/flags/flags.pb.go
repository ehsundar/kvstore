// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: examples/flags/flags.proto

package flags

import (
	_ "github.com/ehsundar/kvstore/protobuf/kvstore"
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

type FlagKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlagKey) Reset() {
	*x = FlagKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_flags_flags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagKey) ProtoMessage() {}

func (x *FlagKey) ProtoReflect() protoreflect.Message {
	mi := &file_examples_flags_flags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagKey.ProtoReflect.Descriptor instead.
func (*FlagKey) Descriptor() ([]byte, []int) {
	return file_examples_flags_flags_proto_rawDescGZIP(), []int{0}
}

type FlagValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint          string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ExperimentPercent int32  `protobuf:"varint,2,opt,name=experiment_percent,json=experimentPercent,proto3" json:"experiment_percent,omitempty"`
}

func (x *FlagValue) Reset() {
	*x = FlagValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_flags_flags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagValue) ProtoMessage() {}

func (x *FlagValue) ProtoReflect() protoreflect.Message {
	mi := &file_examples_flags_flags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagValue.ProtoReflect.Descriptor instead.
func (*FlagValue) Descriptor() ([]byte, []int) {
	return file_examples_flags_flags_proto_rawDescGZIP(), []int{1}
}

func (x *FlagValue) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *FlagValue) GetExperimentPercent() int32 {
	if x != nil {
		return x.ExperimentPercent
	}
	return 0
}

var File_examples_flags_flags_proto protoreflect.FileDescriptor

var file_examples_flags_flags_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x1a, 0x15, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x07, 0x46, 0x6c,
	0x61, 0x67, 0x4b, 0x65, 0x79, 0x3a, 0x03, 0xc2, 0x3e, 0x00, 0x22, 0x5b, 0x0a, 0x09, 0x46, 0x6c,
	0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x3a, 0x03, 0xca, 0x3e, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x68, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x6b,
	0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_flags_flags_proto_rawDescOnce sync.Once
	file_examples_flags_flags_proto_rawDescData = file_examples_flags_flags_proto_rawDesc
)

func file_examples_flags_flags_proto_rawDescGZIP() []byte {
	file_examples_flags_flags_proto_rawDescOnce.Do(func() {
		file_examples_flags_flags_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_flags_flags_proto_rawDescData)
	})
	return file_examples_flags_flags_proto_rawDescData
}

var file_examples_flags_flags_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_flags_flags_proto_goTypes = []interface{}{
	(*FlagKey)(nil),   // 0: flags.FlagKey
	(*FlagValue)(nil), // 1: flags.FlagValue
}
var file_examples_flags_flags_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_examples_flags_flags_proto_init() }
func file_examples_flags_flags_proto_init() {
	if File_examples_flags_flags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_flags_flags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagKey); i {
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
		file_examples_flags_flags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagValue); i {
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
			RawDescriptor: file_examples_flags_flags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_flags_flags_proto_goTypes,
		DependencyIndexes: file_examples_flags_flags_proto_depIdxs,
		MessageInfos:      file_examples_flags_flags_proto_msgTypes,
	}.Build()
	File_examples_flags_flags_proto = out.File
	file_examples_flags_flags_proto_rawDesc = nil
	file_examples_flags_flags_proto_goTypes = nil
	file_examples_flags_flags_proto_depIdxs = nil
}
