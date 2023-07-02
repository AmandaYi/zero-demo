package logic

import (
	"13goqueue/api/goqueue/internal/svc"
	"context"
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"time"
)

type Producer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Producer {
	return &Producer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Producer) Start() {

	logx.Infof("start  Producer \n")
	threading.GoSafe(func() {
		producer := dq.NewProducer([]dq.Beanstalk{
			{
				Endpoint: "localhost:7771",
				Tube:     "tube1",
			},
			{
				Endpoint: "localhost:7772",
				Tube:     "tube2",
			},
		})
		for i := 1000; i < 1005; i++ {
			_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*1)
			if err != nil {
				logx.Error(err)
			}
		}
	})
}

func (l *Producer) Stop() {
	logx.Infof("stop Producer \n")
}
