// Code generated by goctl. DO NOT EDIT!
// Source: permissionBase.proto

package server

import (
	"context"
	userbaselogic "github.com/liangqianyang/permissionBase/internal/logic/userbase"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"
)

type UserBaseServer struct {
	svcCtx *svc.ServiceContext
	permissionBase.UnimplementedUserBaseServer
}

func NewUserBaseServer(svcCtx *svc.ServiceContext) *UserBaseServer {
	return &UserBaseServer{
		svcCtx: svcCtx,
	}
}

// CreateUser 创建用户
func (s *UserBaseServer) CreateUser(ctx context.Context, in *permissionBase.CreateUserRequest) (*permissionBase.CreateUserResponse, error) {
	l := userbaselogic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

// Login 登录
func (s *UserBaseServer) Login(ctx context.Context, in *permissionBase.LoginRequest) (*permissionBase.LoginResponse, error) {
	l := userbaselogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
