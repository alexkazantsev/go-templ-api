# This included makefile should define the 'custom' target rule which is called here.
# The (-) sign before `include` will do an implicit check if the file exists.
-include $(INCLUDE_MAKEFILE)

include .env
export

TOOLBOX_IMAGE ?= toolbox:11.11.2024
TOOLBOX_RUN ?= @docker run --rm --network="host" -v "$(PWD):/workspace" -w "/workspace" --platform linux/amd64 $(TOOLBOX_IMAGE)
MIGRATION_PATH ?= modules/database/migrations

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

lint: ### Run golangci linter
	@golangci-lint run
.PONY: lint

test: ### Run unit test
	@go test ./... -v
.PONY: test

generate-sqlc: ### Generate sqlc code based on migrations
	$(TOOLBOX_RUN) sqlc generate
.PONY: generate-sqlc

migrate-create: ### Create migration with specified filename
	$(TOOLBOX_RUN) migrate create -ext sql -dir $(MIGRATION_PATH) '$(shell bash -c 'read -p "Enter migration name: " name; echo $$name')'
.PONY: migrate-create

migrate-up: ### Run all migrations
	$(TOOLBOX_RUN) migrate -path $(MIGRATION_PATH) -database '$(DB_DSN)?sslmode=disable' up
.PONY: migrate-up