package model

import (
	"context"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"time"
)

type Roles struct {
	Id          int64     `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`                                                 // 主键
	Name        string    `gorm:"column:name;type:varchar(32);not null;default:''" json:"name"`                                                      // 角色名称
	Description string    `gorm:"column:description;type:varchar(128);not null;default:''" json:"description"`                                       // 角色描述
	State       int64     `gorm:"column:state;type:tinyint;not null;default:1" json:"state"`                                                         // 状态 1:正常 -1:禁用
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 系统自动更新时间
}

// CheckRoleUnique 检查角色是否唯一
func (m *Roles) CheckRoleUnique(ctx context.Context, svcCtx *svc.ServiceContext, name string) error {
	var count int64
	svcCtx.Db.WithContext(ctx).Model(&Roles{}).Where(&Roles{Name: name}).Count(&count)
	if count > 0 {
		return errors.New("role already exists")
	}
	return nil
}

// CheckRoleIsExists 检查角色是否存在
func (m *Roles) CheckRoleIsExists(ctx context.Context, svcCtx *svc.ServiceContext, id int64) error {
	var count int64
	svcCtx.Db.WithContext(ctx).Model(&Roles{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return errors.New("role not exists")
	}
	return nil
}
