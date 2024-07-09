// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package clientProto

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

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateClient(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*ClientID, error)
	GetClientById(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientInfo, error)
	GetAllClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_GetAllClientsClient, error)
	GetClientLocation(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientLocation, error)
	DeleteClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateClient(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*ClientID, error) {
	out := new(ClientID)
	err := c.cc.Invoke(ctx, "/ClientService/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientById(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientInfo, error) {
	out := new(ClientInfo)
	err := c.cc.Invoke(ctx, "/ClientService/GetClientById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetAllClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_GetAllClientsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], "/ClientService/GetAllClients", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceGetAllClientsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_GetAllClientsClient interface {
	Recv() (*ClientInfo, error)
	grpc.ClientStream
}

type clientServiceGetAllClientsClient struct {
	grpc.ClientStream
}

func (x *clientServiceGetAllClientsClient) Recv() (*ClientInfo, error) {
	m := new(ClientInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) GetClientLocation(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientLocation, error) {
	out := new(ClientLocation)
	err := c.cc.Invoke(ctx, "/ClientService/GetClientLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/ClientService/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateClient(context.Context, *ClientInfo) (*ClientID, error)
	GetClientById(context.Context, *ClientID) (*ClientInfo, error)
	GetAllClients(*Empty, ClientService_GetAllClientsServer) error
	GetClientLocation(context.Context, *ClientID) (*ClientLocation, error)
	DeleteClient(context.Context, *ClientID) (*ClientResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateClient(context.Context, *ClientInfo) (*ClientID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedClientServiceServer) GetClientById(context.Context, *ClientID) (*ClientInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientById not implemented")
}
func (UnimplementedClientServiceServer) GetAllClients(*Empty, ClientService_GetAllClientsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllClients not implemented")
}
func (UnimplementedClientServiceServer) GetClientLocation(context.Context, *ClientID) (*ClientLocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientLocation not implemented")
}
func (UnimplementedClientServiceServer) DeleteClient(context.Context, *ClientID) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateClient(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/GetClientById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientById(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetAllClients_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).GetAllClients(m, &clientServiceGetAllClientsServer{stream})
}

type ClientService_GetAllClientsServer interface {
	Send(*ClientInfo) error
	grpc.ServerStream
}

type clientServiceGetAllClientsServer struct {
	grpc.ServerStream
}

func (x *clientServiceGetAllClientsServer) Send(m *ClientInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_GetClientLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/GetClientLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientLocation(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteClient(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientService_CreateClient_Handler,
		},
		{
			MethodName: "GetClientById",
			Handler:    _ClientService_GetClientById_Handler,
		},
		{
			MethodName: "GetClientLocation",
			Handler:    _ClientService_GetClientLocation_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ClientService_DeleteClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllClients",
			Handler:       _ClientService_GetAllClients_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/clientProto/client.proto",
}
