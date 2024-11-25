include .envrc
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: new_migration
new_migration:
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filter-out $@, $(MAKECMDGOALS))

.PHONY: migrate_up
migrate_up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) -verbose up $(filter-out $@, $(MAKECMDGOALS))

.PHONY: migrate_down
migrate_down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) -verbose down $(filter-out $@, $(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/seed/main.go

.PHONY: test
test:
	@go test -v ./...
