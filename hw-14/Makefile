

migration_create:
	migrate create -ext sql -dir internal/data/migrations -seq tables

migration_up:
	migrate -path internal/data/migrations -database "mysql://root:root@tcp(localhost:33306)/tours" up