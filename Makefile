.PHONY: tools
tools:
	go install github.com/cosmtrek/air@v1.49.0
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1

# Lint, Format
.PHONY: lint-initial
lint-init: tools
	golangci-lint run ./... --timeout=5m

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m

.PHONY: format-initial
format-init: tools
	golangci-lint run ./... --fix

.PHONY: format
format:
	golangci-lint run ./... --fix

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
