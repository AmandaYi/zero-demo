package logic

import (
	"context"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/types/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRealUrlByShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRealUrlByShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRealUrlByShortLinkLogic {
	return &GetRealUrlByShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRealUrlByShortLink 根据短连接返回真实的url
func (l *GetRealUrlByShortLinkLogic) GetRealUrlByShortLink(in *transform.ShortLinkToRealUrlReq) (*transform.ShortLinkToRealUrlResp, error) {
	//直接查询model
	result, err := l.svcCtx.Model.FindOne(l.ctx, in.ShortLink)
	if err != nil {
		return nil, err
	}
	return &transform.ShortLinkToRealUrlResp{RealUrl: result.RealUrl}, nil
}
