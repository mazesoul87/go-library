package dorm

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func NewXormPostgresClient(config *XormClientConfigXorm) (*XormClient, error) {

	var err error
	c := &XormClient{config: config}

	c.db, err = xorm.NewEngine("postgres", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	return c, nil
}
