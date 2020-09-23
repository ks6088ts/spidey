GOGET ?= go get -u -v

# ---
# Common
# ---

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

# ---
# gqlgen: https://gqlgen.com/
# ---

.PHONY: install
install: install-gqlgen ## install

.PHONY: install-gqlgen
install-gqlgen: ## install gqlgen
	$(GOGET) github.com/99designs/gqlgen
