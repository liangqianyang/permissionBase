package userbaselogic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/qy-blog/permissionBase/internal/model"
	"github.com/qy-blog/permissionBase/internal/svc"
	pb "github.com/qy-blog/permissionBase/pb/permissionBase"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 创建用户
func (l *CreateUserLogic) CreateUser(in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := &model.Users{
		LoginName: in.LoginName,
		NickName:  in.Nickname,
		Mobile:    in.Mobile,
		Email:     in.Email,
		Gender:    in.Gender,
		State:     int64(pb.UserState_USER_STATE_ENABLE),
	}

	err := newUser.CheckUnique(l.svcCtx, in)

	if err := newUser.CheckUnique(l.svcCtx, in); err != nil {
		logc.Errorf(l.ctx, "create user failed: user already exists")
		return nil, errors.New("user already exists")
	}

	// 创建一个 MD5 的哈希对象
	hasher := md5.New()
	// 写入要计算哈希值的数据
	hasher.Write([]byte(in.Password))
	// 计算哈希值，并返回一个 16 字节的哈希值切片
	hash := hasher.Sum(nil)
	// 将哈希值切片转换为十六进制格式的字符串
	hashStr := hex.EncodeToString(hash)
	newUser.Password = hashStr

	err = l.svcCtx.Db.Create(newUser).Error
	if err != nil {
		logc.Errorf(l.ctx, "create user failed: %v", err)
		return nil, errors.New("create user failed,reason: " + err.Error())
	}
	return &pb.CreateUserResponse{
		Id: int64(newUser.Id),
	}, nil
}
