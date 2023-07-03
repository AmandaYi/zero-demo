// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"15jwt/user/rpc/internal/logic"
	"15jwt/user/rpc/internal/svc"
	"15jwt/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) GetMoney(ctx context.Context, in *user.IdRequest) (*user.UserMoneyResponse, error) {
	l := logic.NewGetMoneyLogic(ctx, s.svcCtx)
	return l.GetMoney(in)
}