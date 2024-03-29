// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.5
// source: permissionBase.proto

package permissionBase

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

const (
	PermissionBase_CreateMenu_FullMethodName       = "/permissionBase.PermissionBase/CreateMenu"
	PermissionBase_CreatePermission_FullMethodName = "/permissionBase.PermissionBase/CreatePermission"
	PermissionBase_UpdateMenu_FullMethodName       = "/permissionBase.PermissionBase/UpdateMenu"
	PermissionBase_CreateRole_FullMethodName       = "/permissionBase.PermissionBase/CreateRole"
)

// PermissionBaseClient is the client API for PermissionBase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionBaseClient interface {
	// CreateMenu 创建菜单
	CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error)
	// CreatePermission 创建权限
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
	// UpdateMenu 更新菜单
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error)
	// CreateRole 创建角色
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
}

type permissionBaseClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionBaseClient(cc grpc.ClientConnInterface) PermissionBaseClient {
	return &permissionBaseClient{cc}
}

func (c *permissionBaseClient) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error) {
	out := new(CreateMenuResponse)
	err := c.cc.Invoke(ctx, PermissionBase_CreateMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionBaseClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	out := new(CreatePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionBase_CreatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionBaseClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error) {
	out := new(UpdateMenuResponse)
	err := c.cc.Invoke(ctx, PermissionBase_UpdateMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionBaseClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, PermissionBase_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionBaseServer is the server API for PermissionBase service.
// All implementations must embed UnimplementedPermissionBaseServer
// for forward compatibility
type PermissionBaseServer interface {
	// CreateMenu 创建菜单
	CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuResponse, error)
	// CreatePermission 创建权限
	CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error)
	// UpdateMenu 更新菜单
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error)
	// CreateRole 创建角色
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	mustEmbedUnimplementedPermissionBaseServer()
}

// UnimplementedPermissionBaseServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionBaseServer struct {
}

func (UnimplementedPermissionBaseServer) CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedPermissionBaseServer) CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedPermissionBaseServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedPermissionBaseServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedPermissionBaseServer) mustEmbedUnimplementedPermissionBaseServer() {}

// UnsafePermissionBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionBaseServer will
// result in compilation errors.
type UnsafePermissionBaseServer interface {
	mustEmbedUnimplementedPermissionBaseServer()
}

func RegisterPermissionBaseServer(s grpc.ServiceRegistrar, srv PermissionBaseServer) {
	s.RegisterService(&PermissionBase_ServiceDesc, srv)
}

func _PermissionBase_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionBaseServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionBase_CreateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionBaseServer).CreateMenu(ctx, req.(*CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionBase_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionBaseServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionBase_CreatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionBaseServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionBase_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionBaseServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionBase_UpdateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionBaseServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionBase_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionBaseServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionBase_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionBaseServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionBase_ServiceDesc is the grpc.ServiceDesc for PermissionBase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionBase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permissionBase.PermissionBase",
	HandlerType: (*PermissionBaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMenu",
			Handler:    _PermissionBase_CreateMenu_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _PermissionBase_CreatePermission_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _PermissionBase_UpdateMenu_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _PermissionBase_CreateRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permissionBase.proto",
}

const (
	UserBase_CreateUser_FullMethodName = "/permissionBase.UserBase/CreateUser"
	UserBase_Login_FullMethodName      = "/permissionBase.UserBase/Login"
)

// UserBaseClient is the client API for UserBase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserBaseClient interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Login 登录
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userBaseClient struct {
	cc grpc.ClientConnInterface
}

func NewUserBaseClient(cc grpc.ClientConnInterface) UserBaseClient {
	return &userBaseClient{cc}
}

func (c *userBaseClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserBase_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userBaseClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserBase_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserBaseServer is the server API for UserBase service.
// All implementations must embed UnimplementedUserBaseServer
// for forward compatibility
type UserBaseServer interface {
	// CreateUser 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedUserBaseServer()
}

// UnimplementedUserBaseServer must be embedded to have forward compatible implementations.
type UnimplementedUserBaseServer struct {
}

func (UnimplementedUserBaseServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserBaseServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserBaseServer) mustEmbedUnimplementedUserBaseServer() {}

// UnsafeUserBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserBaseServer will
// result in compilation errors.
type UnsafeUserBaseServer interface {
	mustEmbedUnimplementedUserBaseServer()
}

func RegisterUserBaseServer(s grpc.ServiceRegistrar, srv UserBaseServer) {
	s.RegisterService(&UserBase_ServiceDesc, srv)
}

func _UserBase_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBaseServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBase_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBaseServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserBase_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBaseServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBase_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBaseServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserBase_ServiceDesc is the grpc.ServiceDesc for UserBase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserBase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permissionBase.UserBase",
	HandlerType: (*UserBaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserBase_CreateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserBase_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permissionBase.proto",
}
