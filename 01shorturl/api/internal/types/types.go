// Code generated by goctl. DO NOT EDIT.
package types

type RealUrlToShortLinkReq struct {
	RealUrl string `form:"realUrl"`
}

type RealUrlToShortLinkResp struct {
	ShortLink string `json:"shortLink"`
}

type ShortLinkToRealUrlReq struct {
	ShortLink string `form:"shortLink"`
}

type ShortLinkToRealUrlResp struct {
	RealUrl string `json:"realUrl"`
}
