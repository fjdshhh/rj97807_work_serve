package logic

import (
	"context"
	"math"

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
	var data []types.Article
	var totalNum int64
	var totalSize int64
	err = l.svcCtx.EngineColly.Table("crow").Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Order("data desc").Find(&data).Error
	if err != nil {
		return
	}
	err = l.svcCtx.EngineColly.Table("crow").Count(&totalNum).Error
	totalSize = int64(math.Ceil(float64(totalNum / req.PageSize)))
	resp = new(types.CollyShowResponse)
	resp.Data = data
	resp.TotalSize = totalSize
	resp.TotalNum = totalNum
	return
}
