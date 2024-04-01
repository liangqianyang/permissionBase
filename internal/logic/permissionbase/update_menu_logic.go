package permissionbaselogic

import (
	"context"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenu 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(in *permissionBase.UpdateMenuRequest) (*permissionBase.UpdateMenuResponse, error) {
	menu := &model.Menus{}
	if err := menu.CheckMenuIsExists(l.ctx, l.svcCtx, in.Id); err != nil {
		return nil, err
	}
	if in.ParentId != nil && in.ParentId.Value > 0 {
		var count int64
		l.svcCtx.Db.Model(&model.Menus{}).Where("id = ?", in.ParentId).Count(&count)
		if count == 0 {
			return nil, errors.New("parent menu not exists")
		}
		menu.ParentId = in.ParentId.Value
	}
	if in.Name != nil && in.Name.Value != "" {
		var count int64
		l.svcCtx.Db.Model(&model.Menus{}).Where("name = ?", in.Name.Value).Count(&count)
		if count > 0 {
			return nil, errors.New("menu already exists")
		}
		menu.Name = in.Name.Value
	}

	if in.Icon != nil && in.Icon.Value != "" {
		menu.Icon = in.Icon.Value
	}

	if in.Path != nil && in.Path.Value != "" {
		menu.Path = in.Path.Value
	}

	if in.MenuType != nil && in.MenuType.Value > 0 {
		menu.MenuType = in.MenuType.Value
	}

	if in.Sort != nil && in.Sort.Value > 0 {
		menu.Sort = in.Sort.Value
	}

	if in.State != nil && in.State.Value != 0 {
		menu.State = in.State.Value
	}

	if err := l.svcCtx.Db.Save(menu).Error; err != nil {
		logc.Errorf(l.ctx, "update menu failed: %v", err)
		return nil, err
	}

	return &permissionBase.UpdateMenuResponse{}, nil
}
