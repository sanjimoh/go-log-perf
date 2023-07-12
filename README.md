# go-log-perf

## slog

`$ go test -bench=. -benchmem -benchtime 1000000x -cpu=1,2,4`

| goos   | goarch | pkg                                   | cpu                                       |
|--------|--------|---------------------------------------|-------------------------------------------|
| darwin | amd64  | github.com/go-log-perf/benchmark/slog | Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz |


| Tests                                  | Number of iterations | Average time taken /per benchmark iteration | Average number of bytes allocated /per benchmark iteration | Average number of allocations /per benchmark iteration |
|----------------------------------------|----------------------|---------------------------------------------|------------------------------------------------------------|--------------------------------------------------------|
| BenchmarkSlogInfoPositive              | 1000000              | 6002 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkSlogInfoPositive-2            | 1000000              | 3591 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkSlogInfoPositive-4            | 1000000              | 2227 ns/op                                  | 568 B/op                                                   | 6 allocs/op                                            |
| BenchmarkSlogInfoNegative              | 1000000              | 26.78 ns/op	                                | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkSlogInfoNegative-2            | 1000000              | 36.25 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkSlogInfoNegative-4            | 1000000              | 36.98 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkSlogTenAttributesWithCtx      | 1000000              | 10152 ns/op                                 | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithCtx-2    | 1000000              | 6087 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithCtx-4    | 1000000              | 3685 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithoutCtx   | 1000000              | 8760 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithoutCtx-2 | 1000000              | 5580 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkSlogTenAttributesWithoutCtx-4 | 1000000              | 3311 ns/op                                  | 776 B/op                                                   | 7 allocs/op                                            |
| BenchmarkGoSLog                        | 1000000              | 4442 ns/op                                  | 392 B/op                                                   | 5 allocs/op                                            |
| BenchmarkGoSLog-2                      | 1000000              | 4178 ns/op                                  | 392 B/op                                                   | 5 allocs/op                                            |
| BenchmarkGoSLog-4                      | 1000000              | 4117 ns/op                                  | 392 B/op                                                   | 5 allocs/op                                            |
| BenchmarkGoSLogWithFields              | 1000000              | 5057 ns/op                                  | 392 B/op                                                   | 5 allocs/op                                            |
| BenchmarkGoSLogWithFields-2            | 1000000              | 4761 ns/op                                  | 392 B/op                                                   | 5 allocs/op                                            |
| BenchmarkGoSLogWithFields-4            | 1000000              | 4782 ns/op                                  | 392 B/op                                                   | 5 allocs/op                                            |

## zerolog

`$ go test -bench=. -benchmem -benchtime 1000000x -cpu=1,2,4`

| goos   | goarch | pkg                                      | cpu                                       |
|--------|--------|------------------------------------------|-------------------------------------------|
| darwin | amd64  | github.com/go-log-perf/benchmark/zerolog | Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz |


| Tests                                     | Number of iterations | Average time taken /per benchmark iteration | Average number of bytes allocated /per benchmark iteration | Average number of allocations /per benchmark iteration |
|-------------------------------------------|----------------------|---------------------------------------------|------------------------------------------------------------|--------------------------------------------------------|
| BenchmarkZeroLogInfoPositive              | 1000000              | 203.5 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoPositive-2            | 1000000              | 111.2 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoPositive-4            | 1000000              | 59.30 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoNegative              | 1000000              | 11.00 ns/op	                                | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoNegative-2            | 1000000              | 17.78 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogInfoNegative-4            | 1000000              | 18.76 ns/op                                 | 0 B/op                                                     | 0 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithCtx      | 1000000              | 2323 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithCtx-2    | 1000000              | 1192 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithCtx-4    | 1000000              | 761.5 ns/op                                 | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithoutCtx   | 1000000              | 2379 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithoutCtx-2 | 1000000              | 1255 ns/op                                  | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLogTenAttributesWithoutCtx-4 | 1000000              | 836.0 ns/op                                 | 624 B/op                                                   | 4 allocs/op                                            |
| BenchmarkZeroLog                          | 1000000              | 44379 ns/op                                 | 2808 B/op                                                  | 45 allocs/op                                           |
| BenchmarkZeroLog-2                        | 1000000              | 41234 ns/op                                 | 2808 B/op                                                  | 45 allocs/op                                           |
| BenchmarkZeroLog-4                        | 1000000              | 32881 ns/op                                 | 2809 B/op                                                  | 45 allocs/op                                           |
| BenchmarkZeroLogWithFields                | 1000000              | 32881 ns/op                                 | 3072 B/op                                                  | 60 allocs/op                                           |
| BenchmarkZeroLogWithFields-2              | 1000000              | 32881 ns/op                                 | 3072 B/op                                                  | 60 allocs/op                                           |
| BenchmarkZeroLogWithFields-4              | 1000000              | 32881 ns/op                                 | 3073 B/op                                                  | 60 allocs/op                                           |