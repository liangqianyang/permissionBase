package permissionbaselogic

import (
	"context"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/model"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePermission 更新权限
func (l *UpdatePermissionLogic) UpdatePermission(in *permissionBase.UpdatePermissionRequest) (*permissionBase.UpdatePermissionResponse, error) {
	permissionModel := &model.Permissions{}
	l.svcCtx.Db.Model(&model.Permissions{}).Where("id = ?", in.Id).First(permissionModel)
	if permissionModel.Id == 0 {
		return nil, errors.New("permission not exists")
	}
	if in.Name != "" {
		permissionModel.Name = in.Name
	}
	if in.IdentifyName != "" {
		permissionModel.IdentifyName = in.IdentifyName
	}
	if in.Controller != "" {
		permissionModel.Controller = in.Controller
	}
	if in.Action != "" {
		permissionModel.Action = in.Action
	}
	if in.Description != "" {
		permissionModel.Description = in.Description
	}
	if in.State != 0 {
		permissionModel.State = int64(in.State)
	}
	if err := l.svcCtx.Db.Save(permissionModel).Error; err != nil {
		return nil, err
	}
	return &permissionBase.UpdatePermissionResponse{}, nil
}
