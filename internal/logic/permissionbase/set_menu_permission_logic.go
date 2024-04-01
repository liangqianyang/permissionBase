package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMenuPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetMenuPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenuPermissionLogic {
	return &SetMenuPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetMenuPermission 设置菜单权限
func (l *SetMenuPermissionLogic) SetMenuPermission(in *permissionBase.SetMenuPermissionRequest) (*permissionBase.SetMenuPermissionResponse, error) {
	menuModel := &model.Menus{}
	if err := menuModel.CheckMenuIsExists(l.ctx, l.svcCtx, in.GetMenuId()); err != nil {
		return nil, err
	}
	items := make([]*model.MenuPermissions, 0, len(in.PermissionId))
	for _, id := range in.PermissionId {
		items = append(items, &model.MenuPermissions{
			MenuId:       in.MenuId,
			PermissionId: id,
		})
	}
	if err := l.svcCtx.Db.Transaction(func(tx *gorm.DB) error {
		// 删除原有的菜单权限
		if err := tx.Where(&model.MenuPermissions{MenuId: in.GetMenuId()}).Delete(&model.MenuPermissions{}).Error; err != nil {
			return err
		}
		// 创建新的菜单权限
		if err := tx.Create(items).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &permissionBase.SetMenuPermissionResponse{}, nil
}
