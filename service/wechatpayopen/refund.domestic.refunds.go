package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type RefundDomesticRefundsResponse struct {
	RefundId            string    `json:"refund_id"`
	OutRefundNo         string    `json:"out_refund_no"`
	TransactionId       string    `json:"transaction_id"`
	OutTradeNo          string    `json:"out_trade_no"`
	Channel             string    `json:"channel"`
	UserReceivedAccount string    `json:"user_received_account"`
	SuccessTime         time.Time `json:"success_time"`
	CreateTime          time.Time `json:"create_time"`
	Status              string    `json:"status"`
	FundsAccount        string    `json:"funds_account"`
	Amount              struct {
		Total  int `json:"total"`
		Refund int `json:"refund"`
		From   []struct {
			Account string `json:"account"`
			Amount  int    `json:"amount"`
		} `json:"from"`
		PayerTotal       int    `json:"payer_total"`
		PayerRefund      int    `json:"payer_refund"`
		SettlementRefund int    `json:"settlement_refund"`
		SettlementTotal  int    `json:"settlement_total"`
		DiscountRefund   int    `json:"discount_refund"`
		Currency         string `json:"currency"`
	} `json:"amount"`
	PromotionDetail []struct {
		PromotionId  string `json:"promotion_id"`
		Scope        string `json:"scope"`
		Type         string `json:"type"`
		Amount       int    `json:"amount"`
		RefundAmount int    `json:"refund_amount"`
		GoodsDetail  struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`
			WechatpayGoodsId string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			UnitPrice        int    `json:"unit_price"`
			RefundAmount     int    `json:"refund_amount"`
			RefundQuantity   int    `json:"refund_quantity"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
}

type RefundDomesticRefundsResult struct {
	Result   RefundDomesticRefundsResponse // 结果
	Body     []byte                        // 内容
	Http     gorequest.Response            // 请求
	Err      error                         // 错误
	ApiError ApiError                      // 接口错误
}

func newRefundDomesticRefundsResult(result RefundDomesticRefundsResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *RefundDomesticRefundsResult {
	return &RefundDomesticRefundsResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// RefundDomesticRefunds 申请退款API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_9.shtml
func (c *Client) RefundDomesticRefunds(ctx context.Context, notMustParams ...*gorequest.Params) *RefundDomesticRefundsResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/refund/domestic/refunds", params, http.MethodPost)
	if err != nil {
		return newRefundDomesticRefundsResult(RefundDomesticRefundsResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response RefundDomesticRefundsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newRefundDomesticRefundsResult(response, request.ResponseBody, request, err, apiError)
}
