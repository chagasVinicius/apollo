SHELL := /bin/bash

# ==============================================================================
# help

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

# ==============================================================================
# Setup

GOLANGCI_LINT_VERSION=v1.51.1
GOBIN := $(shell go version)
GOPATH := $(shell go env GOPATH)

check.go:
	go version >/dev/null 2>&1 || (echo "ERROR: go is not installed" && exit 1)

check.docker:
	docker version >/dev/null 2>&1 || (echo "ERROR: docker is not installed" && exit 1)
	docker compose version >/dev/null 2>&1 || (echo "ERROR: docker-compose is not installed" && exit 1)

## Install go tools
setup: check.go check.docker
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install mvdan.cc/gofumpt@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin ${GOLANGCI_LINT_VERSION}
	docker pull postgres:15-alpine

run-migrations: docker exec "$(docker ps -aqf "name=apollo-api")" ./admin migrate

# ==============================================================================
# GO

all: lint test vulncheck

## Download GO dependencies
deps:
	go mod tidy
	go mod vendor

## Upgrade GO dependencies
deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

## Format code (w/ gofumpt)
formt:
	gofumpt -l -w .

## Execute code lint
lint:
	golangci-lint run --fix

## Execute code vulnerability check
vulncheck:
	govulncheck ./...

## Execute unit tests
test:
	go test -v -count=1 ./...
