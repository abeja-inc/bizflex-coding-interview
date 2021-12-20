.PHONY: test clean bench

test:
	go test -v ./...

bench:
	go test -test.bench BenchmarkKVS -benchmem

clean:
	go clean -testcache