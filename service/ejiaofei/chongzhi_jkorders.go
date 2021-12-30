package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type ChOngZhiJkOrdersResponse struct {
	XMLName   xml.Name `xml:"response"`
	UserID    string   `xml:"userid"`    // 会员账号
	PorderID  string   `xml:"Porderid"`  // 鼎信平台订单号
	OrderID   string   `xml:"orderid"`   // 用户订单号
	Account   string   `xml:"account"`   // 需要充值的手机号码
	Face      string   `xml:"face"`      // 充值面值
	Amount    string   `xml:"amount"`    // 购买数量
	StartTime string   `xml:"starttime"` // 开始时间
	State     string   `xml:"state"`     // 订单状态
	EndTime   string   `xml:"endtime"`   // 结束时间
	Error     string   `xml:"error"`     // 错误提示
}

type ChOngZhiJkOrdersResult struct {
	Result ChOngZhiJkOrdersResponse // 结果
	Body   []byte                   // 内容
	Err    error                    // 错误
}

func NewChOngZhiJkOrdersResult(result ChOngZhiJkOrdersResponse, body []byte, err error) *ChOngZhiJkOrdersResult {
	return &ChOngZhiJkOrdersResult{Result: result, Body: body, Err: err}
}

// ChOngZhiJkOrders 话费充值接口
// orderid 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
// face 充值面值	以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信
// account 手机号码	需要充值的手机号码
func (app *App) ChOngZhiJkOrders(orderID string, face int, account string) *ChOngZhiJkOrdersResult {
	// 参数
	param := NewParams()
	param.Set("orderid", orderID)
	param.Set("face", face)
	param.Set("account", account)
	param.Set("amount", 1)
	params := app.NewParamsWith(param)
	// 签名
	app.signStr = fmt.Sprintf("userid%vpwd%vorderid%vface%vaccount%vamount1", app.UserID, app.Pwd, orderID, face, account)
	// 请求
	body, err := app.request("http://api.ejiaofei.net:11140/chongzhi_jkorders.do", params, http.MethodGet)
	// 定义
	var response ChOngZhiJkOrdersResponse
	err = xml.Unmarshal(body, &response)
	return NewChOngZhiJkOrdersResult(response, body, err)
}
