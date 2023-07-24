.PHONY: help

help: ## You are here! showing all command documenentation.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml

environment: ## Setup environment.
environment:
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d

server: ## Running Application
server:
	go run cmd/main.go

test:  ## Running golang testing.
test:
	go test ./... -count=1 -coverprofile=coverage.out

test-cover:  ## Open golang testing coverage
test-cover:
	go tool cover -html=coverage.out

mock-gen:
	mockgen -source=internal/app/service/movies.go -destination=internal/app/mocks/mock_movies.go -package=mocks