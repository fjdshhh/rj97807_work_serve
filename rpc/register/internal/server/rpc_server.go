// Code generated by goctl. DO NOT EDIT!
// Source: register.proto

package server

import (
	"context"

	"rj97807_work_serve/rpc/register/internal/logic"
	"rj97807_work_serve/rpc/register/internal/svc"
	"rj97807_work_serve/rpc/register/types/rpc"
)

type RpcServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedRpcServer
}

func NewRpcServer(svcCtx *svc.ServiceContext) *RpcServer {
	return &RpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcServer) GetMailChecked(ctx context.Context, in *rpc.MailRequest) (*rpc.MailResponse, error) {
	l := logic.NewGetMailCheckedLogic(ctx, s.svcCtx)
	return l.GetMailChecked(in)
}
