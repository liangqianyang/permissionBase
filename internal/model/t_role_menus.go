package model

type RoleMenus struct {
	Id     int64 `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"` // 主键
	RoleId int64 `gorm:"column:role_id;type:bigint;not null;default:0" json:"role_id"`      // 角色ID
	MenuId int64 `gorm:"column:menu_id;type:bigint;not null" json:"menu_id"`                // 菜单ID
}
