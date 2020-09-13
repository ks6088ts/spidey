OS ?= $(shell uname)
GOGET ?= go get -u -v

# ---
# Common
# ---

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: install
install: install-protoc install-protoc-go ## install

.PHONY: test
test: ## test
	@echo "[grpc]"
	@echo "Running OS=$(OS)"
	protoc --version

# ---
# gRPC: https://grpc.io/
# ---

# http://google.github.io/proto-lens/installing-protoc.html
.PHONY: install-protoc
install-protoc: ## install protocol buffers
ifeq ($(OS),Darwin)
	$(eval PROTOC_VERSION ?= 3.13.0)
	$(eval PROTOC_PLATFORM ?= osx-x86_64)
	$(eval PROTOC_ZIP ?= protoc-$(PROTOC_VERSION)-$(PROTOC_PLATFORM).zip)
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)
else ifeq ($(OS),Linux)
	sudo apt-get install -y protobuf-compiler
endif

# https://grpc.io/docs/languages/go/quickstart/
.PHONY: install-protoc-go
install-protoc-go: ## install Go plugin for protocol buffers
	$(GOGET) github.com/golang/protobuf/protoc-gen-go

# ---
# Project
# ---

.PHONY: protoc
protoc:
	$(eval PROTO_DIR ?= todo/protos)
	$(eval PROTO_FILE ?= $(PROTO_DIR)/todo.proto)
	protoc \
		--proto_path $(PROTO_DIR) \
		--go_out=plugins=grpc:. \
		$(PROTO_FILE)
