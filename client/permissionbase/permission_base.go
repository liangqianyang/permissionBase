// Code generated by goctl. DO NOT EDIT!
// Source: permissionBase.proto

package permissionbase

import (
	"context"

	"github.com/qy-blog/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePermissionRequest  = permissionBase.CreatePermissionRequest
	CreatePermissionResponse = permissionBase.CreatePermissionResponse
	CreateUserRequest        = permissionBase.CreateUserRequest
	CreateUserResponse       = permissionBase.CreateUserResponse
	LoginRequest             = permissionBase.LoginRequest
	LoginResponse            = permissionBase.LoginResponse

	PermissionBase interface {
		// CreatePermission 创建权限
		CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
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

// CreatePermission 创建权限
func (m *defaultPermissionBase) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	client := permissionBase.NewPermissionBaseClient(m.cli.Conn())
	return client.CreatePermission(ctx, in, opts...)
}
