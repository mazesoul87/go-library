package wechatminiprogram

import (
	"encoding/json"
	"fmt"
)

// GetCallBackIpResponse 返回参数
type GetCallBackIpResponse struct {
	IpList []string `json:"ip_list"`
}

// NewGetCallBackIpResult 构造函数
func NewGetCallBackIpResult(result GetCallBackIpResponse, byte []byte, err error) *Result {
	return &Result{
		GetCallBackIpResponse: result,
		Byte:                  byte,
		Err:                   err,
	}
}

// GetCallBackIp 获取微信callback IP地址
// callback IP即微信调用开发者服务器所使用的出口IP。
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html#2.%20%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1callback%20IP%E5%9C%B0%E5%9D%80
func (app *App) GetCallBackIp() *Result {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s", app.AccessToken), map[string]interface{}{}, "GET")
	// 定义
	var response GetCallBackIpResponse
	err = json.Unmarshal(body, &response)
	return NewGetCallBackIpResult(response, body, err)
}
