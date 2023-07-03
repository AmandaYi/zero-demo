package logic

import (
	"16periodlimit/greet/internal/svc"
	"16periodlimit/greet/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	limitStore, _ := redis.NewRedis(redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		Pass:        "123456",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})
	periodLimit := limit.NewPeriodLimit(1, 10, limitStore, "periodLimit")
	code, err := periodLimit.Take("first")
	//local limit = tonumber(ARGV[1])
	//local window = tonumber(ARGV[2])
	//-- incrbt key 1 => key visis++
	//自增key，但是又有过期时间，1秒就过期
	//local current = redis.call("INCRBY", KEYS[1], 1)
	//-- 如果是第一次访问，设置过期时间 => TTL = window size
	//-- 因为是只限制一段时间的访问次数
	//if current == 1 then
	//redis.call("expire", KEYS[1], window)
	//return 1
	//elseif current < limit then
	//return 1
	//elseif current == limit then
	//return 2
	//else
	//return 0
	//end
	switch code {
	case -1:
		{
			logx.Info("Redis打开失败")
		}
	case limit.Allowed:
		{

			logx.Info("允许通过")
		}
	case limit.HitQuota:
		{
			logx.Info("请求正好达到配额")
		}
	case limit.OverQuota:
		{
			logx.Info("请求超出配额")
		}

	}
	//fmt.Println(req.Name)
	return &types.Response{Message: "limit test"}, nil
}
