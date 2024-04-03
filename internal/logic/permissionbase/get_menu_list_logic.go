package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMenuList 获取菜单列表
func (l *GetMenuListLogic) GetMenuList(in *permissionBase.GetMenuListRequest) (*permissionBase.GetMenuListResponse, error) {
	var menus []model.Menus
	var total int64
	db := l.svcCtx.Db
	if in.ParentId != 0 {
		db = db.Where("parent_id = ?", in.ParentId)
	}
	if in.Name != "" {
		db = db.Where("name like ?", in.Name+"%")
	}
	if in.MenuType != 0 {
		db = db.Where("menu_type = ?", in.MenuType)
	}
	if in.State != 0 {
		db = db.Where("status = ?", in.State)
	}
	if err := db.Model(&model.Menus{}).Count(&total).Error; err != nil {
		return nil, err
	}
	page := in.Page
	pageSize := in.PageSize
	if page <= 0 {
		page = model.DefaultPage
	}
	if pageSize <= 0 {
		pageSize = model.DefaultPageSize
	}
	db = db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Order("sort asc")
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}
	list := make([]*permissionBase.MenuInfo, 0)
	for _, menu := range menus {
		list = append(list, &permissionBase.MenuInfo{
			Id:         menu.Id,
			Name:       menu.Name,
			ParentId:   menu.ParentId,
			MenuType:   permissionBase.MenuType(menu.MenuType),
			Icon:       menu.Icon,
			Url:        menu.Url,
			Sort:       menu.Sort,
			State:      permissionBase.MenuState(menu.State),
			OperatorId: menu.OperatorId,
			Operator:   menu.Operator,
			CreatedAt:  menu.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  menu.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &permissionBase.GetMenuListResponse{
		List: list,
		Page: &permissionBase.PageInfo{
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	}, nil
}
