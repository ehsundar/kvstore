// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: examples/example.proto

package examples

import (
	_ "github.com/ehsundar/kvstore/protobuf/kvstore"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StaticKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StaticKey) Reset() {
	*x = StaticKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticKey) ProtoMessage() {}

func (x *StaticKey) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticKey.ProtoReflect.Descriptor instead.
func (*StaticKey) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{0}
}

type StaticValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  bool                     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Phones []string                 `protobuf:"bytes,2,rep,name=phones,proto3" json:"phones,omitempty"`
	Items  *StaticValue_NestedItems `protobuf:"bytes,3,opt,name=items,proto3" json:"items,omitempty"`
}

func (x *StaticValue) Reset() {
	*x = StaticValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticValue) ProtoMessage() {}

func (x *StaticValue) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticValue.ProtoReflect.Descriptor instead.
func (*StaticValue) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{1}
}

func (x *StaticValue) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

func (x *StaticValue) GetPhones() []string {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *StaticValue) GetItems() *StaticValue_NestedItems {
	if x != nil {
		return x.Items
	}
	return nil
}

type DynamicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpcName  string `protobuf:"bytes,1,opt,name=rpc_name,json=rpcName,proto3" json:"rpc_name,omitempty"`
	CallerId string `protobuf:"bytes,2,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	Bucket   int64  `protobuf:"varint,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *DynamicKey) Reset() {
	*x = DynamicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicKey) ProtoMessage() {}

func (x *DynamicKey) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicKey.ProtoReflect.Descriptor instead.
func (*DynamicKey) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{2}
}

func (x *DynamicKey) GetRpcName() string {
	if x != nil {
		return x.RpcName
	}
	return ""
}

func (x *DynamicKey) GetCallerId() string {
	if x != nil {
		return x.CallerId
	}
	return ""
}

func (x *DynamicKey) GetBucket() int64 {
	if x != nil {
		return x.Bucket
	}
	return 0
}

type RateLimitCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Limit uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RateLimitCount) Reset() {
	*x = RateLimitCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitCount) ProtoMessage() {}

func (x *RateLimitCount) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitCount.ProtoReflect.Descriptor instead.
func (*RateLimitCount) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitCount) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RateLimitCount) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type OnlineSessionsKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnlineSessionsKey) Reset() {
	*x = OnlineSessionsKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineSessionsKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineSessionsKey) ProtoMessage() {}

func (x *OnlineSessionsKey) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineSessionsKey.ProtoReflect.Descriptor instead.
func (*OnlineSessionsKey) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{4}
}

type OnlineSessionsValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *OnlineSessionsValue) Reset() {
	*x = OnlineSessionsValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineSessionsValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineSessionsValue) ProtoMessage() {}

func (x *OnlineSessionsValue) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineSessionsValue.ProtoReflect.Descriptor instead.
func (*OnlineSessionsValue) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{5}
}

func (x *OnlineSessionsValue) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StaticValue_NestedItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []int32 `protobuf:"varint,1,rep,packed,name=items,proto3" json:"items,omitempty"`
}

func (x *StaticValue_NestedItems) Reset() {
	*x = StaticValue_NestedItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_example_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticValue_NestedItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticValue_NestedItems) ProtoMessage() {}

func (x *StaticValue_NestedItems) ProtoReflect() protoreflect.Message {
	mi := &file_examples_example_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticValue_NestedItems.ProtoReflect.Descriptor instead.
func (*StaticValue_NestedItems) Descriptor() ([]byte, []int) {
	return file_examples_example_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StaticValue_NestedItems) GetItems() []int32 {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_examples_example_proto protoreflect.FileDescriptor

var file_examples_example_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x3a, 0x03, 0xc2, 0x3e, 0x00, 0x22, 0x9d, 0x01, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x23, 0x0a, 0x0b, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x03, 0xca, 0x3e, 0x00, 0x22, 0x6f, 0x0a, 0x0a,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x70,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x70,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x11, 0xc2, 0x3e, 0x0e, 0x0a,
	0x0a, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x00, 0x22, 0x4d, 0x0a,
	0x0e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0f, 0xca, 0x3e, 0x0c,
	0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x18, 0x0a, 0x11,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4b, 0x65,
	0x79, 0x3a, 0x03, 0xc2, 0x3e, 0x00, 0x22, 0x30, 0x0a, 0x13, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x3a, 0x03, 0xca, 0x3e, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x68, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x2f,
	0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_example_proto_rawDescOnce sync.Once
	file_examples_example_proto_rawDescData = file_examples_example_proto_rawDesc
)

func file_examples_example_proto_rawDescGZIP() []byte {
	file_examples_example_proto_rawDescOnce.Do(func() {
		file_examples_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_example_proto_rawDescData)
	})
	return file_examples_example_proto_rawDescData
}

var file_examples_example_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_examples_example_proto_goTypes = []interface{}{
	(*StaticKey)(nil),               // 0: example.StaticKey
	(*StaticValue)(nil),             // 1: example.StaticValue
	(*DynamicKey)(nil),              // 2: example.DynamicKey
	(*RateLimitCount)(nil),          // 3: example.RateLimitCount
	(*OnlineSessionsKey)(nil),       // 4: example.OnlineSessionsKey
	(*OnlineSessionsValue)(nil),     // 5: example.OnlineSessionsValue
	(*StaticValue_NestedItems)(nil), // 6: example.StaticValue.NestedItems
}
var file_examples_example_proto_depIdxs = []int32{
	6, // 0: example.StaticValue.items:type_name -> example.StaticValue.NestedItems
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_examples_example_proto_init() }
func file_examples_example_proto_init() {
	if File_examples_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticKey); i {
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
		file_examples_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticValue); i {
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
		file_examples_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicKey); i {
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
		file_examples_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitCount); i {
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
		file_examples_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineSessionsKey); i {
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
		file_examples_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineSessionsValue); i {
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
		file_examples_example_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticValue_NestedItems); i {
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
			RawDescriptor: file_examples_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_example_proto_goTypes,
		DependencyIndexes: file_examples_example_proto_depIdxs,
		MessageInfos:      file_examples_example_proto_msgTypes,
	}.Build()
	File_examples_example_proto = out.File
	file_examples_example_proto_rawDesc = nil
	file_examples_example_proto_goTypes = nil
	file_examples_example_proto_depIdxs = nil
}
