package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type GreetMiddleware2Middleware struct {
}

func NewGreetMiddleware2Middleware() *GreetMiddleware2Middleware {
	return &GreetMiddleware2Middleware{}
}

func (m *GreetMiddleware2Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("NewGreetMiddleware2Middleware req:...")
		next(w, r)
		logx.Info("NewGreetMiddleware2Middleware resp:...")
	}
}
