// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/short/to/url",
				Handler: GetRealUrlByShortLinkHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/url/to/short",
				Handler: GetShortLinkByRealUrlHandler(serverCtx),
			},
		},
	)
}