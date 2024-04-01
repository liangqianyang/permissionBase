package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRole 更新角色
func (l *UpdateRoleLogic) UpdateRole(in *permissionBase.UpdateRoleRequest) (*permissionBase.UpdateRoleResponse, error) {
	roleModel := &model.Roles{}
	if err := roleModel.CheckRoleIsExists(l.ctx, l.svcCtx, in.Id); err != nil {
		return nil, err
	}
	if in.Name != "" {
		roleModel.Name = in.Name
	}
	if in.Description != "" {
		roleModel.Description = in.Description
	}
	if in.State != 0 {
		roleModel.State = int64(in.State)
	}
	if err := l.svcCtx.Db.Save(roleModel).Error; err != nil {
		return nil, err
	}

	return &permissionBase.UpdateRoleResponse{}, nil
}
