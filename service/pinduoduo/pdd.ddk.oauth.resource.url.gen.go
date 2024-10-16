package pinduoduo

import (
	"context"
	"github.com/mazesoul87/go-library/utils/gorequest"
)

type PddDdkOauthResourceUrlGenResponse struct {
	OrderUrlGenResponse struct {
		SepMarketFee          int    `json:"sep_market_fee"`
		ResourcePrice         int    `json:"Resource_price"`
		SepDuoId              int    `json:"sep_duo_id"`
		Pid                   string `json:"pid"`
		PromotionRate         int    `json:"promotion_rate"`
		CpsSign               string `json:"cps_sign"`
		Type                  int    `json:"type"`
		SubsidyDuoAmountLevel int    `json:"subsidy_duo_amount_level"`
		OrderStatus           int    `json:"order_status"`
		CatIds                []int  `json:"cat_ids"`
		OrderCreateTime       int64  `json:"order_create_time"`
		IsDirect              int    `json:"is_direct"`
		OrderGroupSuccessTime int    `json:"order_group_success_time"`
		MallId                int    `json:"mall_id"`
		OrderAmount           int64  `json:"order_amount"`
		PriceCompareStatus    int    `json:"price_compare_status"`
		MallName              string `json:"mall_name"`
		OrderModifyAt         int    `json:"order_modify_at"`
		AuthDuoId             int    `json:"auth_duo_id"`
		CpaNew                int    `json:"cpa_new"`
		ResourceName          string `json:"Resource_name"`
		BatchNo               string `json:"batch_no"`
		RedPacketType         int    `json:"red_packet_type"`
		UrlLastGenerateTime   int    `json:"url_last_generate_time"`
		ResourceQuantity      int    `json:"Resource_quantity"`
		ResourceId            int64  `json:"Resource_id"`
		SepParameters         string `json:"sep_parameters"`
		SepRate               int    `json:"sep_rate"`
		SubsidyType           int    `json:"subsidy_type"`
		ShareRate             int    `json:"share_rate"`
		CustomParameters      string `json:"custom_parameters"`
		ResourceThumbnailUrl  string `json:"Resource_thumbnail_url"`
		PromotionAmount       int64  `json:"promotion_amount"`
		OrderPayTime          int    `json:"order_pay_time"`
		GroupId               int64  `json:"group_id"`
		SepPid                string `json:"sep_pid"`
		ReturnStatus          int    `json:"return_status"`
		OrderStatusDesc       string `json:"order_status_desc"`
		ShareAmount           int    `json:"share_amount"`
		ResourceCategoryName  string `json:"Resource_category_name"`
		RequestId             string `json:"request_id"`
		ResourceSign          string `json:"Resource_sign"`
		OrderSn               string `json:"order_sn"`
		ZsDuoId               int    `json:"zs_duo_id"`
	} `json:"order_UrlGen_response"`
}

type PddDdkOauthResourceUrlGenResult struct {
	Result PddDdkOauthResourceUrlGenResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newPddDdkOauthResourceUrlGenResult(result PddDdkOauthResourceUrlGenResponse, body []byte, http gorequest.Response) *PddDdkOauthResourceUrlGenResult {
	return &PddDdkOauthResourceUrlGenResult{Result: result, Body: body, Http: http}
}

// OauthResourceUrlGen 拼多多主站频道推广接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.resource.url.gen
func (c *Client) OauthResourceUrlGen(ctx context.Context, notMustParams ...gorequest.Params) (*PddDdkOauthResourceUrlGenResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.oauth.resource.url.gen")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.resource.url.gen", notMustParams...)

	// 请求
	var response PddDdkOauthResourceUrlGenResponse
	request, err := c.request(ctx, span, params, &response)
	return newPddDdkOauthResourceUrlGenResult(response, request.ResponseBody, request), err
}
