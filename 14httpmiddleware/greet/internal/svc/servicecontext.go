package svc

import (
	"14httpmiddleware/greet/internal/config"
	"14httpmiddleware/greet/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	GreetMiddleware1 rest.Middleware
	GreetMiddleware2 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		GreetMiddleware1: middleware.NewGreetMiddleware1Middleware().Handle,
		GreetMiddleware2: middleware.NewGreetMiddleware2Middleware().Handle,
	}
}
