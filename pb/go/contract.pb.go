// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: contract.proto

package _go

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

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAlbumRequest) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type CreateAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=Album,proto3" json:"Album,omitempty"`
}

func (x *CreateAlbumResponse) Reset() {
	*x = CreateAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumResponse) ProtoMessage() {}

func (x *CreateAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumResponse.ProtoReflect.Descriptor instead.
func (*CreateAlbumResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type GetAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAlbumRequest) Reset() {
	*x = GetAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumRequest) ProtoMessage() {}

func (x *GetAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

type GetAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album []*Album `protobuf:"bytes,1,rep,name=album,proto3" json:"album,omitempty"`
}

func (x *GetAlbumResponse) Reset() {
	*x = GetAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumResponse) ProtoMessage() {}

func (x *GetAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumResponse.ProtoReflect.Descriptor instead.
func (*GetAlbumResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *GetAlbumResponse) GetAlbum() []*Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type GetAlbumByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAlbumByIdRequest) Reset() {
	*x = GetAlbumByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumByIdRequest) ProtoMessage() {}

func (x *GetAlbumByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumByIdRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumByIdRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *GetAlbumByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAlbumByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *GetAlbumByIdResponse) Reset() {
	*x = GetAlbumByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumByIdResponse) ProtoMessage() {}

func (x *GetAlbumByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumByIdResponse.ProtoReflect.Descriptor instead.
func (*GetAlbumByIdResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{5}
}

func (x *GetAlbumByIdResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x22, 0x42, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x43, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x32, 0x94, 0x02, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x12, 0x58, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x12, 0x22, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_contract_proto_goTypes = []interface{}{
	(*CreateAlbumRequest)(nil),   // 0: hello.world.v1.CreateAlbumRequest
	(*CreateAlbumResponse)(nil),  // 1: hello.world.v1.CreateAlbumResponse
	(*GetAlbumRequest)(nil),      // 2: hello.world.v1.GetAlbumRequest
	(*GetAlbumResponse)(nil),     // 3: hello.world.v1.GetAlbumResponse
	(*GetAlbumByIdRequest)(nil),  // 4: hello.world.v1.GetAlbumByIdRequest
	(*GetAlbumByIdResponse)(nil), // 5: hello.world.v1.GetAlbumByIdResponse
	(*Album)(nil),                // 6: hello.world.v1.Album
}
var file_contract_proto_depIdxs = []int32{
	6, // 0: hello.world.v1.CreateAlbumRequest.album:type_name -> hello.world.v1.Album
	6, // 1: hello.world.v1.CreateAlbumResponse.Album:type_name -> hello.world.v1.Album
	6, // 2: hello.world.v1.GetAlbumResponse.album:type_name -> hello.world.v1.Album
	6, // 3: hello.world.v1.GetAlbumByIdResponse.album:type_name -> hello.world.v1.Album
	0, // 4: hello.world.v1.HelloWorld.CreateAlbum:input_type -> hello.world.v1.CreateAlbumRequest
	2, // 5: hello.world.v1.HelloWorld.GetAlbum:input_type -> hello.world.v1.GetAlbumRequest
	4, // 6: hello.world.v1.HelloWorld.GetAlbumById:input_type -> hello.world.v1.GetAlbumByIdRequest
	1, // 7: hello.world.v1.HelloWorld.CreateAlbum:output_type -> hello.world.v1.CreateAlbumResponse
	3, // 8: hello.world.v1.HelloWorld.GetAlbum:output_type -> hello.world.v1.GetAlbumResponse
	5, // 9: hello.world.v1.HelloWorld.GetAlbumById:output_type -> hello.world.v1.GetAlbumByIdResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	file_album_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumResponse); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumRequest); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumResponse); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumByIdRequest); i {
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
		file_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumByIdResponse); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
