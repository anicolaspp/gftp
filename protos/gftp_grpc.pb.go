// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

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

// GftpClient is the client API for Gftp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GftpClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	List(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirResponse, error)
	Get(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileContent, error)
	Put(ctx context.Context, in *FilePutRequest, opts ...grpc.CallOption) (*DirResponse, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*DirResponse, error)
	Move(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirResponse, error)
	Remove(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*DirResponse, error)
}

type gftpClient struct {
	cc grpc.ClientConnInterface
}

func NewGftpClient(cc grpc.ClientConnInterface) GftpClient {
	return &gftpClient{cc}
}

func (c *gftpClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gftpClient) List(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirResponse, error) {
	out := new(DirResponse)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gftpClient) Get(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileContent, error) {
	out := new(FileContent)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gftpClient) Put(ctx context.Context, in *FilePutRequest, opts ...grpc.CallOption) (*DirResponse, error) {
	out := new(DirResponse)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gftpClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*DirResponse, error) {
	out := new(DirResponse)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gftpClient) Move(ctx context.Context, in *DirRequest, opts ...grpc.CallOption) (*DirResponse, error) {
	out := new(DirResponse)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gftpClient) Remove(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*DirResponse, error) {
	out := new(DirResponse)
	err := c.cc.Invoke(ctx, "/gftp.Gftp/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GftpServer is the server API for Gftp service.
// All implementations must embed UnimplementedGftpServer
// for forward compatibility
type GftpServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	List(context.Context, *DirRequest) (*DirResponse, error)
	Get(context.Context, *FileRequest) (*FileContent, error)
	Put(context.Context, *FilePutRequest) (*DirResponse, error)
	Login(context.Context, *User) (*DirResponse, error)
	Move(context.Context, *DirRequest) (*DirResponse, error)
	Remove(context.Context, *FileRequest) (*DirResponse, error)
	mustEmbedUnimplementedGftpServer()
}

// UnimplementedGftpServer must be embedded to have forward compatible implementations.
type UnimplementedGftpServer struct {
}

func (UnimplementedGftpServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedGftpServer) List(context.Context, *DirRequest) (*DirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGftpServer) Get(context.Context, *FileRequest) (*FileContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGftpServer) Put(context.Context, *FilePutRequest) (*DirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedGftpServer) Login(context.Context, *User) (*DirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGftpServer) Move(context.Context, *DirRequest) (*DirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedGftpServer) Remove(context.Context, *FileRequest) (*DirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedGftpServer) mustEmbedUnimplementedGftpServer() {}

// UnsafeGftpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GftpServer will
// result in compilation errors.
type UnsafeGftpServer interface {
	mustEmbedUnimplementedGftpServer()
}

func RegisterGftpServer(s grpc.ServiceRegistrar, srv GftpServer) {
	s.RegisterService(&Gftp_ServiceDesc, srv)
}

func _Gftp_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gftp_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).List(ctx, req.(*DirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gftp_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).Get(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gftp_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).Put(ctx, req.(*FilePutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gftp_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gftp_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).Move(ctx, req.(*DirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gftp_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GftpServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gftp.Gftp/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GftpServer).Remove(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gftp_ServiceDesc is the grpc.ServiceDesc for Gftp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gftp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gftp.Gftp",
	HandlerType: (*GftpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Gftp_Status_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Gftp_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Gftp_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Gftp_Put_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Gftp_Login_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _Gftp_Move_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Gftp_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/gftp.proto",
}
