// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package order

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateAllOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_CreateAllOrdersClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateAllOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_CreateAllOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], "/OrderService/CreateAllOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceCreateAllOrdersClient{stream}
	return x, nil
}

type OrderService_CreateAllOrdersClient interface {
	Send(*OrderRequest) error
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type orderServiceCreateAllOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceCreateAllOrdersClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceCreateAllOrdersClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateAllOrders(OrderService_CreateAllOrdersServer) error
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateAllOrders(OrderService_CreateAllOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateAllOrders not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateAllOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).CreateAllOrders(&orderServiceCreateAllOrdersServer{stream})
}

type OrderService_CreateAllOrdersServer interface {
	Send(*OrderResponse) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderServiceCreateAllOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceCreateAllOrdersServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceCreateAllOrdersServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateAllOrders",
			Handler:       _OrderService_CreateAllOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/order/order.proto",
}
