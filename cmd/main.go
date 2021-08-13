package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jskelcy/promscale-bench/pkg/benchmark"
	"github.com/jskelcy/promscale-bench/pkg/parser"
	"github.com/olekukonko/tablewriter"
)

func main() {
	fileNamePtr := flag.String("input", "", "path to CSV containing sample queries")
	flag.Parse()
	if *fileNamePtr == "" {
		log.Fatalln("no input file value provided")
	}
	file, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Fatalf("Couldn't open the file %v", err)
	}

	p := parser.NewCSVParser(file)
	queries, err := p.Parse()
	if err != nil {
		log.Fatalf("error parsing file %v", err)
	}

	bench := benchmark.NewBenchmarker("http://localhost:9201")
	report := bench.Benchmark(queries)

	p50, err := report.AverageDuration()
	if err != nil {
		log.Fatalf("error generating average for the report %v", err)
	}
	p95, err := report.P95Duration()
	if err != nil {
		log.Fatalf("error generating p95 for the report %v", err)
	}
	data := [][]string{
		{"Requests Count", fmt.Sprintf("%d", report.NumRequest())},
		{"Total Duration", fmt.Sprintf("%d", report.TotalTime().Milliseconds())},
		{"Average Duration", fmt.Sprintf("%f", p50)},
		{"P95 Duration", fmt.Sprintf("%f", p95)},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	slowest := report.MaxDuration()
	fmt.Printf("Slowest request %s took %v\n", slowest.Query.Expression, slowest.Duration)
	fastest := report.MinDuration()
	fmt.Printf("Fastest request %s took %v\n", fastest.Query.Expression, fastest.Duration)
}
