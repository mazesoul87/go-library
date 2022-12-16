package pintoto

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
