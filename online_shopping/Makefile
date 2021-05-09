APPLICATION_NAME := $(shell grep "name\": " package.json | sed -E 's/.*:\s*"(.+)",$$/\1/')
BASE_VERSION := $(shell grep "version\": " package.json | sed -E 's/.*:\s*"(.+)",$$/\1/')

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
	yarn

docker-build: ## Build docker image
	test "${shell docker images -q ${APPLICATION_NAME}:${BASE_VERSION}}" != "" || docker build -t ${APPLICATION_NAME}:${BASE_VERSION} ./
