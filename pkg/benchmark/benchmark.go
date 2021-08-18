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

func (b *benchmarker) Benchmark(queries []Query) *report {
	var wg sync.WaitGroup

	r := NewReport()
	totalStart := time.Now()
	for _, query := range queries {
		wg.Add(1)
		go func(query Query) {
			defer wg.Done()
			start := time.Now()
			req, _ := http.NewRequest("GET", b.formatURL(), nil)
			req.URL.RawQuery = query.formatRangeQuery()
			resp, err := b.client.Do(req)
			if err != nil {
				fmt.Println("Errored when sending request to the server")
				return
			}
			if resp.StatusCode > 200 {
				r.addFailure()
				return
			}
			d := time.Since(start)
			r.append(d)
		}(query)
	}
	wg.Wait()
	r.totalDuration = time.Since(totalStart)
	return r
}

func (b benchmarker) formatURL() string {
	return b.address + "/api/v1/query_range?"
}
