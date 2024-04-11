package permissionbaselogic

import (
	"context"
	"fmt"
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
	db = db.Order("sort asc")
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}
	list := make([]*permissionBase.MenuInfo, 0)
	for _, menu := range menus {
		list = append(list, &permissionBase.MenuInfo{
			ParentId:  menu.ParentId,
			Id:        menu.Id,
			Name:      menu.Name,
			Component: menu.Component,
			Path:      menu.Path,
			Title:     menu.Name,
			MenuType:  permissionBase.MenuType(menu.MenuType),
			State:     permissionBase.MenuState(menu.State),
			Meta: &permissionBase.MenuMeta{
				Title:      menu.Name,
				Icon:       menu.Icon,
				Hidden:     menu.Hidden,
				AlwaysShow: menu.AlwaysShow,
				NoCache:    menu.NoCache,
				Breadcrumb: menu.Breadcrumb,
				Affix:      menu.Affix,
				NoTagsView: menu.NoTagsView,
			},
		})
	}
	menuTree := l.buildMenuTree(list)
	return &permissionBase.GetMenuListResponse{
		List: menuTree,
	}, nil
}

// buildMenuTree 构建菜单树
func (l *GetMenuListLogic) buildMenuTree(items []*permissionBase.MenuInfo) []*permissionBase.MenuInfo {
	list := make([]*permissionBase.MenuInfo, 0)
	for i := range items {
		if items[i].ParentId == 0 {
			list = append(list, items[i])
		} else {
			for j := range items {
				if items[j].Id == items[i].ParentId {
					if items[j].Children == nil {
						items[j].Children = make([]*permissionBase.MenuInfo, 0)
					}
					items[j].Children = append(items[j].Children, items[i])
					if items[j].Redirect == "" {
						items[j].Redirect = fmt.Sprintf("%s/%s", items[j].Path, items[i].Path)
					}
				}
			}
		}
	}
	return list
}
