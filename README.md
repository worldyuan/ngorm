## Nebula Graph ORM框架


NGORM 是一个专注性能的Nebula Graph的ORM框架，根据PO（序列化实体）自动生成数据库操作相关函数，更高效的完成数据库操作。

#### NGORM有如下主要特点
- 极致的性能
- 简易上手
- 事先使用部分约束，以降低程序的复杂度
- 自动生成创建tag和edge，以及增删改查
- 无侵入式自动生成代码

**NGORM交流群**
```
正在创建中...
```

## 快速安装

**1. 安装工具和依赖**

```shell
go install github.com/jeek120/ngorm/cmd/ngormgen
go get github.com/jeek120/ngorm
```

**2. 编写实体**

```go
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
```

**3.通过命令生成代码**

```shell
ngormgen ./test
```

**4.通过注解生成**

```go
package main

//go:generate ngormgen ../po
func main() {
	
}
```
