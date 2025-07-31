package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user_api/internal/logic"
	"user_api/internal/svc"
	"user_api/internal/types"
)

func User_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUser_apiLogic(r.Context(), svcCtx)
		resp, err := l.User_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
