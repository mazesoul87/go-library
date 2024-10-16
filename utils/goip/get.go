package goip

import (
	"github.com/mazesoul87/go-library/utils/goip/geoip"
	"github.com/mazesoul87/go-library/utils/goip/qqwry"
)

func (c *Client) GetGeo() *geoip.Client {
	return c.geoIpClient
}

func (c *Client) GetQqWry() *qqwry.Client {
	return c.qqwryClient
}
