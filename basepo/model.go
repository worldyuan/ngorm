package basepo

import (
	"github.com/jeek120/ngorm/util/snowflake"
)

type ITag interface {
	TagName() string
	Id() int64
}

type IEdge interface {
	EdgeName() string
	Src() int64
	Dst() int64
	Rank() int
}

type Tag struct {
	id int64
}

type Edge struct {
	src  int64
	dst  int64
	rank int
}

func (t *Tag) GenId() int64 {
	t.id = snowflake.Id().Generate().Int64()
	return t.id
}

func (t *Tag) SetId(id int64) {
	t.id = id
}

func (t *Tag) Id() int64 {
	return t.id
}

func (t *Tag) Id2() int64 {
	if t.id == 0 {
		t.GenId()
	}
	return t.id
}

func NewEdge(src, dst int64) *Edge {
	return &Edge{
		src: src,
		dst: dst,
	}
}

func NewEdgeWithRank(src, dst int64, rank int) *Edge {
	return &Edge{
		src:  src,
		dst:  dst,
		rank: rank,
	}
}

func (e *Edge) Src() int64 {
	return e.src
}

func (e *Edge) Dst() int64 {
	return e.dst
}

func (e *Edge) Rank() int {
	return e.rank
}
