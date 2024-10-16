package pinduoduo

import (
	"context"
	"github.com/mazesoul87/go-library/utils/gorequest"
)

type PddDdkOauthCashGiftCreateResponse struct {
	CreateCashgiftResponse struct {
		CashGiftId float64 `json:"cash_gift_id"` // 礼金ID
		Success    bool    `json:"success"`      // 创建结果
	} `json:"create_cashgift_response"`
}

type PddDdkOauthCashGiftCreateResult struct {
	Result PddDdkOauthCashGiftCreateResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newPddDdkOauthCashGiftCreateResult(result PddDdkOauthCashGiftCreateResponse, body []byte, http gorequest.Response) *PddDdkOauthCashGiftCreateResult {
	return &PddDdkOauthCashGiftCreateResult{Result: result, Body: body, Http: http}
}

// OauthCashGiftCreate 创建多多礼金
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.create
func (c *Client) OauthCashGiftCreate(ctx context.Context, notMustParams ...gorequest.Params) (*PddDdkOauthCashGiftCreateResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.oauth.cashgift.create")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.create", notMustParams...)

	// 请求
	var response PddDdkOauthCashGiftCreateResponse
	request, err := c.request(ctx, span, params, &response)
	return newPddDdkOauthCashGiftCreateResult(response, request.ResponseBody, request), err
}
