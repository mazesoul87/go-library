package golog

import (
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/datatypes"
)

// ApiPostgresqlLog 结构体
type ApiPostgresqlLog struct {
	LogId                 uint           `gorm:"primaryKey" json:"log_id,omitempty"`                   //【记录】编号
	RequestTime           TimeString     `gorm:"index" json:"request_time,omitempty"`                  //【请求】时间
	RequestUri            string         `gorm:"type:text" json:"request_uri,omitempty"`               //【请求】链接
	RequestUrl            string         `gorm:"type:text" json:"request_url,omitempty"`               //【请求】链接
	RequestApi            string         `gorm:"type:text;index" json:"request_api,omitempty"`         //【请求】接口
	RequestMethod         string         `gorm:"type:text;index" json:"request_method,omitempty"`      //【请求】方式
	RequestParams         datatypes.JSON `gorm:"type:jsonb" json:"request_params,omitempty"`           //【请求】参数
	RequestHeader         datatypes.JSON `gorm:"type:jsonb" json:"request_header,omitempty"`           //【请求】头部
	ResponseHeader        datatypes.JSON `gorm:"type:jsonb" json:"response_header,omitempty"`          //【返回】头部
	ResponseStatusCode    int            `gorm:"type:bigint" json:"response_status_code,omitempty"`    //【返回】状态码
	ResponseBody          datatypes.JSON `gorm:"type:jsonb" json:"response_body,omitempty"`            //【返回】内容
	ResponseContentLength int64          `gorm:"type:bigint" json:"response_content_length,omitempty"` //【返回】大小
	ResponseTime          TimeString     `gorm:"index" json:"response_time,omitempty"`                 //【返回】时间
	SystemHostName        string         `gorm:"type:text" json:"system_host_name,omitempty"`          //【系统】主机名
	SystemInsideIp        string         `gorm:"type:text" json:"system_inside_ip,omitempty"`          //【系统】内网ip
	GoVersion             string         `gorm:"type:text" json:"go_version,omitempty"`                //【程序】Go版本
}

// GormMiddleware 中间件
func (c *ApiClient) GormMiddleware(request gorequest.Response) {
	c.GormRecord(ApiPostgresqlLog{
		RequestTime:           TimeString{Time: request.RequestTime},                            //【请求】时间
		RequestUri:            request.RequestUri,                                               //【请求】链接
		RequestUrl:            gorequest.UriParse(request.RequestUri).Url,                       //【请求】链接
		RequestApi:            gorequest.UriParse(request.RequestUri).Path,                      //【请求】接口
		RequestMethod:         request.RequestMethod,                                            //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),  //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),  //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                       //【返回】状态码
		ResponseBody:          request.ResponseBody,                                             //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                    //【返回】大小
		ResponseTime:          TimeString{Time: request.ResponseTime},                           //【返回】时间
	})
}
