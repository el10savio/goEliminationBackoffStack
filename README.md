# goEliminationBackoffStack

A basic elimination backoff stack without locks to test its performance.

### Tests:

```
$ go test -v -race --cover ./...
Testing goEliminationBackoffStack
go test -v -race --cover ./...
=== RUN   TestStack
--- PASS: TestStack (0.00s)
=== RUN   TestRaceConditions
--- PASS: TestRaceConditions (0.10s)
=== RUN   TestThroughput
    stack_test.go:185: Throughput: 767751.20 ops/sec with 10 goroutines
--- PASS: TestThroughput (5.01s)
PASS
coverage: 100.0% of statements
ok  	github.com/el10savio/goEliminationBackoffStack/eliminationBackoffStack	6.309s	coverage: 100.0% of statements
```

### Benchmarks:

```
$ go test -bench=. -count=15 ./...
Running benchmarks on goEliminationBackoffStack
go test -bench=. -count=15 ./...
goos: darwin
goarch: arm64
pkg: github.com/el10savio/goEliminationBackoffStack/eliminationBackoffStack
cpu: Apple M4
BenchmarkPushSingleThreaded-10    	31228603	        46.11 ns/op
BenchmarkPushSingleThreaded-10    	32533317	        44.68 ns/op
BenchmarkPushSingleThreaded-10    	30721302	        37.54 ns/op
BenchmarkPushSingleThreaded-10    	31005693	        45.73 ns/op
BenchmarkPushSingleThreaded-10    	33576772	        37.77 ns/op
BenchmarkPushSingleThreaded-10    	30530712	        36.09 ns/op
BenchmarkPushSingleThreaded-10    	31964174	        37.04 ns/op
BenchmarkPushSingleThreaded-10    	30296348	        40.89 ns/op
BenchmarkPushSingleThreaded-10    	31402711	        42.42 ns/op
BenchmarkPushSingleThreaded-10    	31694659	        42.83 ns/op
BenchmarkPushSingleThreaded-10    	29026586	        36.91 ns/op
BenchmarkPushSingleThreaded-10    	32637966	        35.04 ns/op
BenchmarkPushSingleThreaded-10    	40670266	        35.30 ns/op
BenchmarkPushSingleThreaded-10    	40151292	        35.65 ns/op
BenchmarkPushSingleThreaded-10    	41294170	        35.21 ns/op
BenchmarkPopSingleThreaded-10     	202944169	         8.534 ns/op
BenchmarkPopSingleThreaded-10     	204272269	        12.14 ns/op
BenchmarkPopSingleThreaded-10     	100000000	        12.00 ns/op
BenchmarkPopSingleThreaded-10     	202868389	        15.10 ns/op
BenchmarkPopSingleThreaded-10     	203625249	        11.92 ns/op
BenchmarkPopSingleThreaded-10     	201389193	        14.95 ns/op
BenchmarkPopSingleThreaded-10     	205082251	        12.19 ns/op
BenchmarkPopSingleThreaded-10     	201312318	        12.73 ns/op
BenchmarkPopSingleThreaded-10     	175292604	         9.957 ns/op
BenchmarkPopSingleThreaded-10     	163426416	        10.51 ns/op
BenchmarkPopSingleThreaded-10     	142790118	         8.285 ns/op
BenchmarkPopSingleThreaded-10     	199385780	        15.73 ns/op
BenchmarkPopSingleThreaded-10     	198049010	        16.96 ns/op
BenchmarkPopSingleThreaded-10     	198732843	        15.67 ns/op
BenchmarkPopSingleThreaded-10     	202238186	        13.67 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	21942907	        49.85 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24377856	        50.09 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	21464568	        51.26 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24002659	        54.06 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	23491966	        49.61 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	26738240	        51.63 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24344762	        48.33 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20478513	        57.61 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	23075868	        48.26 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24035131	        48.43 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24302894	        48.22 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24443485	        53.83 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20229988	        52.33 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	24549102	        48.07 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	21883300	        51.86 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 9266799	       142.9 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 7940444	       145.1 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8634070	       139.9 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 7634390	       162.9 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8774962	       139.2 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 9032847	       132.5 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 9471768	       121.4 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	11967373	       123.0 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8346550	       146.7 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	10046444	       142.3 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8021979	       142.2 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8226784	       140.8 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8733534	       132.8 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 8396665	       157.5 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	 7963494	       137.4 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5507420	       240.9 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5366869	       242.9 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6892708	       184.6 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6666822	       208.7 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 7238730	       208.5 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 7242514	       177.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 7669486	       182.1 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 9693320	       202.5 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 7812037	       195.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 9866246	       186.1 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5491891	       195.5 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6655484	       205.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6878980	       221.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6812278	       201.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 7528482	       200.9 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7956192	       176.3 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 6961665	       176.3 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7713139	       179.8 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8540167	       173.5 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7126372	       185.8 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7147305	       167.5 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8129839	       178.7 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8208976	       180.3 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7068364	       199.0 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7403941	       191.1 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7572820	       183.9 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7049458	       173.5 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8157267	       182.6 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 9000380	       141.5 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8404812	       141.4 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	18752624	        58.51 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	21273027	        58.19 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	23962777	        59.82 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	24262602	        58.93 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	24683823	        59.55 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	21846681	        60.09 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	19378018	        57.61 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	20811368	        57.65 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	21459738	        54.68 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	20732251	        57.69 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	22189125	        58.13 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	19811596	        52.70 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	21624414	        55.18 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	20601771	        59.09 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	19982238	        55.84 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18620829	        68.47 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17231053	        67.62 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18730244	        71.46 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18066789	        70.07 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17395232	        70.88 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18085214	        72.73 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17454946	        73.89 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17735223	        71.81 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17215365	        68.83 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17936035	        70.90 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17853080	        70.78 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	16888750	        71.63 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17501746	        71.88 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19757869	        71.08 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	20813578	        71.57 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	17213163	        73.99 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	17527779	        77.69 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	17391933	        88.30 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13544494	        95.24 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	12465070	        90.95 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13825228	        97.75 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	12893572	        86.61 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13404363	        95.99 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	12799988	        91.56 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13915018	        95.29 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	12955033	        91.05 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13886016	        96.73 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	12758638	        89.55 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13171879	        91.35 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	13459981	        96.20 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	13683692	        88.95 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	13540635	        81.19 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	15708304	        75.23 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	17263909	        74.87 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	17484247	        74.07 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	18253159	        76.96 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	17704032	        72.33 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	16950090	        74.69 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	16614620	        72.46 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	16515171	        74.34 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	17909767	        73.75 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	15384754	        74.62 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	17505160	        72.61 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	16079072	        74.33 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	15455120	        75.11 ns/op
BenchmarkLatency-10                             	14910447	        78.94 ns/op	        42.00 p50-ns	        42.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	13309081	        84.21 ns/op	        42.00 p50-ns	        83.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	12345704	        90.56 ns/op	        42.00 p50-ns	        83.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	10714360	        95.28 ns/op	        42.00 p50-ns	        83.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	12222147	        98.43 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	11491060	       100.8 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	10016038	       100.9 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	11835744	        99.07 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	10043031	       102.5 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	12278322	        94.58 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	10928902	        97.91 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	13832852	        95.29 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	12629894	        97.77 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	12059725	        93.03 ns/op	        42.00 p50-ns	        83.00 p90-ns	       125.0 p99-ns
BenchmarkLatency-10                             	10806124	        97.14 ns/op	        42.00 p50-ns	        84.00 p90-ns	       125.0 p99-ns
PASS
ok  	github.com/el10savio/goEliminationBackoffStack/eliminationBackoffStack	514.353s
```


