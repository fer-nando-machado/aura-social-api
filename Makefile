PROJECT_NAME = aura-social-api
MODULE_NAME = api

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info commands:)
	$(info -> setup                   installs dependencies)
	$(info -> format                  formats go files)
	$(info -> build                   builds binary)
	$(info -> test                    runs available tests)
	$(info -> run                     starts server locally)
	$(info -> publish                 pushes master to heroku)
	$(info -> logs                	  views heroku logs)

.PHONY: setup
setup:
	go get -d -v -t ./...
	go install -v ./...
	go mod tidy -v

.PHONY: format
format:
	go fmt ./...

.PHONY: build
build:
	go build -v -o $(MODULE_NAME).bin ./$(MODULE_NAME)
	chmod +x $(MODULE_NAME).bin
	echo $(MODULE_NAME).bin

.PHONY: test
test:
	go test ./... -v -covermode=count

.PHONY: run
run:
	go run ./$(MODULE_NAME)

.PHONY: publish
publish:
	git push heroku master

.PHONY: logs
logs:
	heroku logs --tail
