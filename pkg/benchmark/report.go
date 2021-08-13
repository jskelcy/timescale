package benchmark

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"
)

const (
	minDuration time.Duration = -1 << 63
	maxDuration time.Duration = 1<<63 - 1
)

type report struct {
	reqs          []requestInfo
	totalDuration time.Duration
}

func (r report) minDuration() requestInfo {
	min := requestInfo{Duration: maxDuration}
	for _, info := range r.reqs {
		if info.Duration < min.Duration {
			min = info
		}
	}
	return min
}

func (r report) maxDuration() requestInfo {
	min := requestInfo{Duration: minDuration}
	for _, info := range r.reqs {
		if info.Duration > min.Duration {
			min = info
		}
	}
	return min
}

func (r report) averageDuration() (float64, error) {
	vals := stats.Float64Data{}
	for _, req := range r.reqs {
		vals = append(vals, float64(req.Duration.Milliseconds()))
	}
	avg, err := stats.Mean(vals)
	if err != nil {
		return 0, err
	}
	return avg, nil
}

func (r report) medianDuration() (float64, error) {
	vals := stats.Float64Data{}
	for _, req := range r.reqs {
		vals = append(vals, float64(req.Duration.Milliseconds()))
	}
	avg, err := stats.Median(vals)
	if err != nil {
		return 0, err
	}
	return avg, nil
}

func (r report) p95Duration() (float64, error) {
	vals := stats.Float64Data{}
	for _, req := range r.reqs {
		vals = append(vals, float64(req.Duration.Milliseconds()))
	}
	avg, err := stats.Percentile(vals, 95)
	if err != nil {
		return 0, err
	}
	return avg, nil
}

func (r report) TotalTime() time.Duration {
	return r.totalDuration
}

func (r report) NumRequest() int {
	return len(r.reqs)
}

func (r report) Render() {
	p50, err := r.averageDuration()
	if err != nil {
		log.Fatalf("error generating average for the report %v", err)
	}
	median, err := r.medianDuration()
	if err != nil {
		log.Fatalf("error generating median for the report %v", err)
	}
	p95, err := r.p95Duration()
	if err != nil {
		log.Fatalf("error generating p95 for the report %v", err)
	}
	data := [][]string{
		{"Requests Count", fmt.Sprintf("%d", r.NumRequest())},
		{"Total Duration", fmt.Sprintf("%d ms", r.TotalTime().Milliseconds())},
		{"Average Duration", fmt.Sprintf("%d ms", int(p50))},
		{"Median Duration", fmt.Sprintf("%d ms", int(median))},
		{"Minimum Duration", fmt.Sprintf("%d ms", r.minDuration().Duration.Milliseconds())},
		{"Max Duration", fmt.Sprintf("%d ms", r.maxDuration().Duration.Milliseconds())},
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

type requestInfo struct {
	Query      Query
	Duration   time.Duration
	StatusCode int
}
