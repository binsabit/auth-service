run:
	go run main.go -config=config.yaml
migration:
	@read -p "Enter migration name: " name;\
		migrate create -ext sql -dir migrations/postgres -seq $$name

migrate-up:
	migrate -path migrations/postgres \
		-database "postgresql://yerda:postgres@localhost:5432/dbauth?sslmode=disable" -verbose up
migrate-down:
	migrate -path migrations/postgres \
		-database "postgresql://yerda:postgres@localhost:5432/dbauth?sslmode=disable" -verbose down
migrate-force:
	@read -p "Enter version to force :" version;\
		migrate -path migrations/postgres \
			-database "postgresql://yerda:postgres@localhost:5432/dbauth?sslmode=disable" -verbose force $$version