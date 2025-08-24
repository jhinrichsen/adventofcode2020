.PHONY: all
all: lint test

.PHONY: bench
bench:
	go test -run=^$ -bench=. -benchmem -count=10

.PHONY: lint
lint:
	go vet
	staticcheck

.PHONY: test
test:
	go test -v -run=Day..Part.$

.PHONY: gopls
gopls:
	# go run golang.org/x/tools/gopls@latest serve -listen 127.0.0.1:4389
	go run golang.org/x/tools/gopls@latest serve -listen 127.0.0.1:4389 -rpc.trace -logfile gopls.log

prof:
	go test -run=^$ -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	go tool pprof cpu.profile
	# go tool pprof mem.profile
