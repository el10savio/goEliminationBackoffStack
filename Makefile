
fmt:
	@echo "go fmt goEliminationBackoffStack"
	go fmt ./...

vet:
	@echo "go vet goEliminationBackoffStack"
	go vet ./...

lint:
	@echo "go lint goEliminationBackoffStack"
	golint ./...

test:
	@echo "Testing goEliminationBackoffStack"
	go test -v -race --cover ./...

bench:
	@echo "Running benchmarks on goEliminationBackoffStack"
	go test -bench=. -count=15 ./...

codespell:
	@echo "checking app spellings"
	codespell
