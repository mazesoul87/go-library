package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpDeleteResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CgiBinWxOpenQrCodeJumpDeleteResult struct {
	Result CgiBinWxOpenQrCodeJumpDeleteResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
}

func newCgiBinWxOpenQrCodeJumpDeleteResult(result CgiBinWxOpenQrCodeJumpDeleteResponse, body []byte, http gorequest.Response) *CgiBinWxOpenQrCodeJumpDeleteResult {
	return &CgiBinWxOpenQrCodeJumpDeleteResult{Result: result, Body: body, Http: http}
}

// CgiBinWxOpenQrCodeJumpDelete 删除已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpdelete.html
func (c *Client) CgiBinWxOpenQrCodeJumpDelete(ctx context.Context, prefix string, notMustParams ...gorequest.Params) (*CgiBinWxOpenQrCodeJumpDeleteResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newCgiBinWxOpenQrCodeJumpDeleteResult(CgiBinWxOpenQrCodeJumpDeleteResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("prefix", prefix)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/wxopen/qrcodejumpdelete?access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newCgiBinWxOpenQrCodeJumpDeleteResult(CgiBinWxOpenQrCodeJumpDeleteResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinWxOpenQrCodeJumpDeleteResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinWxOpenQrCodeJumpDeleteResult(response, request.ResponseBody, request), err
}
