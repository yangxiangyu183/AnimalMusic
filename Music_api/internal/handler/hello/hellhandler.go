package hello

import (
	"net/http"

	"Music_api/internal/logic/hello"
	"Music_api/internal/svc"
	"Music_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HellHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hello.NewHellLogic(r.Context(), svcCtx)
		resp, err := l.Hell(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
