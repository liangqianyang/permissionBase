package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateRole 创建角色
func (l *CreateRoleLogic) CreateRole(in *permissionBase.CreateRoleRequest) (*permissionBase.CreateRoleResponse, error) {
	newRole := &model.Roles{
		Name:        in.Name,
		Description: in.Description,
	}

	err := newRole.CheckRoleUnique(l.ctx, l.svcCtx, in.Name)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Db.Create(newRole).Error
	if err != nil {
		logc.Errorf(l.ctx, "create role failed: %v", err)
		return nil, err
	}

	return &permissionBase.CreateRoleResponse{}, nil
}
