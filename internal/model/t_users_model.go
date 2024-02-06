package model

import (
	"errors"
	"github.com/qy-blog/permissionBase/internal/svc"
	pb "github.com/qy-blog/permissionBase/pb/permissionBase"
	"time"
)

type Users struct {
	Id                uint64    `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"`                     // 主键
	LoginName         string    `gorm:"column:login_name" json:"login_name"`                                                   // 登录名
	NickName          string    `gorm:"column:nickname" json:"nickname"`                                                       // 昵称
	Password          string    `gorm:"column:password" json:"password"`                                                       // 密码
	Mobile            string    `gorm:"column:mobile" json:"mobile"`                                                           // 手机号
	Email             string    `gorm:"column:email" json:"email"`                                                             // 邮箱
	Gender            int64     `gorm:"column:gender" json:"gender"`                                                           // 性别 0:未知 1:男 2:女
	State             int64     `gorm:"column:state" json:"state"`                                                             // 状态 1:正常 -1:禁用
	CreateTime        time.Time `gorm:"column:create_time;type:timestamp;autoCreateTime" json:"create_time"`                   // 创建时间
	SysAutoUpdateTime time.Time `gorm:"column:sys_auto_update_time;type:timestamp;autoUpdateTime" json:"sys_auto_update_time"` // 系统自动更新时间
}

// CheckUnique 检查用户是否唯一
func (m *Users) CheckUnique(svcCtx *svc.ServiceContext, in *pb.CreateUserRequest) error {
	var count int64
	svcCtx.Db.Model(&Users{}).Where(&Users{LoginName: in.LoginName}).Or(&Users{Mobile: in.Mobile}).Or(&Users{Mobile: in.Email}).Count(&count)
	if count > 0 {
		return errors.New("user already exists")
	}
	return nil
}
