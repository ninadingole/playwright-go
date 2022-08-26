
GO ?= go

test:
	@echo "Testing..."
	@$(GO) test -race -shuffle=on ./...


