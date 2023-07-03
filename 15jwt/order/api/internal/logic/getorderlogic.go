package logic

import (
	"15jwt/order/api/internal/svc"
	"15jwt/order/api/internal/types"
	"15jwt/user/rpc/userclient"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReplay, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{Id: req.Id})

	if err != nil {
		return nil, err
	}
	if user.Name == "test" {
		return nil, errors.New("用户不存在")
	}
	var result = types.OrderReplay{
		Id:   user.Id,
		Name: user.Name,
	}
	fmt.Println("客户端", req.Id)
	user2, err := l.svcCtx.UserRpc2.GetMoney(l.ctx, &userclient.IdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	result.Amount = user2.GetAmount()
	return &result, nil
}
