package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"gorm.io/gorm"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleMenuLogic {
	return &SetRoleMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetRoleMenu 设置角色菜单
func (l *SetRoleMenuLogic) SetRoleMenu(in *permissionBase.SetRoleMenuRequest) (*permissionBase.SetRoleMenuResponse, error) {
	roleModel := &model.Roles{}
	if err := roleModel.CheckRoleIsExists(l.ctx, l.svcCtx, in.GetRoleId()); err != nil {
		return nil, err
	}
	items := make([]*model.RoleMenus, 0, len(in.MenuId))
	for _, id := range in.MenuId {
		items = append(items, &model.RoleMenus{
			RoleId: in.RoleId,
			MenuId: id,
		})
	}
	if err := l.svcCtx.Db.Transaction(func(tx *gorm.DB) error {
		// 删除原有的角色菜单
		if err := tx.Where(&model.RoleMenus{RoleId: in.GetRoleId()}).Delete(&model.RoleMenus{}).Error; err != nil {
			return err
		}
		// 创建新的角色菜单
		if err := tx.Create(items).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &permissionBase.SetRoleMenuResponse{}, nil
}
