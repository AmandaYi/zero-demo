package handler

import (
	"net/http"

	"15jwt/user/api/internal/logic"
	"15jwt/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		result, err := l.UserInfo(userId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//httpx.Ok(w)
			httpx.OkJsonCtx(r.Context(), w, result)
		}
	}
}
