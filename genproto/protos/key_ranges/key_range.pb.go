// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/key_ranges/key_range.proto

package key_ranges

import (
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

type KeyRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeyRangeRequest) Reset() {
	*x = KeyRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_ranges_key_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRangeRequest) ProtoMessage() {}

func (x *KeyRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_ranges_key_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRangeRequest.ProtoReflect.Descriptor instead.
func (*KeyRangeRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_ranges_key_range_proto_rawDescGZIP(), []int{0}
}

type KeyRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowerBound []byte `protobuf:"bytes,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound []byte `protobuf:"bytes,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
}

func (x *KeyRange) Reset() {
	*x = KeyRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_ranges_key_range_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRange) ProtoMessage() {}

func (x *KeyRange) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_ranges_key_range_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRange.ProtoReflect.Descriptor instead.
func (*KeyRange) Descriptor() ([]byte, []int) {
	return file_protos_key_ranges_key_range_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRange) GetLowerBound() []byte {
	if x != nil {
		return x.LowerBound
	}
	return nil
}

func (x *KeyRange) GetUpperBound() []byte {
	if x != nil {
		return x.UpperBound
	}
	return nil
}

type KeyRangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyRanges []*KeyRange `protobuf:"bytes,1,rep,name=key_ranges,json=keyRanges,proto3" json:"key_ranges,omitempty"`
}

func (x *KeyRangeReply) Reset() {
	*x = KeyRangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_ranges_key_range_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRangeReply) ProtoMessage() {}

func (x *KeyRangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_ranges_key_range_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRangeReply.ProtoReflect.Descriptor instead.
func (*KeyRangeReply) Descriptor() ([]byte, []int) {
	return file_protos_key_ranges_key_range_proto_rawDescGZIP(), []int{2}
}

func (x *KeyRangeReply) GetKeyRanges() []*KeyRange {
	if x != nil {
		return x.KeyRanges
	}
	return nil
}

var File_protos_key_ranges_key_range_proto protoreflect.FileDescriptor

var file_protos_key_ranges_key_range_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72,
	0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x4b,
	0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c,
	0x0a, 0x08, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x50, 0x0a, 0x0d,
	0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f, 0x0a,
	0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e,
	0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x32, 0x73,
	0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x60, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e,
	0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_key_ranges_key_range_proto_rawDescOnce sync.Once
	file_protos_key_ranges_key_range_proto_rawDescData = file_protos_key_ranges_key_range_proto_rawDesc
)

func file_protos_key_ranges_key_range_proto_rawDescGZIP() []byte {
	file_protos_key_ranges_key_range_proto_rawDescOnce.Do(func() {
		file_protos_key_ranges_key_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_key_ranges_key_range_proto_rawDescData)
	})
	return file_protos_key_ranges_key_range_proto_rawDescData
}

var file_protos_key_ranges_key_range_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_key_ranges_key_range_proto_goTypes = []interface{}{
	(*KeyRangeRequest)(nil), // 0: yandex.router.key_ranges.KeyRangeRequest
	(*KeyRange)(nil),        // 1: yandex.router.key_ranges.KeyRange
	(*KeyRangeReply)(nil),   // 2: yandex.router.key_ranges.KeyRangeReply
}
var file_protos_key_ranges_key_range_proto_depIdxs = []int32{
	1, // 0: yandex.router.key_ranges.KeyRangeReply.key_ranges:type_name -> yandex.router.key_ranges.KeyRange
	0, // 1: yandex.router.key_ranges.KeyRangeService.ListKeyRange:input_type -> yandex.router.key_ranges.KeyRangeRequest
	2, // 2: yandex.router.key_ranges.KeyRangeService.ListKeyRange:output_type -> yandex.router.key_ranges.KeyRangeReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_key_ranges_key_range_proto_init() }
func file_protos_key_ranges_key_range_proto_init() {
	if File_protos_key_ranges_key_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_key_ranges_key_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRangeRequest); i {
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
		file_protos_key_ranges_key_range_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRange); i {
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
		file_protos_key_ranges_key_range_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRangeReply); i {
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
			RawDescriptor: file_protos_key_ranges_key_range_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_key_ranges_key_range_proto_goTypes,
		DependencyIndexes: file_protos_key_ranges_key_range_proto_depIdxs,
		MessageInfos:      file_protos_key_ranges_key_range_proto_msgTypes,
	}.Build()
	File_protos_key_ranges_key_range_proto = out.File
	file_protos_key_ranges_key_range_proto_rawDesc = nil
	file_protos_key_ranges_key_range_proto_goTypes = nil
	file_protos_key_ranges_key_range_proto_depIdxs = nil
}