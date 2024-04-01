package model

import "time"

type MenuPermissions struct {
	Id           int64     `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`                                 // 主键
	MenuId       int64     `gorm:"column:menu_id;type:bigint;not null;default:0" json:"menu_id"`                                      // 菜单id
	PermissionId int64     `gorm:"column:permission_id;type:bigint;not null;default:0" json:"permission_id"`                          // 权限id
	CreatedAt    time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt    time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 系统自动更新时间
}
