package pinduoduo

import (
	"context"
	"github.com/mazesoul87/go-library/gorequest"
)

type GoodsPidGenerateResponse struct {
	PIdGenerateResponse struct {
		PIdList []struct {
			CreateTime int    `json:"create_time,omitempty"` // 推广位创建时间
			PidName    string `json:"pid_name,omitempty"`    // 推广位名称
			PId        string `json:"p_id,omitempty"`        // 调用方推广位ID
			MediaId    int    `json:"media_id,omitempty"`    // 媒体id
		} `json:"p_id_list"`
		RemainPidCount int `json:"remain_pid_count"` // PID剩余数量
	} `json:"p_id_generate_response"`
}

type GoodsPidGenerateResult struct {
	Result GoodsPidGenerateResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newGoodsPidGenerateResult(result GoodsPidGenerateResponse, body []byte, http gorequest.Response) *GoodsPidGenerateResult {
	return &GoodsPidGenerateResult{Result: result, Body: body, Http: http}
}

// GoodsPidGenerate 创建多多进宝推广位
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.pid.generate
func (c *Client) GoodsPidGenerate(ctx context.Context, notMustParams ...gorequest.Params) (*GoodsPidGenerateResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.goods.pid.generate")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.goods.pid.generate", notMustParams...)

	// 请求
	var response GoodsPidGenerateResponse
	request, err := c.request(ctx, span, params, &response)
	return newGoodsPidGenerateResult(response, request.ResponseBody, request), err
}
