package aswzk

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ElectricityBillOrderQueryResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data struct {
		RechargeAccount string  `json:"recharge_account"` // 充值账号
		RechargeMoney   float64 `json:"recharge_money"`   // 充值金额
		RechargeType    string  `json:"recharge_type"`    // 充值联系
		OrderNo         string  `json:"order_no"`         // 订单编号
		Remark          string  `json:"remark"`           // 订单备注
		OrderStatus     string  `json:"order_status"`     // 订单状态
	} `json:"data"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type ElectricityBillOrderQueryResult struct {
	Result ElectricityBillOrderQueryResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newElectricityBillOrderQueryResult(result ElectricityBillOrderQueryResponse, body []byte, http gorequest.Response) *ElectricityBillOrderQueryResult {
	return &ElectricityBillOrderQueryResult{Result: result, Body: body, Http: http}
}

// ElectricityBillOrderQuery 电费订单查询
func (c *Client) ElectricityBillOrderQuery(ctx context.Context, orderNo string, notMustParams ...gorequest.Params) (*ElectricityBillOrderQueryResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_no", orderNo)
	// 请求
	request, err := c.request(ctx, apiUrl+"/electricity_bill/order", params, http.MethodGet)
	if err != nil {
		return newElectricityBillOrderQueryResult(ElectricityBillOrderQueryResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ElectricityBillOrderQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newElectricityBillOrderQueryResult(response, request.ResponseBody, request), err
}
