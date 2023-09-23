package pinduoduo

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PddDdkOauthGoodsPidGenerateResponse struct {
	OrderDetailResponse struct {
		SepMarketFee          int    `json:"sep_market_fee"`
		GoodsPrice            int    `json:"goods_price"`
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
		GoodsName             string `json:"goods_name"`
		BatchNo               string `json:"batch_no"`
		RedPacketType         int    `json:"red_packet_type"`
		UrlLastGenerateTime   int    `json:"url_last_generate_time"`
		GoodsQuantity         int    `json:"goods_quantity"`
		GoodsId               int64  `json:"goods_id"`
		SepParameters         string `json:"sep_parameters"`
		SepRate               int    `json:"sep_rate"`
		SubsidyType           int    `json:"subsidy_type"`
		ShareRate             int    `json:"share_rate"`
		CustomParameters      string `json:"custom_parameters"`
		GoodsThumbnailUrl     string `json:"goods_thumbnail_url"`
		PromotionAmount       int64  `json:"promotion_amount"`
		OrderPayTime          int    `json:"order_pay_time"`
		GroupId               int64  `json:"group_id"`
		SepPid                string `json:"sep_pid"`
		ReturnStatus          int    `json:"return_status"`
		OrderStatusDesc       string `json:"order_status_desc"`
		ShareAmount           int    `json:"share_amount"`
		GoodsCategoryName     string `json:"goods_category_name"`
		RequestId             string `json:"request_id"`
		GoodsSign             string `json:"goods_sign"`
		OrderSn               string `json:"order_sn"`
		ZsDuoId               int    `json:"zs_duo_id"`
	} `json:"order_detail_response"`
}

type PddDdkOauthGoodsPidGenerateResult struct {
	Result PddDdkOauthGoodsPidGenerateResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
}

func newPddDdkOauthGoodsPidGenerateResult(result PddDdkOauthGoodsPidGenerateResponse, body []byte, http gorequest.Response) *PddDdkOauthGoodsPidGenerateResult {
	return &PddDdkOauthGoodsPidGenerateResult{Result: result, Body: body, Http: http}
}

// PidGenerate 多多进宝推广位创建接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.goods.pid.generate
func (c *PddDdkOauthGoodsApi) PidGenerate(ctx context.Context, notMustParams ...*gorequest.Params) (*PddDdkOauthGoodsPidGenerateResult, error) {
	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.goods.pid.generate", notMustParams...)
	// 请求
	request, err := c.client.request(ctx, params)
	if err != nil {
		return newPddDdkOauthGoodsPidGenerateResult(PddDdkOauthGoodsPidGenerateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PddDdkOauthGoodsPidGenerateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPddDdkOauthGoodsPidGenerateResult(response, request.ResponseBody, request), err
}
