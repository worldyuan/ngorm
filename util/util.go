package util

import (
	"fmt"
	nebula_go "github.com/vesoft-inc/nebula-go/v3"
)

func CheckResultSet(prefix string, res *nebula_go.ResultSet) {
	if !res.IsSucceed() {
		panic(fmt.Sprintf("%s, ErrorCode: %v, ErrorMsg: %s", prefix, res.GetErrorCode(), res.GetErrorMsg()))
	}
}