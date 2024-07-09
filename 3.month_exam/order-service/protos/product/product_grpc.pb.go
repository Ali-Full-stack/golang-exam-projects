// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProducts(ctx context.Context, opts ...grpc.CallOption) (ProductService_CreateProductsClient, error)
	DeleteProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*ProductResponse, error)
	GetAllProducts(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (ProductService_GetAllProductsClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProducts(ctx context.Context, opts ...grpc.CallOption) (ProductService_CreateProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/ProductService/CreateProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceCreateProductsClient{stream}
	return x, nil
}

type ProductService_CreateProductsClient interface {
	Send(*ProductInfo) error
	CloseAndRecv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceCreateProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceCreateProductsClient) Send(m *ProductInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceCreateProductsClient) CloseAndRecv() (*ProductResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllProducts(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (ProductService_GetAllProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[1], "/ProductService/GetAllProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceGetAllProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_GetAllProductsClient interface {
	Recv() (*ProductInfo, error)
	grpc.ClientStream
}

type productServiceGetAllProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceGetAllProductsClient) Recv() (*ProductInfo, error) {
	m := new(ProductInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProducts(ProductService_CreateProductsServer) error
	DeleteProduct(context.Context, *ProductID) (*ProductResponse, error)
	GetAllProducts(*CategoryRequest, ProductService_GetAllProductsServer) error
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProducts(ProductService_CreateProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProducts not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *ProductID) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) GetAllProducts(*CategoryRequest, ProductService_GetAllProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).CreateProducts(&productServiceCreateProductsServer{stream})
}

type ProductService_CreateProductsServer interface {
	SendAndClose(*ProductResponse) error
	Recv() (*ProductInfo, error)
	grpc.ServerStream
}

type productServiceCreateProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceCreateProductsServer) SendAndClose(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceCreateProductsServer) Recv() (*ProductInfo, error) {
	m := new(ProductInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CategoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).GetAllProducts(m, &productServiceGetAllProductsServer{stream})
}

type ProductService_GetAllProductsServer interface {
	Send(*ProductInfo) error
	grpc.ServerStream
}

type productServiceGetAllProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceGetAllProductsServer) Send(m *ProductInfo) error {
	return x.ServerStream.SendMsg(m)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateProducts",
			Handler:       _ProductService_CreateProducts_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetAllProducts",
			Handler:       _ProductService_GetAllProducts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/product/product.proto",
}
