package parser

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/jskelcy/promscale-bench/pkg/benchmark"
)

type Parser interface {
	Prase() ([]benchmark.Query, error)
}

type csvParser struct {
	csvReader *csv.Reader
}

func NewCSVParser(f io.Reader) csvParser {
	r := csv.NewReader(f)
	r.Comma = '|'
	r.LazyQuotes = true
	return csvParser{
		csvReader: r,
	}
}

func (p *csvParser) Parse() ([]benchmark.Query, error) {
	out := []benchmark.Query{}
	for {
		// Read each record from csv
		record, err := p.csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) != 4 {
			continue
		}

		q := benchmark.Query{
			Expression: record[0],
		}
		if q.Start, err = strconv.Atoi(record[1]); err != nil {
			continue
		}
		if q.End, err = strconv.Atoi(record[2]); err != nil {
			continue
		}
		if q.Step, err = strconv.Atoi(record[3]); err != nil {
			continue
		}
		out = append(out, q)
	}

	return out, nil
}
