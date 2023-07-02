package logic

import (
	"context"

	"12zrpc/internal/svc"
	"12zrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *__.HelloRequest) (*__.HelloReply, error) {
	// todo: add your logic here and delete this line

	return &__.HelloReply{}, nil
}
