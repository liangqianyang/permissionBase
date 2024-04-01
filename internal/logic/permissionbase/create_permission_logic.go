package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"github.com/zeromicro/go-zero/core/logc"

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

// CreatePermission 创建权限
func (l *CreatePermissionLogic) CreatePermission(in *permissionBase.CreatePermissionRequest) (*permissionBase.CreatePermissionResponse, error) {
	newPermission := &model.Permissions{
		Name:         in.Name,
		IdentifyName: in.IdentifyName,
		Controller:   in.Controller,
		Action:       in.Action,
		Description:  in.Description,
	}

	if err := newPermission.CheckPermissionUnique(l.ctx, l.svcCtx, in.IdentifyName); err != nil {
		return nil, err
	}

	if err := l.svcCtx.Db.Create(newPermission).Error; err != nil {
		logc.Errorf(l.ctx, "create permission failed: %v", err)
		return nil, err
	}

	return &permissionBase.CreatePermissionResponse{}, nil
}
