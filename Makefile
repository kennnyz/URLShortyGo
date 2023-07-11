migrate-up:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/ozon?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/ozon?sslmode=disable" -verbose down

createdb:
	docker run --name ozondb -e POSTGRES_DB=ozon -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

run-up:
	docker-compose up -d

.PHONY: migrate-down migrate-up run-up
