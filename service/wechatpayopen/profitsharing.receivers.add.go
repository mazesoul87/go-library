package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingReceiversAddResponse struct {
	SubMchid       string `json:"sub_mchid"`                 // 子商户号
	Type           string `json:"type"`                      // 分账接收方类型
	Account        string `json:"account"`                   // 分账接收方账号
	Name           string `json:"name,omitempty"`            // 分账接收方全称
	RelationType   string `json:"relation_type"`             // 与分账方的关系类型
	CustomRelation string `json:"custom_relation,omitempty"` // 自定义的分账关系
}

type ProfitSharingReceiversAddResult struct {
	Result   ProfitSharingReceiversAddResponse // 结果
	Body     []byte                            // 内容
	Http     gorequest.Response                // 请求
	Err      error                             // 错误
	ApiError ApiError                          // 接口错误
}

func newProfitSharingReceiversAddResult(result ProfitSharingReceiversAddResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *ProfitSharingReceiversAddResult {
	return &ProfitSharingReceiversAddResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// ProfitSharingReceiversAdd 添加分账接收方API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_8.shtml
func (c *Client) ProfitSharingReceiversAdd(ctx context.Context, Type, account, name, relationType, customRelation string) *ProfitSharingReceiversAddResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSpAppid())      // 应用ID
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("type", Type)                 // 分账接收方类型
	if Type == MERCHANT_ID {
		params.Set("account", account) // 商户号
		params.Set("name", name)       // 商户全称
	}
	if Type == PERSONAL_OPENID && name != "" {
		params.Set("account", account) // 个人openid
		params.Set("name", name)       // 个人姓名
	}
	if Type == PERSONAL_SUB_OPENID && name != "" {
		params.Set("account", account) // 个人sub_openid
		params.Set("name", name)       // 个人姓名
	}
	params.Set("relation_type", relationType) // 与分账方的关系类型
	if relationType == "CUSTOM" {
		params.Set("custom_relation", customRelation) // 自定义的分账关系
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/receivers/delete", params, http.MethodPost)
	if err != nil {
		return newProfitSharingReceiversAddResult(ProfitSharingReceiversAddResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response ProfitSharingReceiversAddResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingReceiversAddResult(response, request.ResponseBody, request, err, apiError)
}
