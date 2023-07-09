.PHONY: migrateup migratedown run create-migration gqlgen

DATABASE_URL := "postgres://postgres:postgres@localhost:5432/book_management_api?sslmode=disable"

run:
	go run .

migrate-up:
	migrate -database $(DATABASE_URL) -path migrations up

migrate-down:
	migrate -database $(DATABASE_URL) -path migrations down

create-migration:
ifeq ($(filter-out $@,$(MAKECMDGOALS)),)
	@echo "Error: 'name' parameter is required. Usage: make create-migration <name>"
else
	migrate create -ext sql -dir migrations -seq $(filter-out $@,$(MAKECMDGOALS))
endif

gqlgen:
	go run github.com/99designs/gqlgen generate

%:
	@:
