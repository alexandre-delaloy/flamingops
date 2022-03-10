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

build: ## Build the go app
	go build cmd/sample-api/main.go

init-web: ## Init the web app 
	cd web ; npm i ; npm run build

init: ## Init the app
	make setup-env
	make init-web
	make start 
	make logs


wiki: ## Sync the wiki
	cp -r docs/wiki/ ../flamingops.wiki/
	cd ../flamingops.wiki
	git add . 
	git commit -m "wiki"
	git push

.PHONY: help
