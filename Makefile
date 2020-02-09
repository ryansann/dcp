.PHONY: run

run: run-js run-go

run-js:
	@./run-js.sh

run-go:
	@go test -race -v ./...