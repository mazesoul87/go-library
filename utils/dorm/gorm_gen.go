package dorm

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

// GormGenClientFun *GormClient 驱动
type GormGenClientFun func() *GormGenClient

// GormGenClientTableFun *GormClient 驱动
// string 表名
type GormGenClientTableFun func() (*GormGenClient, string)

type GormGenClientConfig struct {
	Dns    string     // dns地址
	Config gen.Config // gen配置
}

// GormGenClient
// https://gorm.io/zh_CN/gen/index.html
type GormGenClient struct {
	db        *gorm.DB             // 驱动
	generator *gen.Generator       // 驱动
	config    *GormGenClientConfig // 配置
}
