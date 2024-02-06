package svc

import (
	"context"
	"github.com/qy-blog/permissionBase/internal/config"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		logc.Errorf(context.Background(), "gorm.Open error: %v", err)
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
