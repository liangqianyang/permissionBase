package model

import "time"

type RoleMenus struct {
	Id        int64     `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`                                 // 主键
	RoleId    int       `gorm:"column:role_id;type:int;not null;default:0" json:"role_id"`                                         // 角色ID
	MenuId    int       `gorm:"column:menu_id;type:int;not null" json:"menu_id"`                                                   // 菜单ID
	CreatedAt time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 系统自动更新时间
}
