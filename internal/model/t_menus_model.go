package model

import (
	"context"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"
	"time"
)

type Menus struct {
	Id            int64     `json:"id" gorm:"primaryKey;column:id;type:bigint;autoIncrement;comment:主键"`
	ParentId      int64     `json:"parentId" gorm:"column:parent_id;type:bigint;not null;default:0;comment:父级菜单id"`
	MenuType      int64     `json:"menuType" gorm:"column:menu_type;type:tinyint;not null;default:1;comment:菜单类型 1:目录 2:菜单"`
	Name          string    `json:"name" gorm:"column:name;type:varchar(32);not null;default:'';comment:菜单名称"`
	Icon          string    `json:"icon" gorm:"column:icon;type:varchar(32);not null;default:'';comment:菜单图标"`
	Component     string    `json:"component" gorm:"column:component;type:varchar(128);not null;default:'';comment:菜单组件"`
	ComponentName string    `json:"componentName" gorm:"column:component_name;type:varchar(32);not null;default:'';comment:菜单组件名称"`
	Path          string    `json:"path" gorm:"column:path;type:varchar(128);not null;default:'';comment:菜单路径"`
	Sort          int64     `json:"sort" gorm:"column:sort;type:int;not null;default:0;comment:排序"`
	State         int64     `json:"state" gorm:"column:state;type:tinyint;not null;default:1;comment:状态 1:正常 -1:禁用"`
	Hidden        int64     `json:"hidden" gorm:"column:hidden;type:tinyint;not null;default:0;comment:是否隐藏 0:显示 1:隐藏"`
	AlwaysShow    int64     `json:"alwaysShow" gorm:"column:always_show;type:tinyint;not null;default:0;comment:是否一直显示 0:不显示 1:显示"`
	NoCache       int64     `json:"noCache" gorm:"column:no_cache;type:tinyint;not null;default:0;comment:是否缓存 0:缓存 1:不缓存"`
	Breadcrumb    int64     `json:"breadcrumb" gorm:"column:breadcrumb;type:tinyint;not null;default:0;comment:是否显示面包屑 0:显示 1:不显示"`
	Affix         int64     `json:"affix" gorm:"column:affix;type:tinyint;not null;default:0;comment:是否固定在tagsView 0:不固定 1:固定"`
	NoTagsView    int64     `json:"noTagsView" gorm:"column:no_tags_view;type:tinyint;not null;default:0;comment:是否在tagsView显示 0:显示 1:不显示"`
	OperatorId    int64     `json:"operatorId" gorm:"column:operator_id;type:bigint;not null;default:0;comment:操作人id"`
	Operator      string    `json:"operator" gorm:"column:operator;type:varchar(32);not null;default:'';comment:操作人名称"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
}

// CheckParentIdIsExists 检查父级菜单是否存在
func (m *Menus) CheckParentIdIsExists(ctx context.Context, svcCtx *svc.ServiceContext, in *permissionBase.CreateMenuRequest) error {
	if in.ParentId > 0 {
		var count int64
		svcCtx.Db.WithContext(ctx).Model(&Menus{}).Where("id = ?", in.ParentId).Count(&count)
		if count == 0 {
			return errors.New("parent menu not exists")
		}
	}
	return nil
}

// CheckMenuUnique 检查菜单是否唯一
func (m *Menus) CheckMenuUnique(ctx context.Context, svcCtx *svc.ServiceContext, in *permissionBase.CreateMenuRequest) error {
	var count int64
	svcCtx.Db.WithContext(ctx).Model(&Menus{}).Where(&Menus{Name: in.Name}).Count(&count)
	if count > 0 {
		return errors.New("menu already exists")
	}
	return nil
}

// CheckMenuIsExists 检查菜单是否存在
func (m *Menus) CheckMenuIsExists(ctx context.Context, svcCtx *svc.ServiceContext, id int64) error {
	var count int64
	svcCtx.Db.WithContext(ctx).Model(&Menus{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return errors.New("menu not exists")
	}
	return nil
}
