// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: club.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Query) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{2}
}

func (x *GetReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type IndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   int32    `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`    // from 20
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`    // take 30
	Filter []*Query `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty"` // filter
	Token  string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *IndexRequest) Reset() {
	*x = IndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRequest) ProtoMessage() {}

func (x *IndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRequest.ProtoReflect.Descriptor instead.
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{3}
}

func (x *IndexRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *IndexRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *IndexRequest) GetFilter() []*Query {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *IndexRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IndexReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items   []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IndexReply) Reset() {
	*x = IndexReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReply) ProtoMessage() {}

func (x *IndexReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexReply.ProtoReflect.Descriptor instead.
func (*IndexReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{4}
}

func (x *IndexReply) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *IndexReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *IndexReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Point uint64 `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{5}
}

func (x *Item) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *Item) GetPoint() uint64 {
	if x != nil {
		return x.Point
	}
	return 0
}

var File_club_proto protoreflect.FileDescriptor

var file_club_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x2f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x22, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6f, 0x0a, 0x0c, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5e, 0x0a, 0x0a,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0x56,
	0x0a, 0x04, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x23, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_club_proto_rawDescOnce sync.Once
	file_club_proto_rawDescData = file_club_proto_rawDesc
)

func file_club_proto_rawDescGZIP() []byte {
	file_club_proto_rawDescOnce.Do(func() {
		file_club_proto_rawDescData = protoimpl.X.CompressGZIP(file_club_proto_rawDescData)
	})
	return file_club_proto_rawDescData
}

var file_club_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_club_proto_goTypes = []interface{}{
	(*Query)(nil),        // 0: pb.Query
	(*GetRequest)(nil),   // 1: pb.GetRequest
	(*GetReply)(nil),     // 2: pb.GetReply
	(*IndexRequest)(nil), // 3: pb.IndexRequest
	(*IndexReply)(nil),   // 4: pb.IndexReply
	(*Item)(nil),         // 5: pb.Item
}
var file_club_proto_depIdxs = []int32{
	0, // 0: pb.IndexRequest.filter:type_name -> pb.Query
	5, // 1: pb.IndexReply.items:type_name -> pb.Item
	1, // 2: pb.Club.Get:input_type -> pb.GetRequest
	3, // 3: pb.Club.Index:input_type -> pb.IndexRequest
	2, // 4: pb.Club.Get:output_type -> pb.GetReply
	4, // 5: pb.Club.Index:output_type -> pb.IndexReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_club_proto_init() }
func file_club_proto_init() {
	if File_club_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_club_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_club_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_club_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
		file_club_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRequest); i {
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
		file_club_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexReply); i {
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
		file_club_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_club_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_club_proto_goTypes,
		DependencyIndexes: file_club_proto_depIdxs,
		MessageInfos:      file_club_proto_msgTypes,
	}.Build()
	File_club_proto = out.File
	file_club_proto_rawDesc = nil
	file_club_proto_goTypes = nil
	file_club_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClubClient is the client API for Club service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClubClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexReply, error)
}

type clubClient struct {
	cc grpc.ClientConnInterface
}

func NewClubClient(cc grpc.ClientConnInterface) ClubClient {
	return &clubClient{cc}
}

func (c *clubClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/pb.Club/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexReply, error) {
	out := new(IndexReply)
	err := c.cc.Invoke(ctx, "/pb.Club/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClubServer is the server API for Club service.
type ClubServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Index(context.Context, *IndexRequest) (*IndexReply, error)
}

// UnimplementedClubServer can be embedded to have forward compatible implementations.
type UnimplementedClubServer struct {
}

func (*UnimplementedClubServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedClubServer) Index(context.Context, *IndexRequest) (*IndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}

func RegisterClubServer(s *grpc.Server, srv ClubServer) {
	s.RegisterService(&_Club_serviceDesc, srv)
}

func _Club_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Club/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Club_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Club/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Club_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Club",
	HandlerType: (*ClubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Club_Get_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Club_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "club.proto",
}
