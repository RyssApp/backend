// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetStoreRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoreRequest) Reset()         { *m = GetStoreRequest{} }
func (m *GetStoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoreRequest) ProtoMessage()    {}
func (*GetStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9619da4b41557f, []int{0}
}

func (m *GetStoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreRequest.Unmarshal(m, b)
}
func (m *GetStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreRequest.Marshal(b, m, deterministic)
}
func (m *GetStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreRequest.Merge(m, src)
}
func (m *GetStoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoreRequest.Size(m)
}
func (m *GetStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreRequest proto.InternalMessageInfo

func (m *GetStoreRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetStoreResponse struct {
	Store                *Store   `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoreResponse) Reset()         { *m = GetStoreResponse{} }
func (m *GetStoreResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoreResponse) ProtoMessage()    {}
func (*GetStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9619da4b41557f, []int{1}
}

func (m *GetStoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreResponse.Unmarshal(m, b)
}
func (m *GetStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreResponse.Marshal(b, m, deterministic)
}
func (m *GetStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreResponse.Merge(m, src)
}
func (m *GetStoreResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoreResponse.Size(m)
}
func (m *GetStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreResponse proto.InternalMessageInfo

func (m *GetStoreResponse) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type GetStoresRequest struct {
	Location             *Location   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	DisplayName          string      `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetStoresRequest) Reset()         { *m = GetStoresRequest{} }
func (m *GetStoresRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoresRequest) ProtoMessage()    {}
func (*GetStoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9619da4b41557f, []int{2}
}

func (m *GetStoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoresRequest.Unmarshal(m, b)
}
func (m *GetStoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoresRequest.Marshal(b, m, deterministic)
}
func (m *GetStoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoresRequest.Merge(m, src)
}
func (m *GetStoresRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoresRequest.Size(m)
}
func (m *GetStoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoresRequest proto.InternalMessageInfo

func (m *GetStoresRequest) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *GetStoresRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *GetStoresRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type GetStoresResponse struct {
	Stores               []*Store `protobuf:"bytes,1,rep,name=stores,proto3" json:"stores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoresResponse) Reset()         { *m = GetStoresResponse{} }
func (m *GetStoresResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoresResponse) ProtoMessage()    {}
func (*GetStoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9619da4b41557f, []int{3}
}

func (m *GetStoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoresResponse.Unmarshal(m, b)
}
func (m *GetStoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoresResponse.Marshal(b, m, deterministic)
}
func (m *GetStoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoresResponse.Merge(m, src)
}
func (m *GetStoresResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoresResponse.Size(m)
}
func (m *GetStoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoresResponse proto.InternalMessageInfo

func (m *GetStoresResponse) GetStores() []*Store {
	if m != nil {
		return m.Stores
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStoreRequest)(nil), "pb.GetStoreRequest")
	proto.RegisterType((*GetStoreResponse)(nil), "pb.GetStoreResponse")
	proto.RegisterType((*GetStoresRequest)(nil), "pb.GetStoresRequest")
	proto.RegisterType((*GetStoresResponse)(nil), "pb.GetStoresResponse")
}

func init() {
	proto.RegisterFile("store_service.proto", fileDescriptor_bb9619da4b41557f)
}

var fileDescriptor_bb9619da4b41557f = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x9d, 0x14, 0x4b, 0x73, 0x13, 0x62, 0x9d, 0x2a, 0x84, 0x6c, 0x6c, 0x66, 0xd5, 0x55,
	0x16, 0x2d, 0x28, 0xf8, 0x02, 0x6e, 0x44, 0x24, 0x7d, 0x80, 0x32, 0x69, 0x06, 0x19, 0x68, 0x33,
	0x63, 0x66, 0x14, 0x5c, 0xfb, 0x00, 0xbe, 0xb2, 0x64, 0x7e, 0xf2, 0xd3, 0x65, 0xbe, 0x7b, 0xce,
	0xc9, 0xb9, 0x77, 0x60, 0xa5, 0xb4, 0x68, 0xd9, 0x41, 0xb1, 0xf6, 0x9b, 0x1f, 0x59, 0x21, 0x5b,
	0xa1, 0x05, 0x0e, 0x64, 0x95, 0x25, 0x27, 0x71, 0xa4, 0x9a, 0x8b, 0xc6, 0xb2, 0x6c, 0x29, 0xe9,
	0x07, 0x6f, 0xc6, 0x24, 0x32, 0x56, 0xfb, 0x41, 0x72, 0xb8, 0x79, 0x61, 0x7a, 0xdf, 0x91, 0x92,
	0x7d, 0x7e, 0x31, 0xa5, 0x71, 0x02, 0x01, 0xaf, 0x53, 0xb4, 0x46, 0x9b, 0xb0, 0x0c, 0x78, 0x4d,
	0x76, 0xb0, 0x1c, 0x24, 0x4a, 0x8a, 0x46, 0x31, 0xfc, 0x00, 0xd7, 0x26, 0xc5, 0xc8, 0xa2, 0x6d,
	0x58, 0xc8, 0xaa, 0xb0, 0x0a, 0xcb, 0xc9, 0x1f, 0x1a, 0x5c, 0xca, 0x27, 0x6f, 0x60, 0xe1, 0xdb,
	0x39, 0x63, 0xdc, 0x19, 0x5f, 0x1d, 0x2b, 0xfb, 0x29, 0xce, 0x21, 0xae, 0xb9, 0x92, 0x27, 0xfa,
	0x73, 0x68, 0xe8, 0x99, 0xa5, 0x81, 0x69, 0x13, 0x39, 0xf6, 0x46, 0xcf, 0x0c, 0x17, 0x00, 0xc3,
	0x6a, 0xe9, 0xcc, 0xc4, 0x25, 0x5d, 0xdc, 0x7b, 0x4f, 0xcb, 0x91, 0x82, 0x3c, 0xc2, 0xed, 0xa8,
	0x90, 0xdb, 0x23, 0x87, 0xb9, 0xe9, 0xab, 0x52, 0xb4, 0x9e, 0x4d, 0x17, 0x71, 0x83, 0xed, 0x2f,
	0x82, 0xd8, 0x90, 0xbd, 0xbd, 0x35, 0x7e, 0x82, 0x85, 0x0f, 0xc2, 0xab, 0x4e, 0x7f, 0x71, 0xc0,
	0xec, 0x6e, 0x0a, 0xed, 0xaf, 0xc8, 0x15, 0x7e, 0x86, 0xb0, 0x6f, 0x80, 0x27, 0x22, 0x7f, 0xa1,
	0xec, 0xfe, 0x82, 0x7a, 0x6f, 0x35, 0x37, 0xcf, 0xb5, 0xfb, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0x8e, 0xe1, 0xa9, 0xf8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*GetStoreResponse, error)
	GetStores(ctx context.Context, in *GetStoresRequest, opts ...grpc.CallOption) (*GetStoresResponse, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*GetStoreResponse, error) {
	out := new(GetStoreResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) GetStores(ctx context.Context, in *GetStoresRequest, opts ...grpc.CallOption) (*GetStoresResponse, error) {
	out := new(GetStoresResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreService/GetStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	GetStore(context.Context, *GetStoreRequest) (*GetStoreResponse, error)
	GetStores(context.Context, *GetStoresRequest) (*GetStoresResponse, error)
}

// UnimplementedStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (*UnimplementedStoreServiceServer) GetStore(ctx context.Context, req *GetStoreRequest) (*GetStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStore not implemented")
}
func (*UnimplementedStoreServiceServer) GetStores(ctx context.Context, req *GetStoresRequest) (*GetStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStores not implemented")
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetStore(ctx, req.(*GetStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_GetStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreService/GetStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetStores(ctx, req.(*GetStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStore",
			Handler:    _StoreService_GetStore_Handler,
		},
		{
			MethodName: "GetStores",
			Handler:    _StoreService_GetStores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store_service.proto",
}
