# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /usr/bin/env bash

export WORKDIR=$(shell pwd)
export APP_NAME=learning

.PHONY: server
# run server
server:
	go run learning/cmd/server

.PHONY: client
# run client
client:
	go run learning/cmd/client

.PHONY: generate
# run go generate
generate:
	go generate ./...

# show help
help:
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\w0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help