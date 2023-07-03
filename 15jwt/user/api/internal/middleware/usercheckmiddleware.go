package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserCheckMiddleware struct {
}

func NewUserCheckMiddleware() *UserCheckMiddleware {
	return &UserCheckMiddleware{}
}

func (m *UserCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("req:...")
		next(w, r)
		logx.Info("resp:...")
	}
}
