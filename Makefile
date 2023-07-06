migrate-up:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/ozon?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/ozon?sslmode=disable" -verbose down

run-up:
	docker-compose up -d
	migrate-up

.PHONY: migrate-down migrate-up run-up
