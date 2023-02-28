package logic

import (
	"context"

	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenuLogic {
	return &GetUserMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenuLogic) GetUserMenu(req *types.GetMenuRequest, userRole string) (resp *types.GetMenuResponse, err error) {
	var data []types.Menu
	err = l.svcCtx.EngineWeb.Table("role_menu").Where("role=?", userRole).Find(&data).Error
	if err != nil {
		return
	}
	resp = new(types.GetMenuResponse)
	resp.Data = data
	return
}
