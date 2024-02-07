package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type ApolloConf struct {
	AppID     string
	Cluster   string
	Ip        string
	Namespace string
	Secret    string
	IsBackup  bool
}

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	Apollo ApolloConf
}
