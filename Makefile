GO_BIN ?= go
OS := $(shell uname -s)

test: deps
	@$(GO_BIN) test recordkeeper/*.go

deps:
	@$(GO_BIN) mod tidy

test-cover: deps
	@$(GO_BIN) test recordkeeper/*.go -cover -coverprofile=c.out
	@$(GO_BIN) tool cover -html=c.out -o coverage.html

# --- Linting
lint:
	@golangci-lint run

install-lint:
ifeq ($(OS), Darwin)
	brew install golangci/tap/golangci-lint
endif
ifeq ($(OS), Linux)
	GO111MODULE=on $(GO_BIN) get github.com/golangci/golangci-lint/cmd/golangci-lint@692dacb773b703162c091c2d8c59f9cd2d6801db
endif
# ---
