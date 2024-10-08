// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package paymentProto

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

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	AddClientCard(ctx context.Context, in *CardRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteClientCard(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CardResponse, error)
	MakeOrderPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	MakePurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) AddClientCard(ctx context.Context, in *CardRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/PaymentService/AddClientCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DeleteClientCard(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*CardResponse, error) {
	out := new(CardResponse)
	err := c.cc.Invoke(ctx, "/PaymentService/DeleteClientCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) MakeOrderPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/PaymentService/MakeOrderPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) MakePurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, "/PaymentService/MakePurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	AddClientCard(context.Context, *CardRequest) (*Empty, error)
	DeleteClientCard(context.Context, *RequestId) (*CardResponse, error)
	MakeOrderPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	MakePurchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) AddClientCard(context.Context, *CardRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClientCard not implemented")
}
func (UnimplementedPaymentServiceServer) DeleteClientCard(context.Context, *RequestId) (*CardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientCard not implemented")
}
func (UnimplementedPaymentServiceServer) MakeOrderPayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeOrderPayment not implemented")
}
func (UnimplementedPaymentServiceServer) MakePurchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePurchase not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_AddClientCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AddClientCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/AddClientCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AddClientCard(ctx, req.(*CardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DeleteClientCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DeleteClientCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/DeleteClientCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DeleteClientCard(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_MakeOrderPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).MakeOrderPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/MakeOrderPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).MakeOrderPayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_MakePurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).MakePurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/MakePurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).MakePurchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClientCard",
			Handler:    _PaymentService_AddClientCard_Handler,
		},
		{
			MethodName: "DeleteClientCard",
			Handler:    _PaymentService_DeleteClientCard_Handler,
		},
		{
			MethodName: "MakeOrderPayment",
			Handler:    _PaymentService_MakeOrderPayment_Handler,
		},
		{
			MethodName: "MakePurchase",
			Handler:    _PaymentService_MakePurchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/paymentProto/payment.proto",
}
