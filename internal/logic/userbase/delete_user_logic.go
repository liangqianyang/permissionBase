package userbaselogic

import (
	"context"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteUser 删除用户
func (l *DeleteUserLogic) DeleteUser(in *permissionBase.DeleteUserRequest) (*permissionBase.DeleteUserResponse, error) {
	// todo: add your logic here and delete this line

	return &permissionBase.DeleteUserResponse{}, nil
}
