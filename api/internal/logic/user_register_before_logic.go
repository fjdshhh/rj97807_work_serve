package logic

import (
	"context"
	"errors"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
	"rj97807_work_serve/rpc/register/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterBeforeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterBeforeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterBeforeLogic {
	return &UserRegisterBeforeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterBeforeLogic) UserRegisterBefore(req *types.RegisterBeforeRequest) (resp *types.RegisterBeforeResponse, err error) {
	//判断用户是否存在
	var cnt int64
	l.svcCtx.EngineWeb.Table("users").Where("name=?", req.Name).Count(&cnt)
	if cnt > 0 {
		err = errors.New("用户名已被注册")
		return
	}
	l.svcCtx.EngineWeb.Table("users").Where("email=?", req.Email).Count(&cnt)
	if cnt > 0 {
		err = errors.New("邮箱已被注册")
		return
	}

	getMsg, err := l.svcCtx.RegisterRpc.GetMailChecked(l.ctx, &rpc.MailRequest{
		Name:  req.Name,
		Email: req.Email,
	})
	if getMsg.Message == "已发送验证码" {
		resp = new(types.RegisterBeforeResponse)
		resp.Message = "已发送验证码"
	}
	return
}
