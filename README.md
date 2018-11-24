# abc

A Golang logging library.

## Installation
```
go get -u gitlab.com/TimSatke/abc
```

## Example
`abc` supports several types of loggers. For better examples and more loggers, see `examples/*` or the wiki.

### SimpleLogger
```go
logger := abc.NewSimpleLogger()
logger.SetLevel(abc.LevelDebug)
logger.Debugf("Hello %v!", "World") // 2018-11-24 20:10:55.300 [DEBG] - Hello World
```

### NamedLogger
```go
logger := abc.NewNamedLogger("MyLogger")
logger.SetLevel(abc.LevelDebug)
logger.Debugf("Hello %v!", "World") // 2018-11-24 20:10:55.300 <MyLogger> [DEBG] - Hello World
```

## Benchmarks
```
$ go test -count 5 -bench . -benchmem
goos: windows
goarch: amd64
pkg: gitlab.com/TimSatke/abc
BenchmarkCustomPatternLogger_Printf-8                    2000000               813 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               796 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               805 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               807 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               795 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8          1000000              2597 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2582 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2541 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2550 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2641 ns/op            1897 B/op         43 allocs/op
BenchmarkNamedLogger_Printf-8                            5000000               241 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                           10000000               244 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                           10000000               245 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                            5000000               245 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                            5000000               245 ns/op             368 B/op          9 allocs/op
BenchmarkSimpleLogger_Printf-8                          10000000               230 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                          10000000               229 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                          10000000               224 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               229 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               225 ns/op             304 B/op          8 allocs/op
BenchmarkStdLogger-8                                     3000000               478 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               481 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               473 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               472 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               472 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1205 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1211 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1213 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1206 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1216 ns/op             176 B/op          2 allocs/op
PASS
ok      gitlab.com/TimSatke/abc 56.298s
```