package logic

import (
	"context"
	"errors"
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
	//判断用户是否存在
	var cnt int64
	l.svcCtx.EngineWeb.Table("users").Where("name=?", req.Name).Count(&cnt)
	if cnt > 0 {
		err = errors.New("用户已被注册")
		return
	}
	err = l.svcCtx.EngineWeb.Table("users").Create(&data).Error
	if err == nil {
		resp = new(types.RegisterResponse)
		resp.Message = "注册成功"
	}
	return
}
