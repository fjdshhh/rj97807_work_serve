package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rj97807_work_serve/api/internal/logic"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
)

func UserRegisterBeforeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterBeforeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRegisterBeforeLogic(r.Context(), svcCtx)
		resp, err := l.UserRegisterBefore(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
