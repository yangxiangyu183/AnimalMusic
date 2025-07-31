package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user_rpc/internal/logic"
	"user_rpc/internal/svc"
	"user_rpc/internal/types"
)

func User_rpcHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUser_rpcLogic(r.Context(), svcCtx)
		resp, err := l.User_rpc(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
