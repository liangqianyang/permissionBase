// Code generated by goctl. DO NOT EDIT!
// Source: permissionBase.proto

package server

import (
	"context"

	"github.com/liangqianyang/permissionBase/internal/logic/permissionbase"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"
)

type PermissionBaseServer struct {
	svcCtx *svc.ServiceContext
	permissionBase.UnimplementedPermissionBaseServer
}

func NewPermissionBaseServer(svcCtx *svc.ServiceContext) *PermissionBaseServer {
	return &PermissionBaseServer{
		svcCtx: svcCtx,
	}
}

// CreateMenu 创建菜单
func (s *PermissionBaseServer) CreateMenu(ctx context.Context, in *permissionBase.CreateMenuRequest) (*permissionBase.CreateMenuResponse, error) {
	l := permissionbaselogic.NewCreateMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(in)
}

// UpdateMenu 更新菜单
func (s *PermissionBaseServer) UpdateMenu(ctx context.Context, in *permissionBase.UpdateMenuRequest) (*permissionBase.UpdateMenuResponse, error) {
	l := permissionbaselogic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

// CreatePermission 创建权限
func (s *PermissionBaseServer) CreatePermission(ctx context.Context, in *permissionBase.CreatePermissionRequest) (*permissionBase.CreatePermissionResponse, error) {
	l := permissionbaselogic.NewCreatePermissionLogic(ctx, s.svcCtx)
	return l.CreatePermission(in)
}

// UpdatePermission 更新权限
func (s *PermissionBaseServer) UpdatePermission(ctx context.Context, in *permissionBase.UpdatePermissionRequest) (*permissionBase.UpdatePermissionResponse, error) {
	l := permissionbaselogic.NewUpdatePermissionLogic(ctx, s.svcCtx)
	return l.UpdatePermission(in)
}

// CreateRole 创建角色
func (s *PermissionBaseServer) CreateRole(ctx context.Context, in *permissionBase.CreateRoleRequest) (*permissionBase.CreateRoleResponse, error) {
	l := permissionbaselogic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

// UpdateRole 更新角色
func (s *PermissionBaseServer) UpdateRole(ctx context.Context, in *permissionBase.UpdateRoleRequest) (*permissionBase.UpdateRoleResponse, error) {
	l := permissionbaselogic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

// SetMenuPermission 设置菜单权限
func (s *PermissionBaseServer) SetMenuPermission(ctx context.Context, in *permissionBase.SetMenuPermissionRequest) (*permissionBase.SetMenuPermissionResponse, error) {
	l := permissionbaselogic.NewSetMenuPermissionLogic(ctx, s.svcCtx)
	return l.SetMenuPermission(in)
}

// SetRoleMenu 设置角色菜单
func (s *PermissionBaseServer) SetRoleMenu(ctx context.Context, in *permissionBase.SetRoleMenuRequest) (*permissionBase.SetRoleMenuResponse, error) {
	l := permissionbaselogic.NewSetRoleMenuLogic(ctx, s.svcCtx)
	return l.SetRoleMenu(in)
}

// SetUserRole 设置用户角色
func (s *PermissionBaseServer) SetUserRole(ctx context.Context, in *permissionBase.SetUserRoleRequest) (*permissionBase.SetUserRoleResponse, error) {
	l := permissionbaselogic.NewSetUserRoleLogic(ctx, s.svcCtx)
	return l.SetUserRole(in)
}
