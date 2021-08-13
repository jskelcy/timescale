package benchmark

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type benchmarker struct {
	address string
	client  *http.Client
}

func NewBenchmarker(address string) *benchmarker {
	return &benchmarker{
		address: address,
		client:  http.DefaultClient,
	}
}

func (b *benchmarker) Benchmark(queries []Query) report {
	var wg sync.WaitGroup

	r := report{
		reqs: make([]requestInfo, len(queries)),
	}

	totalStart := time.Now()
	for i, query := range queries {
		wg.Add(1)
		go func(i int, query Query) {
			start := time.Now()
			req, _ := http.NewRequest("GET", b.formatURL(), nil)
			req.URL.RawQuery = query.formatRangeQuery()
			resp, err := b.client.Do(req)
			if err != nil {
				fmt.Println("Errored when sending request to the server")
				return
			}
			r.reqs[i].Query = query
			r.reqs[i].Duration = time.Since(start)
			r.reqs[i].StatusCode = resp.StatusCode
			wg.Done()
		}(i, query)
	}
	wg.Wait()
	r.totalDuration = time.Since(totalStart)
	return r
}

func (b benchmarker) formatURL() string {
	return b.address + "/api/v1/query_range?"
}
