package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type GreetMiddleware1Middleware struct {
}

func NewGreetMiddleware1Middleware() *GreetMiddleware1Middleware {
	return &GreetMiddleware1Middleware{}
}

func (m *GreetMiddleware1Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("NewGreetMiddleware1Middleware req:...")
		next(w, r)
		logx.Info("NewGreetMiddleware1Middleware resp:...")
	}
}
