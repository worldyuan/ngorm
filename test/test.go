package test

import "github.com/jeek120/ngorm/base"

type User struct {
	*base.Tag
	// 名称
	Name string				// 名称
	Passwd string 			// 密码
	Age 	int64			// 年龄
}

type (
	Group struct {
		*base.Tag
		name string			// 名称
	}
)

type (
	UserGroup struct {
		*base.Edge			// 用户所属群组
	}
)