package wechatoffice

import (
	"errors"
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/goredis"
	"gopkg.in/dtapps/go-library.v3/utils/gotime"
)

var (
	qdTypeDb  = "DB"
	qdTypeRdb = "redis"
)

func (app *App) AuthGetAccessTokenMonitor(qdType string) error {
	result := app.GetCallBackIp()
	if len(result.GetCallBackIpResponse.IpList) <= 0 {
		switch qdType {
		case qdTypeDb:
			token := app.AuthGetAccessToken()
			if token.AuthGetAccessTokenResponse.AccessToken == "" {
				return errors.New("获取AccessToken失败")
			} else {
				app.Db.Create(&WechatAccessTokenDbModel{
					AppID:       app.AppId,
					AppSecret:   app.AppSecret,
					AccessToken: token.AuthGetAccessTokenResponse.AccessToken,
					ExpiresIn:   token.AuthGetAccessTokenResponse.ExpiresIn,
					ExpiresTime: gotime.Current().AfterSeconds(7000).Format(),
				})
				return nil
			}
		case qdTypeRdb:
			cacheName := fmt.Sprintf("wechat_access_token:%v", app.AppId)
			redis := goredis.App{
				Rdb: app.RDb,
			}
			token := app.AuthGetAccessToken()
			if token.AuthGetAccessTokenResponse.AccessToken == "" {
				return errors.New("获取AccessToken失败")
			}
			redis.NewStringOperation().Set(cacheName, token.AuthGetAccessTokenResponse.AccessToken, goredis.WithExpire(7000))
			return nil
		default:
			return errors.New("驱动类型不在范围内")
		}
	}
	return nil
}
