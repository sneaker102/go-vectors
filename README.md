# go-vectors
package for go arithmetics

## Profiling with Benchmark tests
```
go test -cpuprofile cpu.prof -bench=.
```

And the benchmarks will be run and create a prof file with filename cpu.prof (in the above example).


```
go tool pprof cpu.prof
```

Then within the golang profiler tool, you can use the commands:
```
(pprof) top
```

To list all processes narrowed by a regex search
```
(pprof) peek run*
```

To create a visualization with dot, graphviz needs to be installed.
Then you can visualize the call graph in a browser with the web commands
```
(pprof) web
```

## cgo

Additional Info about integration of C into Go

https://sodocumentation.net/go/topic/6125/cgo
