package benchmark

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"
)

const (
	minDuration time.Duration = -1 << 63
	maxDuration time.Duration = 1<<63 - 1
)

type report struct {
	sync.Mutex
	vals          stats.Float64Data
	totalDuration time.Duration
	min           time.Duration
	max           time.Duration
}

func NewReport() *report {
	return &report{
		vals: stats.Float64Data{},
		min:  maxDuration,
		max:  minDuration,
	}
}

func (r *report) append(d time.Duration) {
	r.Lock()
	defer r.Unlock()
	r.vals = append(r.vals, float64(d.Milliseconds()))
	if d < r.min {
		r.min = d
	}
	if d > r.max {
		r.max = d
	}
}

func (r *report) Render() {
	mean, err := stats.Mean(r.vals)
	if err != nil {
		log.Fatalf("error generating average for the report %v", err)
	}
	median, err := stats.Median(r.vals)
	if err != nil {
		log.Fatalf("error generating median for the report %v", err)
	}
	p95, err := stats.Percentile(r.vals, 95)
	if err != nil {
		log.Fatalf("error generating p95 for the report %v", err)
	}
	data := [][]string{
		{"Requests Count", fmt.Sprintf("%d", len(r.vals))},
		{"Total Duration", fmt.Sprintf("%d ms", r.totalDuration.Milliseconds())},
		{"Average Duration", fmt.Sprintf("%d ms", int(mean))},
		{"Median Duration", fmt.Sprintf("%d ms", int(median))},
		{"Minimum Duration", fmt.Sprintf("%d ms", r.min.Milliseconds())},
		{"Max Duration", fmt.Sprintf("%d ms", r.max.Milliseconds())},
		{"P95 Duration", fmt.Sprintf("%d ms", int(p95))},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	for _, v := range data {
		table.Append(v)
	}
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
