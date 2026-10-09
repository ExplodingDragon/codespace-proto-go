GO ?= go
GO_BIN ?= $(shell bin="$$($(GO) env GOBIN)"; if test -n "$$bin"; then printf "%s" "$$bin"; else printf "%s/bin" "$$($(GO) env GOPATH)"; fi)
BUF ?= $(GO_BIN)/buf

BUF_VERSION ?= v1.8.0
PROTOC_GEN_GO_VERSION ?= v1.36.11
PROTOC_GEN_CONNECT_GO_VERSION ?= v1.20.0

.PHONY: build
build: lint generate test

.PHONY: install
install:
	$(GO) install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	$(GO) install connectrpc.com/connect/cmd/protoc-gen-connect-go@$(PROTOC_GEN_CONNECT_GO_VERSION)

.PHONY: lint
lint:
	$(BUF) lint proto
	$(BUF) format proto --diff --exit-code

.PHONY: format
format:
	$(BUF) format proto -w

.PHONY: generate
generate:
	PATH="$(GO_BIN):$$PATH" $(BUF) generate proto

.PHONY: test
test:
	$(GO) test ./...

.PHONY: clean
clean:
	rm -f codespace/v1/*.pb.go
	rm -rf codespace/v1/codespacev1connect
