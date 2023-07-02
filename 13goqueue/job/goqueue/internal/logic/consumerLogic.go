package logic

import (
	"13goqueue/job/goqueue/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type Consumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Consumer {
	return &Consumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Consumer) Start() {
	logx.Infof("start consumer \n")

	threading.GoSafe(func() {
		l.svcCtx.Consumer.Consume(func(body []byte) {
			logx.Infof("consumer job  %s \n", string(body))
		})
	})
}

func (l *Consumer) Stop() {
	logx.Infof("stop consumer \n")
}
