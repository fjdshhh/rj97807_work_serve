package handler

import (
	"net/http"
	"rj97807_work_serve/ws/chat_ws"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rj97807_work_serve/api/internal/logic"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
)

func wsClientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsClientRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		//ws需要修改header元素
		r.Header.Set("Connection", "upgrade")
		chat_ws.ServeWs(w, r, svcCtx.EngineWeb)
		l := logic.NewWsClientLogic(r.Context(), svcCtx)
		resp, err := l.WsClient(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
