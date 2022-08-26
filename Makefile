
GO ?= go


.PHONY: test
test: install-deps
	@echo "Testing..."
	@$(GO) test -race -shuffle=on ./...

install-deps:
	@echo "Downloading browser deps"
	@$(GO) run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps


