package logic

import (
	"context"
	"rj97807_work_serve/funcs"
	"rj97807_work_serve/utils"
	"time"

	"rj97807_work_serve/rpc/register/internal/svc"
	"rj97807_work_serve/rpc/register/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMailCheckedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMailCheckedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMailCheckedLogic {
	return &GetMailCheckedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMailCheckedLogic) GetMailChecked(in *rpc.MailRequest) (*rpc.MailResponse, error) {
	//获取验证码
	code := funcs.RandCode()
	//存储验证码
	l.svcCtx.RDB.Set(l.ctx, in.Email, code, time.Second*time.Duration(utils.CodeExpireTime))
	//发送验证码
	err := funcs.MailSendCode(in.Email, code)
	if err != nil {
		return nil, err
	}
	return &rpc.MailResponse{
		Message: "已发送验证码",
	}, nil
}
