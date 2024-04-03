package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateMenu 创建菜单
func (l *CreateMenuLogic) CreateMenu(in *permissionBase.CreateMenuRequest) (*permissionBase.CreateMenuResponse, error) {
	newMenu := &model.Menus{
		ParentId: in.ParentId,
		Name:     in.Name,
		Icon:     in.Icon,
		Url:      in.Url,
		Sort:     in.Sort,
		MenuType: int64(in.MenuType),
	}

	if err := newMenu.CheckParentIdIsExists(l.ctx, l.svcCtx, in); err != nil {
		return nil, err
	}

	if err := newMenu.CheckMenuUnique(l.ctx, l.svcCtx, in); err != nil {
		return nil, err
	}

	if err := l.svcCtx.Db.Create(newMenu).Error; err != nil {
		logc.Errorf(l.ctx, "create menu failed: %v", err)
		return nil, err
	}

	return &permissionBase.CreateMenuResponse{}, nil
}
