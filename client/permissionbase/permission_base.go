// Code generated by goctl. DO NOT EDIT!
// Source: permissionBase.proto

package permissionbase

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
	LoginRequest              = permissionBase.LoginRequest
	LoginResponse             = permissionBase.LoginResponse
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

	PermissionBase interface {
		// CreateMenu 创建菜单
		CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error)
		// UpdateMenu 更新菜单
		UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error)
		// CreatePermission 创建权限
		CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
		// UpdatePermission 更新权限
		UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error)
		// CreateRole 创建角色
		CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
		// UpdateRole 更新角色
		UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
		// SetMenuPermission 设置菜单权限
		SetMenuPermission(ctx context.Context, in *SetMenuPermissionRequest, opts ...grpc.CallOption) (*SetMenuPermissionResponse, error)
		// SetRoleMenu 设置角色菜单
		SetRoleMenu(ctx context.Context, in *SetRoleMenuRequest, opts ...grpc.CallOption) (*SetRoleMenuResponse, error)
		// SetUserRole 设置用户角色
		SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error)
	}

	defaultPermissionBase struct {
		cli zrpc.Client
	}
)

func NewPermissionBase(cli zrpc.Client) PermissionBase {
	return &defaultPermissionBase{
		cli: cli,
	}
}

// CreateMenu 创建菜单
func (m *defaultPermissionBase) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.CreateMenu(ctx, in, opts...)
}

// UpdateMenu 更新菜单
func (m *defaultPermissionBase) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

// CreatePermission 创建权限
func (m *defaultPermissionBase) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.CreatePermission(ctx, in, opts...)
}

// UpdatePermission 更新权限
func (m *defaultPermissionBase) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.UpdatePermission(ctx, in, opts...)
}

// CreateRole 创建角色
func (m *defaultPermissionBase) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

// UpdateRole 更新角色
func (m *defaultPermissionBase) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

// SetMenuPermission 设置菜单权限
func (m *defaultPermissionBase) SetMenuPermission(ctx context.Context, in *SetMenuPermissionRequest, opts ...grpc.CallOption) (*SetMenuPermissionResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.SetMenuPermission(ctx, in, opts...)
}

// SetRoleMenu 设置角色菜单
func (m *defaultPermissionBase) SetRoleMenu(ctx context.Context, in *SetRoleMenuRequest, opts ...grpc.CallOption) (*SetRoleMenuResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.SetRoleMenu(ctx, in, opts...)
}

// SetUserRole 设置用户角色
func (m *defaultPermissionBase) SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.SetUserRole(ctx, in, opts...)
}
