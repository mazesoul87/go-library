package gojobs

import (
	"context"
	"fmt"
	"github.com/mazesoul87/go-library/utils/gorequest"
	"github.com/mazesoul87/go-library/utils/gotime"
	"time"
)

// GetRedisKeyName 获取Redis键名
func GetRedisKeyName(taskType string) string {
	return "task:run:" + taskType
}

// SetRedisKeyValue 返回设置Redis键值
func SetRedisKeyValue(ctx context.Context, taskType string) (context.Context, string, any, time.Duration) {
	return ctx,
		GetRedisKeyName(taskType),
		fmt.Sprintf(
			"%s-%s-%s-%s",
			gotime.Current().SetFormat(gotime.DateTimeZhFormat),
			gorequest.TraceGetTraceID(ctx),
			gorequest.TraceGetSpanID(ctx),
			gorequest.GetRequestIDContext(ctx),
		),
		0
}

// SetRedisKeyValueExpiration 返回设置Redis键值，有过分时间
func SetRedisKeyValueExpiration(ctx context.Context, taskType string, expiration int64) (context.Context, string, any, time.Duration) {
	return ctx,
		GetRedisKeyName(taskType),
		fmt.Sprintf(
			"%s-%s-%s-%s",
			gotime.Current().SetFormat(gotime.DateTimeZhFormat),
			gorequest.TraceGetTraceID(ctx),
			gorequest.TraceGetSpanID(ctx),
			gorequest.GetRequestIDContext(ctx),
		),
		time.Duration(expiration)
}
