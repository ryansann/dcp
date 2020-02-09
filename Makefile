.PHONY: run

run: run-js run-go

run-js:
	@echo "JS TESTS:"
	@./run-js.sh

run-go:
	@echo "GO TESTS:"
	@go test -race -v ./...