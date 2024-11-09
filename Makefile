# This included makefile should define the 'custom' target rule which is called here.
# The (-) sign before `include` will do an implicit check if the file exists.
-include $(INCLUDE_MAKEFILE)

TOOLBOX_IMAGE ?= toolbox:11.02.2024
TOOLBOX_RUN ?= @docker run --rm -v "$(PWD):/workspace" -w "/workspace" --platform linux/amd64 $(TOOLBOX_IMAGE)
MIGRATION_NAME ?= $(shell bash -c 'read -p "Enter migration name: " name; echo $$name')

lint:
	@golangci-lint run

test:
	@go test ./... -v

generate-migrations:
	@go run ./main.go migration:create --name $(MIGRATION_NAME)

generate-sqlc:
	$(TOOLBOX_RUN) sqlc generate