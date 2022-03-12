package snowflake

import (
	"fmt"
)

var SnowNode *Node
var ids map[string]*Node
func init() {
	ids = make(map[string]*Node)
	var err error
	SnowNode,err = NewNode(1)
	if err != nil {
		panic(err)
	}
}

func NewIdFactory(name string, nodeId int64) {
	var err error
	if ids[name] != nil {
		panic(fmt.Sprintf("%s name id factory already exist!", name))
	}
	ids[name], err = NewNode(nodeId)
	if err != nil {
		panic(err)
	}
}

const DEFAULT_NAME = "default"
func NewDefaultIdFactory(nodeId int64) {
	NewIdFactory(DEFAULT_NAME, nodeId)
}

func GetIdFactory(name string) *Node {
	return ids[name]
}

func Id() *Node {
	return ids[DEFAULT_NAME]
}