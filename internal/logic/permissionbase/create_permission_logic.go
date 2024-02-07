package permissionbaselogic

import (
	"context"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePermissionLogic) CreatePermission(in *permissionBase.CreatePermissionRequest) (*permissionBase.CreatePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &permissionBase.CreatePermissionResponse{}, nil
}
