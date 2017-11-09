# go-graph

Implementation of graph datastructure in golang.

## Features
1. Directed Graph implemented.
2. Concurrent Directed Graph implemented.

## Getting Started

### Installing

```
go get github.com/BlackRabbitt/go-graph
```

### Import in your project

```
import "github.com/BlackRabbitt/go-graph/digraph"
```

### Run Test and Benchmark

To run the tests:
```
go test
```

To run the benchmarks:
```
go test -test.bench 
```

Test a memory consumption and an allocations count. 
```
go test -bench=. -benchmem
```
