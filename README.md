# go-log-perf

Performance benchmarking of go based logging libraries. Right now results include -
- [slog](./benchmark/slog) 
- [zerolog](./benchmark/zerolog)

The tests include running sequential as well as parallel benchmark runs, each covering following scenarios -
- An info level logging use case
- Logging with 10 attributes with Context
- Logging with 10 attributes without Context

## slog

`$ go test -bench=. -benchmem -benchtime 1000000x -cpu=1,2,4`

| goos   | goarch | pkg                                   | cpu                                       |
|--------|--------|---------------------------------------|-------------------------------------------|
| darwin | amd64  | github.com/go-log-perf/benchmark/slog | Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz |


| Tests                                               | Number of iterations | Average time taken /per benchmark iteration | Average number of bytes allocated /per benchmark iteration | Average number of allocations /per benchmark iteration |
|-----------------------------------------------------|----------------------|---------------------------------------------|------------------------------------------------------------|--------------------------------------------------------|
| BenchmarkGoSLogInfoSeq                              | 1000000              | 4330 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkGoSLogInfoSeq-2                            | 1000000              | 3962 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkGoSLogInfoSeq-4                            | 1000000              | 3767 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkGoSLogInfoWithTenAttributesWithCtxSeq      | 1000000              | 6756 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkGoSLogInfoWithTenAttributesWithCtxSeq-2    | 1000000              | 6173 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkGoSLogInfoWithTenAttributesWithCtxSeq-4    | 1000000              | 5950 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkGoSLogInfoWithTenAttributesWithoutCtxSeq   | 1000000              | 5795 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkGoSLogInfoWithTenAttributesWithoutCtxSeq-2 | 1000000              | 5375 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkGoSLogInfoWithTenAttributesWithoutCtxSeq-4 | 1000000              | 5437 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogInfoParallel                           | 1000000              | 9669 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkSlogInfoParallel-2                         | 1000000              | 4791 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkSlogInfoParallel-4                         | 1000000              | 2552 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkSlogTenAttributesWithCtxParallel           | 1000000              | 10857 ns/op	                                | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithCtxParallel-2         | 1000000              | 5774 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithCtxParallel-4         | 1000000              | 3310 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithoutCtxParallel        | 1000000              | 8175 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithoutCtxParallel-2      | 1000000              | 4770 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithoutCtxParallel-4      | 1000000              | 2872 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |

## zerolog

`$ go test -bench=. -benchmem -benchtime 1000000x -cpu=1,2,4`

| goos   | goarch | pkg                                      | cpu                                       |
|--------|--------|------------------------------------------|-------------------------------------------|
| darwin | amd64  | github.com/go-log-perf/benchmark/zerolog | Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz |


| Tests                                                | Number of iterations | Average time taken /per benchmark iteration | Average number of bytes allocated /per benchmark iteration | Average number of allocations /per benchmark iteration |
|------------------------------------------------------|----------------------|---------------------------------------------|------------------------------------------------------------|--------------------------------------------------------|
| BenchmarkZeroLogInfoSeq                              | 1000000              | 223.0 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoSeq-2                            | 1000000              | 228.1 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoSeq-4                            | 1000000              | 271.0 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoWithTenAttributesWithCtxSeq      | 1000000              | 2606 ns/op	                                 | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogInfoWithTenAttributesWithCtxSeq-2    | 1000000              | 2341 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogInfoWithTenAttributesWithCtxSeq-4    | 1000000              | 2367 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogInfoWithTenAttributesWithoutCtxSeq   | 1000000              | 2548 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogInfoWithTenAttributesWithoutCtxSeq-2 | 1000000              | 2413 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogInfoWithTenAttributesWithoutCtxSeq-4 | 1000000              | 2388 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogInfoParallel                         | 1000000              | 268.5 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoParallel-2                       | 1000000              | 167.1 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoParallel-4                       | 1000000              | 116.6 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithCtxParallel         | 1000000              | 2550 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithCtxParallel-2       | 1000000              | 1359 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithCtxParallel-4       | 1000000              | 798.6 ns/op                                 | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithoutCtxParallel      | 1000000              | 2577 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithoutCtxParallel-2    | 1000000              | 1329 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithoutCtxParallel-4    | 1000000              | 796.8 ns/op                                 | 624 B/op                                                   | 4 allocs/op                                            |