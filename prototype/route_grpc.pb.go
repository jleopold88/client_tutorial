// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: prototype/route.proto

package prototype

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

// TutorialClient is the client API for Tutorial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TutorialClient interface {
	CreateUser(ctx context.Context, in *CreateData, opts ...grpc.CallOption) (*SuccessMessage, error)
	GetByID(ctx context.Context, in *GetDatabyID, opts ...grpc.CallOption) (*DataReturn, error)
	GetByName(ctx context.Context, in *GetDatabyName, opts ...grpc.CallOption) (*ListDataReturn, error)
	UpdateUser(ctx context.Context, in *UpdateData, opts ...grpc.CallOption) (*SuccessMessage, error)
	DeleteUser(ctx context.Context, in *DeleteData, opts ...grpc.CallOption) (*SuccessMessage, error)
}

type tutorialClient struct {
	cc grpc.ClientConnInterface
}

func NewTutorialClient(cc grpc.ClientConnInterface) TutorialClient {
	return &tutorialClient{cc}
}

func (c *tutorialClient) CreateUser(ctx context.Context, in *CreateData, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/prototype.Tutorial/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialClient) GetByID(ctx context.Context, in *GetDatabyID, opts ...grpc.CallOption) (*DataReturn, error) {
	out := new(DataReturn)
	err := c.cc.Invoke(ctx, "/prototype.Tutorial/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialClient) GetByName(ctx context.Context, in *GetDatabyName, opts ...grpc.CallOption) (*ListDataReturn, error) {
	out := new(ListDataReturn)
	err := c.cc.Invoke(ctx, "/prototype.Tutorial/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialClient) UpdateUser(ctx context.Context, in *UpdateData, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/prototype.Tutorial/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialClient) DeleteUser(ctx context.Context, in *DeleteData, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/prototype.Tutorial/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TutorialServer is the server API for Tutorial service.
// All implementations must embed UnimplementedTutorialServer
// for forward compatibility
type TutorialServer interface {
	CreateUser(context.Context, *CreateData) (*SuccessMessage, error)
	GetByID(context.Context, *GetDatabyID) (*DataReturn, error)
	GetByName(context.Context, *GetDatabyName) (*ListDataReturn, error)
	UpdateUser(context.Context, *UpdateData) (*SuccessMessage, error)
	DeleteUser(context.Context, *DeleteData) (*SuccessMessage, error)
	mustEmbedUnimplementedTutorialServer()
}

// UnimplementedTutorialServer must be embedded to have forward compatible implementations.
type UnimplementedTutorialServer struct {
}

func (UnimplementedTutorialServer) CreateUser(context.Context, *CreateData) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedTutorialServer) GetByID(context.Context, *GetDatabyID) (*DataReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedTutorialServer) GetByName(context.Context, *GetDatabyName) (*ListDataReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedTutorialServer) UpdateUser(context.Context, *UpdateData) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedTutorialServer) DeleteUser(context.Context, *DeleteData) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedTutorialServer) mustEmbedUnimplementedTutorialServer() {}

// UnsafeTutorialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TutorialServer will
// result in compilation errors.
type UnsafeTutorialServer interface {
	mustEmbedUnimplementedTutorialServer()
}

func RegisterTutorialServer(s grpc.ServiceRegistrar, srv TutorialServer) {
	s.RegisterService(&Tutorial_ServiceDesc, srv)
}

func _Tutorial_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototype.Tutorial/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).CreateUser(ctx, req.(*CreateData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tutorial_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototype.Tutorial/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).GetByID(ctx, req.(*GetDatabyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tutorial_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabyName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototype.Tutorial/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).GetByName(ctx, req.(*GetDatabyName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tutorial_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototype.Tutorial/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).UpdateUser(ctx, req.(*UpdateData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tutorial_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prototype.Tutorial/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).DeleteUser(ctx, req.(*DeleteData))
	}
	return interceptor(ctx, in, info, handler)
}

// Tutorial_ServiceDesc is the grpc.ServiceDesc for Tutorial service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tutorial_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prototype.Tutorial",
	HandlerType: (*TutorialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Tutorial_CreateUser_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Tutorial_GetByID_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _Tutorial_GetByName_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Tutorial_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Tutorial_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prototype/route.proto",
}
