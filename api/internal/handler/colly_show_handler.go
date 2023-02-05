package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rj97807_work_serve/api/internal/logic"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
)

func CollyShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollyShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCollyShowLogic(r.Context(), svcCtx)
		resp, err := l.CollyShow(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
