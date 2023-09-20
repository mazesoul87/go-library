package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsJsapiResponse struct {
	PrepayId string `json:"prepay_id"`
}

type PayPartnerTransactionsJsapiResult struct {
	Result   PayPartnerTransactionsJsapiResponse // 结果
	Body     []byte                              // 内容
	Http     gorequest.Response                  // 请求
	Err      error                               // 错误
	ApiError ApiError                            // 接口错误
}

func newPayPartnerTransactionsJsapiResult(result PayPartnerTransactionsJsapiResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsJsapiResult {
	return &PayPartnerTransactionsJsapiResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsJsapi JSAPI下单
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_1.shtml
func (c *Client) PayPartnerTransactionsJsapi(ctx context.Context, notMustParams ...*gorequest.Params) *PayPartnerTransactionsJsapiResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sp_appid", c.GetSpAppid())   // 服务商应用ID
	params.Set("sp_mchid", c.GetSpMchId())   // 服务商户号
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/pay/partner/transactions/jsapi", params, http.MethodPost)
	if err != nil {
		return newPayPartnerTransactionsJsapiResult(PayPartnerTransactionsJsapiResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 结果
	var response PayPartnerTransactionsJsapiResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsJsapiResult(response, request.ResponseBody, request, err, apiError)
}
