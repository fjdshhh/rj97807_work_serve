package logic

import (
	"context"
	"rj97807_work_serve/api/models"
	"rj97807_work_serve/funcs"
	"rj97807_work_serve/utils"
	"time"

	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	data := new(models.UserModel)
	// 验证存在
	err = l.svcCtx.EngineWeb.Table("users").Where("name=? and password=?", req.Name, funcs.Md5(req.Pwd)).Take(&data).Error
	if err != nil {
		return nil, err
	}
	//生成token
	useToken, err := funcs.YieldToken(utils.TokenExpireTime, data.Uid, data.Name)
	if err != nil {
		return nil, err
	}
	//token过期则刷新
	reFreshToken, err := funcs.YieldToken(utils.ReTokenExpireTime, data.Uid, data.Name)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RDB.Set(l.ctx, req.Name, reFreshToken, time.Second*time.Duration(utils.ReTokenExpireTime)).Err()
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	logx.Info("用户登录:", data.Name)
	resp.Token = useToken
	resp.ReToken = reFreshToken
	resp.Role = data.Role
	return
}
