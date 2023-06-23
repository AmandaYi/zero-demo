package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformclient"
)

type ServiceContext struct {
	Config       config.Config
	TransformRpc transformclient.Transform
}

// NewServiceContext api用到的rpc的内容，都要是client包开头的
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		TransformRpc: transformclient.NewTransform(zrpc.MustNewClient(c.TransformRpc)),
	}
}
