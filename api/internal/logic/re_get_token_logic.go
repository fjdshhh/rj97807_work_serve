package logic

import (
	"context"
	"rj97807_work_serve/api/models"
	"rj97807_work_serve/funcs"
	"rj97807_work_serve/utils"

	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReGetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReGetTokenLogic {
	return &ReGetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReGetTokenLogic) ReGetToken(req *types.ReGetTokenRequest) (resp *types.ReGetTokenResponse, err error) {
	var um models.UserModel
	// 解析Token获取name
	uc, err := funcs.AnalyzeToken(req.ReToken)
	if err != nil {
		return nil, err
	}
	// 跟Redis对比 是否存在
	redisToken, err := l.svcCtx.RDB.Get(l.ctx, uc.Name).Result()
	if err != nil {
		return nil, err
	}
	resp = new(types.ReGetTokenResponse)
	//确认是本人
	if redisToken == req.ReToken {
		err := l.svcCtx.EngineWeb.Table("users").Select("uid").Where("name=?", uc.Name).Find(&um).Error
		if err != nil {
			return nil, err
		}
		useToken, err := funcs.YieldToken(utils.TokenExpireTime, um.Uid, uc.Name)
		if err != nil {
			return nil, err
		}
		reFreshToken, err := funcs.YieldToken(utils.ReTokenExpireTime, um.Uid, uc.Name)
		if err != nil {
			return nil, err
		}
		resp.Token = useToken
		resp.ReToken = reFreshToken
	} else {
		resp.Token = ""
		resp.ReToken = ""
	}
	return
}
