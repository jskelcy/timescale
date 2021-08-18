build:
	go build -o bin/benchmark cmd/main.go

run:
	go run cmd/main.go --input=$(input)