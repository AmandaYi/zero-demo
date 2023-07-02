package handler

import (
	"13goqueue/api/goqueue/internal/logic"
	"context"
	"github.com/zeromicro/go-zero/core/service"

	"13goqueue/api/goqueue/internal/svc"
)

func RegisterJob(serverCtx *svc.ServiceContext, group *service.ServiceGroup) {
	group.Add(logic.NewProducerLogic(context.Background(), serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(), serverCtx))
}
