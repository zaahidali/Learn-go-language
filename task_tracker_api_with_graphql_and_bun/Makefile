.PHONY: migrateup migratedown run create-migration

DATABASE_URL := "postgres://postgres:postgres@localhost:5432/task_tracker_db_with_gql_bun?sslmode=disable"

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

%:
	@:

remove-migration: # CAUTION: Use this for migrations which you have not applied
ifeq ($(filter-out $@,$(MAKECMDGOALS)),)
	@echo "Error: 'migration' parameter is required. Usage: make remove-migration <migration>"
else
	rm -f migrations/$(filter-out $@,$(MAKECMDGOALS)).up.sql
	rm -f migrations/$(filter-out $@,$(MAKECMDGOALS)).down.sql
endif

%:
	@: