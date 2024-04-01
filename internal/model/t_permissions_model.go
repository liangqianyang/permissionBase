package model

import (
	"context"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/svc"
	"time"
)

type Permissions struct {
	Id           int64     `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`                                  // 主键
	Name         string    `gorm:"column:name;type:varchar(32);not null;default:''" json:"name"`                                       // 权限名称
	IdentifyName string    `gorm:"column:identify_name;type:varchar(32);not null;default:''" json:"identify_name"`                     // 权限标识
	Controller   string    `gorm:"column:controller;type:varchar(32);not null;default:''" json:"controller"`                           // 控制器
	Action       string    `gorm:"column:action;type:varchar(32);not null;default:''" json:"action"`                                   // 方法
	Description  string    `gorm:"column:description;type:varchar(128);not null;default:''" json:"description"`                        // 权限描述
	State        int64     `gorm:"column:state;type:tinyint;not null;default:1" json:"state"`                                          // 状态 1:正常 -1:禁用
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 系统自动更新时间
}

// CheckPermissionUnique 检查权限是否唯一
func (m *Permissions) CheckPermissionUnique(ctx context.Context, svcCtx *svc.ServiceContext, identifyName string) error {
	var count int64
	svcCtx.Db.WithContext(ctx).Model(&Permissions{}).Where(&Permissions{IdentifyName: identifyName}).Count(&count)
	if count > 0 {
		return errors.New("permission already exists")
	}
	return nil
}
