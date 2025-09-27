GOLANGCILINT_VERSION := v2.5.0
GO_PATH := $(shell go env GOPATH)/bin
GOLANGCILINT_BIN := $(GO_PATH)/golangci-lint

.PHONY: lint fmt test

all: lint fmt test

$(GOLANGCILINT_BIN):
	@if ! test -x $(GOLANGCILINT_BIN) || ! $(GOLANGCILINT_BIN) --version | grep -q $(GOLANGCILINT_VERSION); then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_PATH) $(GOLANGCILINT_VERSION); \
	fi

fmt: $(GOLANGCILINT_BIN)
	$(GOLANGCILINT_BIN) fmt ./... -v

lint: $(GOLANGCILINT_BIN)
	@$(GOLANGCILINT_BIN) run $(LINT_FLAGS) ./... --fix -v

test:
	go test -v ./... -short
