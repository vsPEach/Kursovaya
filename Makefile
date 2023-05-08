BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"
COMPOSE_PATH = deployments/docker-compose.yml


GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./config/config.yml

version: build
	$(BIN) version

generate:
	go generate ./...

test:
	go test -race ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

up:
	docker compose -f $(COMPOSE_PATH) up -d --build --force-recreate

down:
	docker compose -f $(COMPOSE_PATH) down

.PHONY: build run build-img version test lint up down
