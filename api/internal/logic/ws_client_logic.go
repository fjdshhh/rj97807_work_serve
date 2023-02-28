package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
)

type WsClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsClientLogic {
	return &WsClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsClientLogic) WsClient(req *types.WsClientRequest) (resp *types.WsClientResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
