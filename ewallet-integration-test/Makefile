# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

default: help

get-deps: ## Install projects dependencies with Go Module
	@echo "Getting dependencies"
	go mod vendor

test: get-deps ## Run project tests
	@echo "Running tests"
	docker-compose up -d --quiet-pull
	mkdir -p ./.test/cover
	go test -race -coverpkg= ./... -coverprofile=./.test/cover/cover.out -v