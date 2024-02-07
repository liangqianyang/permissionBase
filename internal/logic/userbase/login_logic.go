package userbaselogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/qy-blog/permissionBase/internal/model"
	"time"

	"github.com/qy-blog/permissionBase/internal/svc"
	pb "github.com/qy-blog/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 登录
func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	userModel := &model.Users{}
	user, err := userModel.CheckLogin(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	token, err := l.generateLoginToken(user.LoginName)
	if err != nil {
		return nil, errors.New("generate token failed")
	}
	context.WithValue(l.ctx, "token", token)
	context.WithValue(l.ctx, "user_id", int64(user.Id))
	context.WithValue(l.ctx, "nick_name", user.NickName)
	return &pb.LoginResponse{
		Id:        int64(user.Id),
		LoginName: user.LoginName,
		Nickname:  user.NickName,
		Token:     token,
	}, nil
}

// generateLoginToken 生成登录token
func (l *LoginLogic) generateLoginToken(username string) (string, error) {
	// 设置过期时间为1小时
	expirationTime := time.Now().Add(time.Hour).Unix()

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
	}

	// 从配置中获取密钥
	secretSalt := l.svcCtx.Apollo.GetStringValue("secret_salt", "")
	secretKey := []byte(secretSalt)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %v", err)
	}

	return signedToken, nil
}
