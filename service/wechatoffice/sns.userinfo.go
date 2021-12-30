package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SnsUserinfoResponse struct {
	Openid     string   `json:"openid"`            // 用户的唯一标识
	Nickname   string   `json:"nickname"`          // 用户昵称
	Sex        int      `json:"sex"`               // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province   string   `json:"province"`          // 用户个人资料填写的省份
	City       string   `json:"city"`              // 普通用户个人资料填写的城市
	Country    string   `json:"country"`           // 国家，如中国为CN
	Headimgurl string   `json:"headimgurl"`        // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege  []string `json:"privilege"`         // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Unionid    string   `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

type SnsUserinfoResult struct {
	Result SnsUserinfoResponse // 结果
	Byte   []byte              // 内容
	Err    error               // 错误
}

func NewSnsUserinfoResult(result SnsUserinfoResponse, byte []byte, err error) *SnsUserinfoResult {
	return &SnsUserinfoResult{Result: result, Byte: byte, Err: err}
}

// SnsUserinfo 拉取用户信息(需scope为 snsapi_userinfo)
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (app *App) SnsUserinfo(accessToken, openid string) *SnsUserinfoResult {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openid), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response SnsUserinfoResponse
	err = json.Unmarshal(body, &response)
	return NewSnsUserinfoResult(response, body, err)
}
