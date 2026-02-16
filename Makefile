SHELL := /bin/bash

.PHONY: help fmt test race tidy run lint proto tools

help:
	@echo "Targets:"
	@echo "  fmt     - gofmt (+ goimports if installed)"
	@echo "  test    - run unit tests"
	@echo "  race    - run tests with race detector"
	@echo "  tidy    - go mod tidy"
	@echo "  run     - run local service entrypoint"
	@echo "  lint    - placeholder (we’ll wire golangci-lint later)"
	@echo "  proto   - placeholder (we’ll wire Buf + protoc-gen-go later)"
	@echo "  tools   - install dev tools (optional)"

fmt:
	@echo "==> gofmt"
	gofmt -w .
	@echo "==> goimports (if available)"
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w . ; \
	else \
		echo "goimports not installed (ok for Day 1). Install via: make tools"; \
	fi

test:
	go test ./...

race:
	go test -race ./...

tidy:
	go mod tidy

run:
	go run ./cmd/todo

lint:
	@echo "lint not wired yet (Day 1 placeholder)."

proto:
	@echo "proto not wired yet (Day 1 placeholder)."

tools:
	@echo "Installing goimports..."
	go install golang.org/x/tools/cmd/goimports@latest
