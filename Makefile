# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

default: help

build-inventory-image: ## Build inventory image
	cd inventory && make docker-build

build-discount-image: ## Build discount image
	cd discount && make docker-build

build-online-shopping-image: ## Build online shopping image
	cd online_shopping && make docker-build

build-images: build-inventory-image build-discount-image build-online-shopping-image ## Build all images

test-inventory: ## Run inventory tests
	cd inventory && make test

test-discount: ## Run discount tests
	cd discount && make test

test-integration: build-images ## Run integration tests
	cd ewallet-integration-test && make test

run: ## Run application
	docker-compose up -d --quiet-pull
	@echo "\nReady to accept requests. Try the following command on your terminal:\n"
	@echo ">: curl -H 'Api-Key: c5b6e72c-5b04-4bd2-ba5e-c85a253191dc' http://127.0.0.1:50054/products?userId=11054c65-89dd-46a6-86ab-c603195100a5"
	@echo "\nVisit http://127.0.0.1:5601/app/apm for APM visualization\n"

clean: ## Stop all containers
	cd inventory && docker-compose down --remove-orphans
	cd discount && docker-compose down --remove-orphans
	cd ewallet-integration-test && docker-compose down --remove-orphans
	docker-compose down --remove-orphans