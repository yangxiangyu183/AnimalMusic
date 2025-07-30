package handler

import (
	"net/http"

	"AnimalMusic/internal/logic"
	"AnimalMusic/internal/svc"
	"AnimalMusic/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AnimalMusicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAnimalMusicLogic(r.Context(), svcCtx)
		resp, err := l.AnimalMusic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
