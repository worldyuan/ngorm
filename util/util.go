package util

import (
	"fmt"
	"github.com/jeek120/ngorm/util/snowflake"
	nebula_go "github.com/vesoft-inc/nebula-go/v3"
)

var SnowNode *snowflake.Node
func init() {
	var err error
	SnowNode,err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func CheckResultSet(prefix string, res *nebula_go.ResultSet) {
	if !res.IsSucceed() {
		panic(fmt.Sprintf("%s, ErrorCode: %v, ErrorMsg: %s", prefix, res.GetErrorCode(), res.GetErrorMsg()))
	}
}