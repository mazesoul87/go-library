package wechatqy

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) Config(key string) *Client {
	c.config.key = key
	return c
}

// ConfigSLogClientFun 日志配置
func (c *Client) ConfigSLogClientFun(sLogFun golog.SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
}

// SetHttp 配置请求
func (c *Client) SetHttp(app *gorequest.App) {
	c.requestClient = app
	c.requestClientStatus = true
	c.requestClient.Uri = apiUrl
}

// DefaultHttp 默认请求
func (c *Client) DefaultHttp() {
	c.requestClient = gorequest.NewHttp()
	c.requestClientStatus = true
	c.requestClient.Uri = apiUrl
}
