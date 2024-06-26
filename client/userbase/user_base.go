// Code generated by goctl. DO NOT EDIT!
// Source: permissionBase.proto

package userbase

import (
	"context"

	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateMenuRequest         = permissionBase.CreateMenuRequest
	CreateMenuResponse        = permissionBase.CreateMenuResponse
	CreatePermissionRequest   = permissionBase.CreatePermissionRequest
	CreatePermissionResponse  = permissionBase.CreatePermissionResponse
	CreateRoleRequest         = permissionBase.CreateRoleRequest
	CreateRoleResponse        = permissionBase.CreateRoleResponse
	CreateUserRequest         = permissionBase.CreateUserRequest
	CreateUserResponse        = permissionBase.CreateUserResponse
	DeleteUserRequest         = permissionBase.DeleteUserRequest
	DeleteUserResponse        = permissionBase.DeleteUserResponse
	GetMenuListRequest        = permissionBase.GetMenuListRequest
	GetMenuListResponse       = permissionBase.GetMenuListResponse
	GetPermissionListRequest  = permissionBase.GetPermissionListRequest
	GetPermissionListResponse = permissionBase.GetPermissionListResponse
	GetRoleListRequest        = permissionBase.GetRoleListRequest
	GetRoleListResponse       = permissionBase.GetRoleListResponse
	GetUserListRequest        = permissionBase.GetUserListRequest
	GetUserListResponse       = permissionBase.GetUserListResponse
	LoginRequest              = permissionBase.LoginRequest
	LoginResponse             = permissionBase.LoginResponse
	MenuInfo                  = permissionBase.MenuInfo
	MenuMeta                  = permissionBase.MenuMeta
	PageInfo                  = permissionBase.PageInfo
	PermissionInfo            = permissionBase.PermissionInfo
	RoleInfo                  = permissionBase.RoleInfo
	SetMenuPermissionRequest  = permissionBase.SetMenuPermissionRequest
	SetMenuPermissionResponse = permissionBase.SetMenuPermissionResponse
	SetRoleMenuRequest        = permissionBase.SetRoleMenuRequest
	SetRoleMenuResponse       = permissionBase.SetRoleMenuResponse
	SetUserRoleRequest        = permissionBase.SetUserRoleRequest
	SetUserRoleResponse       = permissionBase.SetUserRoleResponse
	UpdateMenuRequest         = permissionBase.UpdateMenuRequest
	UpdateMenuResponse        = permissionBase.UpdateMenuResponse
	UpdatePermissionRequest   = permissionBase.UpdatePermissionRequest
	UpdatePermissionResponse  = permissionBase.UpdatePermissionResponse
	UpdateRoleRequest         = permissionBase.UpdateRoleRequest
	UpdateRoleResponse        = permissionBase.UpdateRoleResponse
	UpdateUserRequest         = permissionBase.UpdateUserRequest
	UpdateUserResponse        = permissionBase.UpdateUserResponse
	UserInfo                  = permissionBase.UserInfo

	UserBase interface {
		// CreateUser 创建用户
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		// UpdateUser 更新用户
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
		// DeleteUser 删除用户
		DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
		// Login 登录
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	}

	defaultUserBase struct {
		cli zrpc.Client
	}
)

func NewUserBase(cli zrpc.Client) UserBase {
	return &defaultUserBase{
		cli: cli,
	}
}

// CreateUser 创建用户
func (m *defaultUserBase) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := permissionBase.NewUserBaseClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

// UpdateUser 更新用户
func (m *defaultUserBase) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	client := permissionBase.NewUserBaseClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

// DeleteUser 删除用户
func (m *defaultUserBase) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	client := permissionBase.NewUserBaseClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

// Login 登录
func (m *defaultUserBase) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := permissionBase.NewUserBaseClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
