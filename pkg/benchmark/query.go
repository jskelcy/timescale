package benchmark

import (
	"fmt"
	"net/url"
	"time"
)

type Query struct {
	Expression string
	Start      int
	End        int
	Step       int
}

func (q Query) formatRangeQuery() string {
	params := url.Values{}
	params.Add("query", q.Expression)
	params.Add("start", time.Unix(0, int64(q.Start)*int64(time.Millisecond)).Format(time.RFC3339))
	params.Add("end", time.Unix(0, int64(q.End)*int64(time.Millisecond)).Format(time.RFC3339))
	params.Add("step", fmt.Sprintf("%d", q.Step))
	return params.Encode()
}
