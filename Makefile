GO_BIN ?= go

test:
	@$(GO_BIN) test recordkeeper/*.go

test-cover:
	@$(GO_BIN) test recordkeeper/*.go -cover -coverprofile=c.out
	@$(GO_BIN) tool cover -html=c.out -o coverage.html
