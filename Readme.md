# Promscale Benchmark

Promscale Benchmark is a golang application which benchmarks a Promscale based on a provided set of queries.

## Usage

In order to run the following commands please use go version `1.13.15` or greater.

Running the project assumes Promscale's Prometheus HTTP API is exposed at `localhost:9201`. Follow the Promscale docker installation guide for the simplest setup. Running the application using default configuration can be done as follows:
```
make run
```
The `input` argument can be set providing CSV file containing desired queries, if no `input` is provided the `obs-queries.csv` file will be used by default.:
```
make run input=obs-queries.csv
```
If the Promscale Prometheus HTTP API is not exposed at `localhost:9201` the URL can be overwritten using the `url` argument, for example:
```
make run input=obs-queries.csv url=example.com
```


Compiling this project can be done using the following command, the artifact will be placed in at the 
root of the project under `bin/benchmark`. Compilation is not required to run the project.
```
$ make build
```


