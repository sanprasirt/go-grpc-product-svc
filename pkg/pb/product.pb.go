// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/pb/product.proto

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

type CreateProductRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stock                int64    `protobuf:"varint,2,opt,name=stock,proto3" json:"stock,omitempty"`
	Price                int64    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductRequest) Reset()         { *m = CreateProductRequest{} }
func (m *CreateProductRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProductRequest) ProtoMessage()    {}
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{0}
}

func (m *CreateProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductRequest.Unmarshal(m, b)
}
func (m *CreateProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductRequest.Marshal(b, m, deterministic)
}
func (m *CreateProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductRequest.Merge(m, src)
}
func (m *CreateProductRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProductRequest.Size(m)
}
func (m *CreateProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductRequest proto.InternalMessageInfo

func (m *CreateProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateProductRequest) GetStock() int64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *CreateProductRequest) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type CreateProductResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductResponse) Reset()         { *m = CreateProductResponse{} }
func (m *CreateProductResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProductResponse) ProtoMessage()    {}
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{1}
}

func (m *CreateProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductResponse.Unmarshal(m, b)
}
func (m *CreateProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductResponse.Marshal(b, m, deterministic)
}
func (m *CreateProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductResponse.Merge(m, src)
}
func (m *CreateProductResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProductResponse.Size(m)
}
func (m *CreateProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductResponse proto.InternalMessageInfo

func (m *CreateProductResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateProductResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateProductResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindOneData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Stock                int64    `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOneData) Reset()         { *m = FindOneData{} }
func (m *FindOneData) String() string { return proto.CompactTextString(m) }
func (*FindOneData) ProtoMessage()    {}
func (*FindOneData) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{2}
}

func (m *FindOneData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneData.Unmarshal(m, b)
}
func (m *FindOneData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneData.Marshal(b, m, deterministic)
}
func (m *FindOneData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneData.Merge(m, src)
}
func (m *FindOneData) XXX_Size() int {
	return xxx_messageInfo_FindOneData.Size(m)
}
func (m *FindOneData) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneData.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneData proto.InternalMessageInfo

func (m *FindOneData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FindOneData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindOneData) GetStock() int64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *FindOneData) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type FindOneRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOneRequest) Reset()         { *m = FindOneRequest{} }
func (m *FindOneRequest) String() string { return proto.CompactTextString(m) }
func (*FindOneRequest) ProtoMessage()    {}
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{3}
}

func (m *FindOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneRequest.Unmarshal(m, b)
}
func (m *FindOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneRequest.Marshal(b, m, deterministic)
}
func (m *FindOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneRequest.Merge(m, src)
}
func (m *FindOneRequest) XXX_Size() int {
	return xxx_messageInfo_FindOneRequest.Size(m)
}
func (m *FindOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneRequest proto.InternalMessageInfo

func (m *FindOneRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindOneResponse struct {
	Status               int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *FindOneData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FindOneResponse) Reset()         { *m = FindOneResponse{} }
func (m *FindOneResponse) String() string { return proto.CompactTextString(m) }
func (*FindOneResponse) ProtoMessage()    {}
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{4}
}

func (m *FindOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneResponse.Unmarshal(m, b)
}
func (m *FindOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneResponse.Marshal(b, m, deterministic)
}
func (m *FindOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneResponse.Merge(m, src)
}
func (m *FindOneResponse) XXX_Size() int {
	return xxx_messageInfo_FindOneResponse.Size(m)
}
func (m *FindOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneResponse proto.InternalMessageInfo

func (m *FindOneResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FindOneResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *FindOneResponse) GetData() *FindOneData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DecreaseStockRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecreaseStockRequest) Reset()         { *m = DecreaseStockRequest{} }
func (m *DecreaseStockRequest) String() string { return proto.CompactTextString(m) }
func (*DecreaseStockRequest) ProtoMessage()    {}
func (*DecreaseStockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{5}
}

func (m *DecreaseStockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecreaseStockRequest.Unmarshal(m, b)
}
func (m *DecreaseStockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecreaseStockRequest.Marshal(b, m, deterministic)
}
func (m *DecreaseStockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecreaseStockRequest.Merge(m, src)
}
func (m *DecreaseStockRequest) XXX_Size() int {
	return xxx_messageInfo_DecreaseStockRequest.Size(m)
}
func (m *DecreaseStockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecreaseStockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecreaseStockRequest proto.InternalMessageInfo

func (m *DecreaseStockRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DecreaseStockRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type DecreaseStockResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecreaseStockResponse) Reset()         { *m = DecreaseStockResponse{} }
func (m *DecreaseStockResponse) String() string { return proto.CompactTextString(m) }
func (*DecreaseStockResponse) ProtoMessage()    {}
func (*DecreaseStockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_024fb40eddc91aac, []int{6}
}

func (m *DecreaseStockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecreaseStockResponse.Unmarshal(m, b)
}
func (m *DecreaseStockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecreaseStockResponse.Marshal(b, m, deterministic)
}
func (m *DecreaseStockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecreaseStockResponse.Merge(m, src)
}
func (m *DecreaseStockResponse) XXX_Size() int {
	return xxx_messageInfo_DecreaseStockResponse.Size(m)
}
func (m *DecreaseStockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecreaseStockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecreaseStockResponse proto.InternalMessageInfo

func (m *DecreaseStockResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DecreaseStockResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateProductRequest)(nil), "product.CreateProductRequest")
	proto.RegisterType((*CreateProductResponse)(nil), "product.CreateProductResponse")
	proto.RegisterType((*FindOneData)(nil), "product.FindOneData")
	proto.RegisterType((*FindOneRequest)(nil), "product.FindOneRequest")
	proto.RegisterType((*FindOneResponse)(nil), "product.FindOneResponse")
	proto.RegisterType((*DecreaseStockRequest)(nil), "product.DecreaseStockRequest")
	proto.RegisterType((*DecreaseStockResponse)(nil), "product.DecreaseStockResponse")
}

func init() { proto.RegisterFile("pkg/pb/product.proto", fileDescriptor_024fb40eddc91aac) }

var fileDescriptor_024fb40eddc91aac = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd7, 0x76, 0x6e, 0xee, 0x8d, 0x4d, 0x08, 0x9d, 0x16, 0x41, 0x19, 0x39, 0xed, 0xb4,
	0xc1, 0xbc, 0x8b, 0xe8, 0x14, 0x3c, 0x39, 0x3a, 0xf4, 0x20, 0x78, 0xc8, 0x9a, 0x87, 0x94, 0x61,
	0x53, 0x93, 0xcc, 0xbf, 0xdc, 0x3f, 0x40, 0x96, 0xa4, 0xa5, 0xab, 0xdd, 0x65, 0xb7, 0x7e, 0xef,
	0xa5, 0x5f, 0x7f, 0xef, 0x7b, 0x0d, 0x84, 0xf9, 0xe6, 0x73, 0x96, 0xaf, 0x67, 0xb9, 0x14, 0x7c,
	0x9b, 0xe8, 0x69, 0x2e, 0x85, 0x16, 0xa4, 0xeb, 0x24, 0x7d, 0x83, 0xf0, 0x41, 0x22, 0xd3, 0xb8,
	0xb4, 0x85, 0x18, 0xbf, 0xb7, 0xa8, 0x34, 0x21, 0xd0, 0xce, 0xd8, 0x17, 0x46, 0xde, 0xd8, 0x9b,
	0xf4, 0x62, 0xf3, 0x4c, 0x42, 0x38, 0x51, 0x5a, 0x24, 0x9b, 0xc8, 0x1f, 0x7b, 0x93, 0x20, 0xb6,
	0x62, 0x57, 0xcd, 0x65, 0x9a, 0x60, 0x14, 0xd8, 0xaa, 0x11, 0xf4, 0x15, 0x46, 0x35, 0x5f, 0x95,
	0x8b, 0x4c, 0x21, 0x39, 0x87, 0x8e, 0xd2, 0x4c, 0x6f, 0x95, 0xb1, 0x0e, 0x62, 0xa7, 0x76, 0x36,
	0x28, 0xa5, 0x90, 0xc6, 0xbc, 0x17, 0x5b, 0x41, 0x86, 0xe0, 0xa7, 0xdc, 0x39, 0xfb, 0x29, 0xa7,
	0x1f, 0xd0, 0x7f, 0x4a, 0x33, 0xfe, 0x92, 0xe1, 0x82, 0x69, 0xe6, 0xda, 0x5e, 0xd1, 0x2e, 0xa9,
	0xfd, 0x26, 0xea, 0xa0, 0x91, 0xba, 0x5d, 0xa5, 0x1e, 0xc3, 0xd0, 0xd9, 0x17, 0x39, 0xd4, 0xbe,
	0x40, 0x53, 0x38, 0x2b, 0x4f, 0x1c, 0x35, 0xd1, 0x04, 0xda, 0x9c, 0x69, 0x66, 0x68, 0xfa, 0xf3,
	0x70, 0x5a, 0xec, 0xa5, 0x32, 0x56, 0x6c, 0x4e, 0xd0, 0x3b, 0x08, 0x17, 0x98, 0x48, 0x64, 0x0a,
	0x57, 0x3b, 0xe6, 0x03, 0x48, 0x24, 0x82, 0xae, 0x90, 0x1c, 0xe5, 0x33, 0x77, 0x8b, 0x29, 0x24,
	0x7d, 0x84, 0x51, 0xcd, 0xe1, 0x18, 0xe4, 0xf9, 0xaf, 0x07, 0x43, 0xb7, 0xc6, 0x15, 0xca, 0x9f,
	0x34, 0x41, 0xb2, 0x84, 0xc1, 0xde, 0x7a, 0xc9, 0x55, 0x39, 0x48, 0xd3, 0xef, 0x74, 0x79, 0x7d,
	0xa8, 0x6d, 0x81, 0x68, 0x8b, 0xdc, 0x42, 0xd7, 0x45, 0x40, 0x2e, 0xea, 0xa1, 0x14, 0x2e, 0xd1,
	0xff, 0x46, 0xf9, 0xfe, 0x12, 0x06, 0x7b, 0xb3, 0x56, 0x88, 0x9a, 0x52, 0xac, 0x10, 0x35, 0x46,
	0x44, 0x5b, 0xf7, 0xf0, 0x7e, 0x3a, 0x9d, 0xd9, 0xdb, 0xb3, 0xee, 0x98, 0x6b, 0x73, 0xf3, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0xaa, 0x69, 0xa1, 0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error)
	DecreaseStock(ctx context.Context, in *DecreaseStockRequest, opts ...grpc.CallOption) (*DecreaseStockResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DecreaseStock(ctx context.Context, in *DecreaseStockRequest, opts ...grpc.CallOption) (*DecreaseStockResponse, error) {
	out := new(DecreaseStockResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/DecreaseStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error)
	DecreaseStock(context.Context, *DecreaseStockRequest) (*DecreaseStockResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedProductServiceServer) FindOne(ctx context.Context, req *FindOneRequest) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedProductServiceServer) DecreaseStock(ctx context.Context, req *DecreaseStockRequest) (*DecreaseStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseStock not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DecreaseStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DecreaseStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/DecreaseStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DecreaseStock(ctx, req.(*DecreaseStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _ProductService_FindOne_Handler,
		},
		{
			MethodName: "DecreaseStock",
			Handler:    _ProductService_DecreaseStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/product.proto",
}
