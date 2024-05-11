migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "postgres://sportgroup_api_user:${db_password}@localhost:${db_port}/sportgroup_auth?sslmode=disable" -verbose up

wire:
	@cd internal/bootstrap && wire
