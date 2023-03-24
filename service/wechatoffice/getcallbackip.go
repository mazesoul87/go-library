package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

type GetCallBackIpResult struct {
	Result GetCallBackIpResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func newGetCallBackIpResult(result GetCallBackIpResponse, body []byte, http gorequest.Response, err error) *GetCallBackIpResult {
	return &GetCallBackIpResult{Result: result, Body: body, Http: http, Err: err}
}

// GetCallBackIp 获取微信callback IP地址
// callback IP即微信调用开发者服务器所使用的出口IP。
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html#2.%20%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1callback%20IP%E5%9C%B0%E5%9D%80
func (c *Client) GetCallBackIp(ctx context.Context) *GetCallBackIpResult {
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/getcallbackip?access_token=%s", c.getAccessToken(ctx)), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response GetCallBackIpResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetCallBackIpResult(response, request.ResponseBody, request, err)
}
