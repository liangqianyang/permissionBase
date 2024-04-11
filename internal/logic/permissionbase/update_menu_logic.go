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
	if in.ParentId > 0 {
		var count int64
		l.svcCtx.Db.Model(&model.Menus{}).Where("id = ?", in.ParentId).Count(&count)
		if count == 0 {
			return nil, errors.New("parent menu not exists")
		}
		menu.ParentId = in.ParentId
	}
	if in.Name != "" {
		var count int64
		l.svcCtx.Db.Model(&model.Menus{}).Where("name = ?", in.Name).Count(&count)
		if count > 0 {
			return nil, errors.New("menu already exists")
		}
		menu.Name = in.Name
	}
	if in.Icon != "" {
		menu.Icon = in.Icon
	}
	if in.Path != "" {
		menu.Path = in.Path
	}
	if in.MenuType > 0 {
		menu.MenuType = int64(in.MenuType)
	}
	if in.Sort > 0 {
		menu.Sort = in.Sort
	}
	if in.State != 0 {
		menu.State = int64(in.State)
	}
	if in.Hidden != 0 {
		menu.Hidden = in.Hidden
	}
	if in.AlwaysShow != 0 {
		menu.AlwaysShow = in.AlwaysShow
	}
	if in.NoCache != 0 {
		menu.NoCache = in.NoCache
	}
	if in.Breadcrumb != 0 {
		menu.Breadcrumb = in.Breadcrumb
	}
	if in.Affix != 0 {
		menu.Affix = in.Affix
	}
	if in.NoTagsView != 0 {
		menu.NoTagsView = in.NoTagsView
	}
	if err := l.svcCtx.Db.Save(menu).Error; err != nil {
		logc.Errorf(l.ctx, "update menu failed: %v", err)
		return nil, err
	}
	return &permissionBase.UpdateMenuResponse{}, nil
}
