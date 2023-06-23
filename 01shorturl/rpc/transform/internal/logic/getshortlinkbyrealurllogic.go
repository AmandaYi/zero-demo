package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/hash"
	"shorturl/rpc/transform/model"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/types/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortLinkByRealUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShortLinkByRealUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortLinkByRealUrlLogic {
	return &GetShortLinkByRealUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShortLinkByRealUrl 根据真实的url返回短连接
func (l *GetShortLinkByRealUrlLogic) GetShortLinkByRealUrl(in *transform.RealUrlToShortLinkReq) (*transform.RealUrlToShortLinkResp, error) {
	key := hash.Md5Hex([]byte(in.RealUrl))[:6]
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		ShortLink: key,
		RealUrl:   in.RealUrl,
	})
	if err != nil {
		return nil, err
	}
	return &transform.RealUrlToShortLinkResp{ShortLink: key}, nil
}
