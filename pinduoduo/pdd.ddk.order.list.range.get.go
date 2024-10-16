package pinduoduo

import (
	"context"
	"github.com/mazesoul87/go-library/gorequest"
)

type OrderListRangeGetResponse struct {
	OrderListGetResponse struct {
		TotalCount  int    `json:"total_count"`
		LastOrderId string `json:"last_order_id"`
		OrderList   []struct {
			SepMarketFee          int    `json:"sep_market_fee"`
			GoodsPrice            int64  `json:"goods_price"`
			SepDuoId              int    `json:"sep_duo_id"`
			PromotionRate         int64  `json:"promotion_rate"`
			Type                  int    `json:"type"`
			SubsidyDuoAmountLevel int    `json:"subsidy_duo_amount_level"`
			CatIds                []int  `json:"cat_ids"`
			OrderStatus           int    `json:"order_status"`
			OrderCreateTime       int64  `json:"order_create_time"`
			IsDirect              int    `json:"is_direct"`
			OrderGroupSuccessTime int    `json:"order_group_success_time"`
			MallId                int    `json:"mall_id"`
			OrderAmount           int64  `json:"order_amount"`
			PriceCompareStatus    int    `json:"price_compare_status"`
			OrderModifyAt         int    `json:"order_modify_at"`
			AuthDuoId             int    `json:"auth_duo_id"`
			CpaNew                int    `json:"cpa_new"`
			GoodsName             string `json:"goods_name"`
			BatchNo               string `json:"batch_no"`
			RedPacketType         int    `json:"red_packet_type"`
			GoodsQuantity         int    `json:"goods_quantity"`
			FailReason            string `json:"fail_reason,omitempty"`
			GoodsId               int64  `json:"goods_id"`
			SepParameters         string `json:"sep_parameters"`
			SepRate               int    `json:"sep_rate"`
			SubsidyType           int    `json:"subsidy_type"`
			CustomParameters      string `json:"custom_parameters"`
			GoodsThumbnailUrl     string `json:"goods_thumbnail_url"`
			ShareRate             int    `json:"share_rate"`
			PromotionAmount       int64  `json:"promotion_amount"`
			OrderPayTime          int64  `json:"order_pay_time"`
			OrderReceiveTime      int64  `json:"order_receive_time"`
			OrderSettleTime       int64  `json:"order_settle_time"`
			ActivityTags          []int  `json:"activity_tags"`
			GroupId               int64  `json:"group_id"`
			SepPid                string `json:"sep_pid"`
			OrderStatusDesc       string `json:"order_status_desc"`
			ShareAmount           int    `json:"share_amount"`
			OrderId               string `json:"order_id"`
			GoodsSign             string `json:"goods_sign"`
			OrderSn               string `json:"order_sn"`
			OrderVerifyTime       int64  `json:"order_verify_time"`
			PId                   string `json:"p_id"`
			ZsDuoId               int    `json:"zs_duo_id"`
		} `json:"order_list"`
		RequestId string `json:"request_id"`
	} `json:"order_list_get_response"`
}

type OrderListRangeGetResult struct {
	Result OrderListRangeGetResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newOrderListRangeGetResult(result OrderListRangeGetResponse, body []byte, http gorequest.Response) *OrderListRangeGetResult {
	return &OrderListRangeGetResult{Result: result, Body: body, Http: http}
}

// OrderListRangeGet 用时间段查询推广订单接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.order.list.range.get
func (c *Client) OrderListRangeGet(ctx context.Context, notMustParams ...gorequest.Params) (*OrderListRangeGetResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.order.list.range.get")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.order.list.range.get", notMustParams...)

	// 请求
	var response OrderListRangeGetResponse
	request, err := c.request(ctx, span, params, &response)
	return newOrderListRangeGetResult(response, request.ResponseBody, request), err
}
