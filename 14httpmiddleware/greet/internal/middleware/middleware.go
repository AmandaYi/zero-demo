package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

func MiddlewareDemoFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logx.Info("request:...")
		next(writer, request)
		logx.Info("response:...")
	}
}
