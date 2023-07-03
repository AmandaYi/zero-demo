package logic

import (
	"context"
	"fmt"

	"15jwt/user/rpc/internal/svc"
	"15jwt/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMoneyLogic {
	return &GetMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMoneyLogic) GetMoney(in *user.IdRequest) (*user.UserMoneyResponse, error) {

	id := in.Id
	var money float64 = 10000000000
	var result user.UserMoneyResponse = user.UserMoneyResponse{}
	fmt.Println("远程调用", id)
	if id == "100" {
		result.Amount = money
	}
	return &result, nil
}
