help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env: ## Copy sample files 
	bash scripts/setup.sh

start: ## Up the docker-compose without cache or orphans
	bash scripts/start.sh

stop: ## Down the docker-compose 
	docker-compose down

logs: ## Display logs of your containers 
	docker-compose logs --follow

lint: ## Lint go files
	gofmt -e -l -s -w .

build: ## build the go app
	go build cmd/sample-api/main.go

init:
	make setup-env
	make init-web
	make start 
	make logs

init-web:
	cd web ; npm i ; npm run build

wiki:
	cp -r ./.wiki ../flamingops.wiki

.PHONY: help

