package gorequest

import (
	"context"
	"github.com/mazesoul87/go-library/gostring"
	"strings"
)

func getIPV6_SpeedNeu6EduCn(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://speed.neu6.edu.cn/getIP.php")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	return gostring.SpaceAndLineBreak(string(response.ResponseBody))
}

func getIPV6_V6IdentMe(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://v6.ident.me")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	return strings.Replace(strings.Replace(string(response.ResponseBody), " ", "", -1), "\n", "", -1)
}

func getIPV6_6IpwCn(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://6.ipw.cn")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	return strings.Replace(strings.Replace(string(response.ResponseBody), " ", "", -1), "\n", "", -1)
}
