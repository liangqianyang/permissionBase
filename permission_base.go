package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/qy-blog/permissionBase/internal/config"
	permissionServer "github.com/qy-blog/permissionBase/internal/server/permissionbase"
	userServer "github.com/qy-blog/permissionBase/internal/server/userbase"
	"github.com/qy-blog/permissionBase/internal/svc"
	"github.com/qy-blog/permissionBase/pb/permissionBase"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configFile = flag.String("f", "etc/permission_base.yaml", "the config file")
	serviceTag = flag.String("t", "", "service tag for dyeing")
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		permissionBase.RegisterPermissionBaseServer(grpcServer, permissionServer.NewPermissionBaseServer(ctx))
		permissionBase.RegisterUserBaseServer(grpcServer, userServer.NewUserBaseServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()

	// 注册服务tag
	if serviceTag != nil && *serviceTag != "" {
		c.Consul.Tag = append(c.Consul.Tag, *serviceTag)
	}

	err := consul.RegisterService(c.ListenOn, c.Consul)
	if err != nil {
		logc.Errorf(context.Background(), "register consul service failed: %v", err)
		return
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
	<-make(chan struct{})
}
