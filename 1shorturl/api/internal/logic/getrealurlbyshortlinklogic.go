package logic

import (
	"context"
	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRealUrlByShortLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRealUrlByShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRealUrlByShortLinkLogic {
	return &GetRealUrlByShortLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRealUrlByShortLinkLogic) GetRealUrlByShortLink(req *types.ShortLinkToRealUrlReq) (resp *types.ShortLinkToRealUrlResp, err error) {
	in := transformclient.ShortLinkToRealUrlReq{ShortLink: req.ShortLink}
	out, err := l.svcCtx.TransformRpc.GetRealUrlByShortLink(l.ctx, &in)
	if err != nil {
		return nil, err
	}
	result := types.ShortLinkToRealUrlResp{RealUrl: out.RealUrl}
	return &result, nil
}
