package main

import (
	"flag"
	"log"
	"os"

	"github.com/jskelcy/promscale-bench/pkg/benchmark"
	"github.com/jskelcy/promscale-bench/pkg/parser"
)

func main() {
	fileNamePtr := flag.String("input", "", "path to CSV containing sample queries")
	promscaleURL := flag.String("url", "http://localhost:9201", "url of promscale if not default")
	flag.Parse()
	if *fileNamePtr == "" {
		log.Fatalln("no input file value provided")
	}
	file, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Fatalf("Couldn't open the file %v", err)
	}
	if *promscaleURL == "" {
		*promscaleURL = "http://localhost:9201"
	}

	p := parser.NewCSVParser(file)
	queries, err := p.Parse()
	if err != nil {
		log.Fatalf("error parsing file %v", err)
	}

	bench := benchmark.NewBenchmarker(*promscaleURL)
	report := bench.Benchmark(queries)
	report.Render()
}
