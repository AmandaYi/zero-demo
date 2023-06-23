package logic

import (
	"context"
	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortLinkByRealUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShortLinkByRealUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortLinkByRealUrlLogic {
	return &GetShortLinkByRealUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShortLinkByRealUrlLogic) GetShortLinkByRealUrl(req *types.RealUrlToShortLinkReq) (resp *types.RealUrlToShortLinkResp, err error) {
	in := transformclient.RealUrlToShortLinkReq{RealUrl: req.RealUrl}
	out, err := l.svcCtx.TransformRpc.GetShortLinkByRealUrl(l.ctx, &in)
	if err != nil {
		return nil, err
	}
	result := types.RealUrlToShortLinkResp{ShortLink: out.ShortLink}
	return &result, nil
}
