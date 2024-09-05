
# юнит тесты
unit_lru:
	go test -v ./lru -run=^$ 

unit_lru_int_pool:
	go test -v ./lru_int_pool -run=^$

unit_lru_pool:
	go test -v ./lru_pool -run=^$


# бенчмарки
bench_lru:
	go test -bench=. ./lru -run=^$

bench_lru_int_pool:
	go test -bench=. ./lru_int_pool -run=^$

bench_lru_pool:
	go test -bench=. ./lru_pool -run=^$


# pprof of BenchmarkRandomAddAndGet in lru_pool
bench_pprof:
	go test -bench=BenchmarkRandomAddAndGet ./lru_pool -run=^$ -cpuprofile=cpu.prof -memprofile=mem.prof -benchmem

# run in web
mem:
	go tool pprof -http=:8080 mem.prof

cpu:
	go tool pprof -http=:8080 cpu.prof

# cmd main
main:
	go run cmd/main.go