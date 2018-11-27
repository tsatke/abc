[![pipeline status](https://gitlab.com/TimSatke/abc/badges/develop/pipeline.svg)](https://gitlab.com/TimSatke/abc/pipelines)
[![coverage report](https://codecov.io/gl/TimSatke/abc/branch/develop/graphs/badge.svg)](https://codecov.io/gl/TimSatke/abc)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/TimSatke/abc)](https://goreportcard.com/report/gitlab.com/TimSatke/abc)
[![GoDoc](https://godoc.org/gitlab.com/TimSatke/abc?status.svg)](https://godoc.org/gitlab.com/TimSatke/abc)

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
BenchmarkColoredLogger_SimpleLogger_Printf-8             1000000              1064 ns/op             320 B/op          9 allocs/op
BenchmarkColoredLogger_SimpleLogger_Printf-8             1000000              1055 ns/op             320 B/op          9 allocs/op
BenchmarkColoredLogger_SimpleLogger_Printf-8             1000000              1058 ns/op             320 B/op          9 allocs/op
BenchmarkColoredLogger_SimpleLogger_Printf-8             1000000              1060 ns/op             320 B/op          9 allocs/op
BenchmarkColoredLogger_SimpleLogger_Printf-8             1000000              1051 ns/op             320 B/op          9 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               785 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               794 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               790 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               777 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf-8                    2000000               763 ns/op             808 B/op         21 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8          1000000              2590 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2575 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2424 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2479 ns/op            1897 B/op         43 allocs/op
BenchmarkCustomPatternLogger_Printf_Stack_Ops-8           500000              2657 ns/op            1897 B/op         43 allocs/op
BenchmarkNamedLogger_Printf-8                           10000000               238 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                           10000000               236 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                            5000000               243 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                            5000000               238 ns/op             368 B/op          9 allocs/op
BenchmarkNamedLogger_Printf-8                           10000000               250 ns/op             368 B/op          9 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               250 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               249 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               244 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               254 ns/op             304 B/op          8 allocs/op
BenchmarkSimpleLogger_Printf-8                           5000000               274 ns/op             304 B/op          8 allocs/op
BenchmarkStdLogger-8                                     3000000               470 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               469 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               475 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               471 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger-8                                     3000000               470 ns/op              96 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1199 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1183 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1184 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1189 ns/op             176 B/op          2 allocs/op
BenchmarkStdLogger_Stackops-8                            1000000              1213 ns/op             176 B/op          2 allocs/op
PASS
ok      gitlab.com/TimSatke/abc 59.162s
```