package logic

import (
	"context"
	"fmt"
	"rj97807_work_serve/api/models"
	"rj97807_work_serve/funcs"

	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	data := new(models.UserModel)
	data.Name = req.Name
	data.Password = funcs.Md5(req.Pwd)
	data.Role = 0
	data.Uid = funcs.GetUUID()
	data.Email = req.Email

	RCode, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return
	}
	//验证通过
	if RCode == req.Code {
		err = l.svcCtx.EngineWeb.Table("users").Create(&data).Error
		if err == nil {
			resp = new(types.RegisterResponse)
			resp.Message = "注册成功"
			//注册完删除key
			_, err := l.svcCtx.RDB.Del(l.ctx, req.Email).Result()
			if err != nil {
				fmt.Printf("%x", "失败"+err.Error())
			}
		}
	} else {
		resp = new(types.RegisterResponse)
		resp.Message = "验证码错误"
	}
	return
}
