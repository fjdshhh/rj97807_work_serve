package logic

import (
	"context"

	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollyShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollyShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollyShowLogic {
	return &CollyShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollyShowLogic) CollyShow(req *types.CollyShowRequest) (resp *types.CollyShowResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
