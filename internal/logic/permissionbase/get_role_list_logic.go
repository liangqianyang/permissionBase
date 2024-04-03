package permissionbaselogic

import (
	"context"
	"github.com/liangqianyang/permissionBase/internal/model"

	"github.com/liangqianyang/permissionBase/internal/svc"
	"github.com/liangqianyang/permissionBase/pb/permissionBase"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRoleList 获取角色列表
func (l *GetRoleListLogic) GetRoleList(in *permissionBase.GetRoleListRequest) (*permissionBase.GetRoleListResponse, error) {
	var roles []model.Roles
	var total int64
	db := l.svcCtx.Db
	if in.Name != "" {
		db = db.Where("name like ?", in.Name+"%")
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
	if err := db.Model(&model.Roles{}).Count(&total).Error; err != nil {
		return nil, err
	}
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	list := make([]*permissionBase.RoleInfo, 0)
	for _, role := range roles {
		list = append(list, &permissionBase.RoleInfo{
			Id:          role.Id,
			Name:        role.Name,
			Description: role.Description,
			State:       permissionBase.RoleState(role.State),
			CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &permissionBase.GetRoleListResponse{
		List: list,
		Page: &permissionBase.PageInfo{Page: page, PageSize: pageSize, Total: total},
	}, nil
}
