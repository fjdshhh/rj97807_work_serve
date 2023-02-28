package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rj97807_work_serve/api/internal/logic"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
)

func GetUserMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMenuRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMenu(&req, r.Header.Get("UserRole"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