### Compared to goTreiberStack Implementation

Throughput of goEliminationBackoffStack:
```
=== RUN   TestThroughput
    stack_test.go:185: Throughput: 767751.20 ops/sec with 10 goroutines
--- PASS: TestThroughput (5.00s)
```

Throughput of goTreiberStack:
```
=== RUN   TestThroughput
    stack_test.go:185: Throughput: 372960.40 ops/sec with 10 goroutines
--- PASS: TestThroughput (5.00s)
```

Benchstat Comparisons:
```
goos: darwin
goarch: arm64
pkg: github.com/el10savio/goEliminationBackoffStack/eliminationBackoffStack
cpu: Apple M4
                                   │   new.txt    │
                                   │    sec/op    │
PushSingleThreaded-10                37.54n ± 14%
PopSingleThreaded-10                 12.19n ± 24%
ConcurrentPush/goroutines-2-10       50.09n ±  4%
ConcurrentPush/goroutines-4-10       140.8n ±  6%
ConcurrentPush/goroutines-8-10       201.3n ±  8%
ConcurrentPush/goroutines-16-10      178.7n ±  3%
ConcurrentPushPop/goroutines-2-10    58.13n ±  4%
ConcurrentPushPop/goroutines-4-10    71.08n ±  1%
ConcurrentPushPop/goroutines-8-10    91.35n ±  5%
ConcurrentPushPop/goroutines-16-10   74.62n ±  1%
Latency-10                           97.14n ±  4%
geomean                              73.08n

           │   new.txt   │
           │   p50-sec   │
Latency-10   42.00n ± 0%

           │   new.txt   │
           │   p90-sec   │
Latency-10   84.00n ± 1%

           │   new.txt   │
           │   p99-sec   │
Latency-10   125.0n ± 0%

pkg: github.com/el10savio/goTreiberStack/treiber
                                   │   old.txt    │
                                   │    sec/op    │
PushSingleThreaded-10                31.89n ±  1%
PopSingleThreaded-10                 10.92n ± 30%
ConcurrentPush/goroutines-2-10       60.34n ±  5%
ConcurrentPush/goroutines-4-10       86.55n ±  4%
ConcurrentPush/goroutines-8-10       244.9n ± 21%
ConcurrentPush/goroutines-16-10      167.5n ± 10%
ConcurrentPushPop/goroutines-2-10    38.30n ±  2%
ConcurrentPushPop/goroutines-4-10    65.80n ±  3%
ConcurrentPushPop/goroutines-8-10    181.5n ±  1%
ConcurrentPushPop/goroutines-16-10   223.9n ±  1%
Latency-10                           71.39n ±  1%
geomean                              76.76n

           │   old.txt   │
           │   p50-sec   │
Latency-10   41.00n ± 0%

           │   old.txt   │
           │   p90-sec   │
Latency-10   42.00n ± 0%

           │   old.txt   │
           │   p99-sec   │
Latency-10   42.00n ± 0%

```
