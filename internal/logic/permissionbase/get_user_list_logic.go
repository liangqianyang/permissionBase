package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserList 获取用户列表
func (l *GetUserListLogic) GetUserList(in *permissionBase.GetUserListRequest) (*permissionBase.GetUserListResponse, error) {
	var users []model.Users
	var total int64
	db := l.svcCtx.Db
	if in.Username != "" {
		db = db.Where("username like ?", in.Username+"%")
	}
	if in.Nickname != "" {
		db = db.Where("nickname like ?", in.Nickname+"%")
	}
	if in.State != permissionBase.UserState_USER_STATE_UNKNOWN {
		db = db.Where("state = ?", in.State)
	}
	page := in.Page
	pageSize := in.PageSize
	if page <= 0 {
		page = model.DefaultPage
	}
	if pageSize <= 0 {
		pageSize = model.DefaultPageSize
	}
	if err := db.Model(&model.Users{}).Count(&total).Error; err != nil {
		return nil, err
	}
	db = db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Order("id desc")
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	list := make([]*permissionBase.UserInfo, 0)
	for _, user := range users {
		list = append(list, &permissionBase.UserInfo{
			Id:        int64(user.Id),
			Username:  user.Username,
			Nickname:  user.NickName,
			Mobile:    user.Mobile,
			Email:     user.Email,
			Gender:    permissionBase.GenderEnum(user.Gender),
			State:     permissionBase.UserState(user.State),
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &permissionBase.GetUserListResponse{
		List: list,
		Page: &permissionBase.PageInfo{Page: page, PageSize: pageSize, Total: total},
	}, nil
}
