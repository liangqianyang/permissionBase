package model

type RoleUsers struct {
	Id     int64 `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"` // 主键
	RoleId int64 `gorm:"column:role_id;type:bigint;not null;default:0" json:"role_id"`      // 角色id
	UserId int64 `gorm:"column:user_id;type:bigint;not null;default:0" json:"user_id"`      // 用户id
}
