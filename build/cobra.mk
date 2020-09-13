BIN ?= bin/$(SERVICE)
SERVICE ?= todo

# ---
# Common
# ---

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: install
install: install-cobra ## install

.PHONY: test
test: ## test
	@echo "[cobra]"
	@cobra --help
	@$(SERVICE)/$(BIN) test

# ---
# Cobra: https://github.com/spf13/cobra
# ---

# https://github.com/spf13/cobra/issues/1215#issuecomment-686429510
.PHONY: install-cobra
install-cobra:
	$(eval GOGET ?= go get -u -v)
	@hash cobra > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) github.com/spf13/cobra/cobra@v1.0.0; \
	fi

# ---
# Project
# ---

CONFIG ?= .cobra.yml

.PHONY: init
init: install-cobra ## initialize cobra cli
	$(eval PKG_NAME ?= github.com/ks6088ts/spidey/$(SERVICE))
	mkdir -p $(SERVICE) && \
	cd $(SERVICE) && \
	cobra init \
		--pkg-name $(PKG_NAME) \
		--config ../$(CONFIG)

.PHONY: add
add: install-cobra ## add cobra command
	$(eval CMD ?= hello)
	$(eval PARENT_CMD ?= rootCmd)
	cd $(SERVICE) && \
	cobra add $(CMD) \
		--config ../$(CONFIG) \
		--parent $(PARENT_CMD)

.PHONY: build
build: ## build applications
	$(eval GOBUILD ?= go build)
	$(eval LDFLAGS ?= -ldflags="-s -w")
	cd $(SERVICE) && \
		$(GOBUILD) $(LDFLAGS) -o $(BIN)
