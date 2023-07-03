package logic

import (
	"context"
	"strconv"

	"15jwt/user/api/internal/svc"
	"15jwt/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(userId string) (resp *types.UserInfoReply, err error) {
	userIdInt, err := strconv.ParseInt(userId, 10, 64)

	var result *types.UserInfoReply = &types.UserInfoReply{
		Id:       userIdInt,
		Username: "123456",
	}

	return result, nil
}
