package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSourceWeb   string
		DataSourceColly string
	}
	Redis struct {
		Addr     string
		Password string
	}
	RegisterRpc zrpc.RpcClientConf
}
