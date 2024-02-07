package model

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/liangqianyang/permissionBase/internal/svc"
	pb "github.com/liangqianyang/permissionBase/pb/permissionBase"
	"time"
)

type Users struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement;column:id;type:bigint;not null" json:"id"` // 主键
	LoginName string    `gorm:"column:login_name" json:"login_name"`                               // 登录名
	NickName  string    `gorm:"column:nickname" json:"nickname"`                                   // 昵称
	Password  string    `gorm:"column:password" json:"password"`                                   // 密码
	Mobile    string    `gorm:"column:mobile" json:"mobile"`                                       // 手机号
	Email     string    `gorm:"column:email" json:"email"`                                         // 邮箱
	Gender    int64     `gorm:"column:gender" json:"gender"`                                       // 性别 0:未知 1:男 2:女
	State     int64     `gorm:"column:state" json:"state"`                                         // 状态 1:正常 -1:禁用
	CreatedAt time.Time `gorm:"column:create_at" json:"created_at"`                                // 创建时间
	UpdatedAt time.Time `gorm:"column:update_at" json:"updated_at"`                                // 系统自动更新时间
}

// CheckUnique 检查用户是否唯一
func (m *Users) CheckUnique(ctx context.Context, svcCtx *svc.ServiceContext, in *pb.CreateUserRequest) error {
	var count int64
	svcCtx.Db.Model(&Users{}).Where(&Users{LoginName: in.LoginName}).Or(&Users{Mobile: in.Mobile}).Or(&Users{Mobile: in.Email}).Count(&count)
	if count > 0 {
		return errors.New("user already exists")
	}
	return nil
}

// CheckLogin 检查用户登录
func (m *Users) CheckLogin(ctx context.Context, svcCtx *svc.ServiceContext, in *pb.LoginRequest) (*Users, error) {
	var count int64
	svcCtx.Db.Model(&Users{}).Where("login_name = ? AND state = ?", in.LoginName, int64(pb.UserState_USER_STATE_ENABLE)).Count(&count)
	if count == 0 {
		return nil, errors.New("user not exists")
	}
	// 创建一个 MD5 的哈希对象
	hasher := md5.New()
	// 写入要计算哈希值的数据
	hasher.Write([]byte(in.Password))
	// 计算哈希值，并返回一个 16 字节的哈希值切片
	hash := hasher.Sum(nil)
	// 将哈希值切片转换为十六进制格式的字符串
	hashStr := hex.EncodeToString(hash)
	user := &Users{}
	svcCtx.Db.Model(&Users{}).Where("login_name = ? AND password = ?", in.LoginName, hashStr).First(user)
	if user.Id == 0 {
		return nil, errors.New("password error")
	}
	return user, nil
}
