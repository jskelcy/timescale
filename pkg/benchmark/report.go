package benchmark

import (
	"time"

	"github.com/montanaflynn/stats"
)

const (
	minDuration time.Duration = -1 << 63
	maxDuration time.Duration = 1<<63 - 1
)

type report struct {
	reqs          []requestInfo
	totalDuration time.Duration
}

func (r report) MinDuration() requestInfo {
	min := requestInfo{Duration: maxDuration}
	for _, info := range r.reqs {
		if info.Duration < min.Duration {
			min = info
		}
	}
	return min
}

func (r report) MaxDuration() requestInfo {
	min := requestInfo{Duration: minDuration}
	for _, info := range r.reqs {
		if info.Duration > min.Duration {
			min = info
		}
	}
	return min
}

func (r report) AverageDuration() (float64, error) {
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

func (r report) P95Duration() (float64, error) {
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

type requestInfo struct {
	Query      Query
	Duration   time.Duration
	StatusCode int
}
