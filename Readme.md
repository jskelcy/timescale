# Promscale Benchmark

Promscale Benchmark is a go application which benchmarks a Promscale based on a provided set of queries.

## Usage

Running the project assumes Promscale is exposed locally at 9201. Follow the Promscale docker installation guide for the simplest setup. Once Promscale is available use the following command to run a benchmark, where `input` is the CSV file containing desired queries:
```
make run input=obs-queries.csv
```
If Promscale is not exposed at `localhost:9201` the URL can be overwritten using the `url` argument, for example:
```
make run input=obs-queries.csv url=example.com
```


Compiling this project can be done using the following command, the artifact will be placed in at the 
root of the project under `bin/benchmark`. Compilation is not required to run the project.
```
$ make build
```


