package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"gorm.io/gorm"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserRoleLogic {
	return &SetUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetUserRole 设置用户角色
func (l *SetUserRoleLogic) SetUserRole(in *permissionBase.SetUserRoleRequest) (*permissionBase.SetUserRoleResponse, error) {
	roleModel := &model.Roles{}
	if err := roleModel.CheckRoleIsExists(l.ctx, l.svcCtx, in.GetRoleId()); err != nil {
		return nil, err
	}
	items := make([]*model.RoleUsers, 0, len(in.UserId))
	for _, id := range in.UserId {
		items = append(items, &model.RoleUsers{
			RoleId: in.RoleId,
			UserId: id,
		})
	}
	if err := l.svcCtx.Db.Transaction(func(tx *gorm.DB) error {
		// 删除原有的用户角色
		if err := tx.Where(&model.RoleUsers{RoleId: in.GetRoleId()}).Delete(&model.RoleUsers{}).Error; err != nil {
			return err
		}
		// 创建新的用户角色
		if err := tx.Create(items).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &permissionBase.SetUserRoleResponse{}, nil
}
