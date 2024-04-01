package model

import (
	"context"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"
	"time"
)

type Menus struct {
	Id        int64     `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`                                                 // 主键
	ParentId  int64     `gorm:"column:parent_id;type:bigint;not null;default:0" json:"parent_id"`                                                  // 父级菜单id
	Name      string    `gorm:"column:name;type:varchar(32);not null;default:''" json:"name"`                                                      // 菜单名称
	Icon      string    `gorm:"column:icon;type:varchar(32);not null;default:''" json:"icon"`                                                      // 菜单图标
	Path      string    `gorm:"column:path;type:varchar(128);not null;default:''" json:"path"`                                                     // 菜单地址
	MenuType  int64     `gorm:"column:menu_type;type:tinyint;not null;default:1" json:"menu_type"`                                                 // 菜单类型 1:目录 2:菜单 3:按钮
	Sort      int64     `gorm:"column:sort;type:int;not null;default:0" json:"sort"`                                                               // 排序
	State     int64     `gorm:"column:state;type:tinyint;not null;default:1" json:"state"`                                                         // 状态 1:正常 -1:禁用
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 系统自动更新时间
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
