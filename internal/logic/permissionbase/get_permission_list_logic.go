package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionListLogic {
	return &GetPermissionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPermissionList 获取权限列表
func (l *GetPermissionListLogic) GetPermissionList(in *permissionBase.GetPermissionListRequest) (*permissionBase.GetPermissionListResponse, error) {
	var permissions []model.Permissions
	var total int64
	db := l.svcCtx.Db
	if in.Name != "" {
		db = db.Where("name like ?", in.Name+"%")
	}
	if in.IdentifyName != "" {
		db = db.Where("identify_name like ?", in.IdentifyName+"%")
	}
	if in.Controller != "" {
		db = db.Where("controller = ?", in.Controller)
	}
	if in.Action != "" {
		db = db.Where("action = ?", in.Action)
	}
	if in.State != 0 {
		db = db.Where("status = ?", in.State)
	}
	page := in.Page
	pageSize := in.PageSize
	if page <= 0 {
		page = model.DefaultPage
	}
	if pageSize <= 0 {
		pageSize = model.DefaultPageSize
	}
	db = db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Order("id desc")
	if err := db.Model(&model.Permissions{}).Count(&total).Error; err != nil {
		return nil, err
	}
	if err := db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	list := make([]*permissionBase.PermissionInfo, 0)
	for _, permission := range permissions {
		list = append(list, &permissionBase.PermissionInfo{
			Id:           permission.Id,
			Name:         permission.Name,
			IdentifyName: permission.IdentifyName,
			Controller:   permission.Controller,
			Action:       permission.Action,
			State:        permissionBase.PermissionState(permission.State),
			CreatedAt:    permission.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    permission.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &permissionBase.GetPermissionListResponse{
		List: list,
		Page: &permissionBase.PageInfo{Page: page, PageSize: pageSize, Total: total},
	}, nil
}
