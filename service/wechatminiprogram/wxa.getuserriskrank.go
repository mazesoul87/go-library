package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGetUserRiskRankResponse struct {
	Errcode  int    `json:"errcode"`   // 错误码
	Errmsg   string `json:"errmsg"`    // 错误信息
	RiskRank int    `json:"risk_rank"` // 用户风险等级，合法值为0,1,2,3,4，数字越大风险越高。
	UnoinId  int64  `json:"unoin_id"`  // 唯一请求标识，标记单次请求
}

type WxaGetUserRiskRankResult struct {
	Result WxaGetUserRiskRankResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func newWxaGetUserRiskRankResult(result WxaGetUserRiskRankResponse, body []byte, http gorequest.Response, err error) *WxaGetUserRiskRankResult {
	return &WxaGetUserRiskRankResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetUserRiskRank 获取用户安全等级
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/safety-control-capability/getUserRiskRank.html
func (c *Client) WxaGetUserRiskRank(ctx context.Context, notMustParams ...gorequest.Params) *WxaGetUserRiskRankResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/getuserriskrank?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	// 定义
	var response WxaGetUserRiskRankResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGetUserRiskRankResult(response, request.ResponseBody, request, err)
}
