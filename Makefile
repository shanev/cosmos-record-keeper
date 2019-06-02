GO_BIN ?= go

test:
	@$(GO_BIN) test recordkeeper/*.go
