package svc

import (
	"context"
	"github.com/apolloconfig/agollo/v4"
	apolloConfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/liangqianyang/permissionBase/internal/config"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Apollo agollo.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	apolloConf := &apolloConfig.AppConfig{
		AppID:          c.Apollo.AppID,
		Cluster:        c.Apollo.Cluster,
		IP:             c.Apollo.Ip,
		NamespaceName:  c.Apollo.Namespace,
		IsBackupConfig: c.Apollo.IsBackup,
		Secret:         c.Apollo.Secret,
	}

	client, err := agollo.StartWithConfig(func() (*apolloConfig.AppConfig, error) {
		return apolloConf, nil
	})
	if err != nil {
		logc.Errorf(context.Background(), "agollo.StartWithConfig error: %v", err)
		return nil
	}

	// 初始化数据库连接
	dataSource := client.GetStringValue("db_blog", "")
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		SkipDefaultTransaction: true,
	})

	if err != nil {
		logc.Errorf(context.Background(), "gorm.Open error: %v", err)
		return nil
	}

	return &ServiceContext{
		Config: c,
		Db:     db,
		Apollo: client,
	}
}
