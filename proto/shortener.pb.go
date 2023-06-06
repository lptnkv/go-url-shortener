// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: shortener.proto

package proto

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

type ShortenUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenUrlRequest) Reset() {
	*x = ShortenUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenUrlRequest) ProtoMessage() {}

func (x *ShortenUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenUrlRequest.ProtoReflect.Descriptor instead.
func (*ShortenUrlRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenUrlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenUrlReply) Reset() {
	*x = ShortenUrlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenUrlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenUrlReply) ProtoMessage() {}

func (x *ShortenUrlReply) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenUrlReply.ProtoReflect.Descriptor instead.
func (*ShortenUrlReply) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenUrlReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetFullUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetFullUrlRequest) Reset() {
	*x = GetFullUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullUrlRequest) ProtoMessage() {}

func (x *GetFullUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullUrlRequest.ProtoReflect.Descriptor instead.
func (*GetFullUrlRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetFullUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetFullUrlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetFullUrlReply) Reset() {
	*x = GetFullUrlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullUrlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullUrlReply) ProtoMessage() {}

func (x *GetFullUrlReply) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullUrlReply.ProtoReflect.Descriptor instead.
func (*GetFullUrlReply) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetFullUrlReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_shortener_proto protoreflect.FileDescriptor

var file_shortener_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x25, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x73, 0x0a, 0x09, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75,
	0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05,
	0x5a, 0x03, 0x2e, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_proto_rawDescOnce sync.Once
	file_shortener_proto_rawDescData = file_shortener_proto_rawDesc
)

func file_shortener_proto_rawDescGZIP() []byte {
	file_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_proto_rawDescData)
	})
	return file_shortener_proto_rawDescData
}

var file_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortener_proto_goTypes = []interface{}{
	(*ShortenUrlRequest)(nil), // 0: ShortenUrlRequest
	(*ShortenUrlReply)(nil),   // 1: ShortenUrlReply
	(*GetFullUrlRequest)(nil), // 2: GetFullUrlRequest
	(*GetFullUrlReply)(nil),   // 3: GetFullUrlReply
}
var file_shortener_proto_depIdxs = []int32{
	0, // 0: Shortener.ShortenUrl:input_type -> ShortenUrlRequest
	2, // 1: Shortener.GetFullUrl:input_type -> GetFullUrlRequest
	1, // 2: Shortener.ShortenUrl:output_type -> ShortenUrlReply
	3, // 3: Shortener.GetFullUrl:output_type -> GetFullUrlReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortener_proto_init() }
func file_shortener_proto_init() {
	if File_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenUrlRequest); i {
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
		file_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenUrlReply); i {
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
		file_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullUrlRequest); i {
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
		file_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullUrlReply); i {
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
			RawDescriptor: file_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_proto_msgTypes,
	}.Build()
	File_shortener_proto = out.File
	file_shortener_proto_rawDesc = nil
	file_shortener_proto_goTypes = nil
	file_shortener_proto_depIdxs = nil
}