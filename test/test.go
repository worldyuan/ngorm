package test

import (
	"github.com/jeek120/ngorm/basepo"
)

type User struct {
	*basepo.Tag
	// 名称
	Name string				`json:"name" idx:"name(10)"`	// 名称
	Passwd string
	Age 	int64
}

type (
	Group struct {
		*basepo.Tag
		name string			// 名称
	}
)

type (
	UserGroup struct {
		*basepo.Edge			// 用户所属群组
	}
)