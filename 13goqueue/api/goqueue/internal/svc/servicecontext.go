package svc

import (
	"13goqueue/api/goqueue/internal/config"
	"github.com/zeromicro/go-queue/dq"
)

type ServiceContext struct {
	Config   config.Config
	Consumer dq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Consumer: dq.NewConsumer(c.DqConf),
	}
}
