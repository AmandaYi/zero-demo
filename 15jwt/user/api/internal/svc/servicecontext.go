package svc

import (
	"15jwt/user/api/internal/config"
	"15jwt/user/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserCheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserCheck: middleware.NewUserCheckMiddleware().Handle,
	}
}
