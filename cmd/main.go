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
	report.Render()
}
