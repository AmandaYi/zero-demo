package svc

import (
	"15jwt/order/api/internal/config"
	"15jwt/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  userclient.User
	UserRpc2 userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserRpc2: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
