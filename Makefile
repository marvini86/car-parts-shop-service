include .env

create_database:
	docker exec -it postgres bash -c "psql -U postgres -c \"CREATE DATABASE ${PG_DB};\""

drop_database:
	docker exec -it postgres bash -c "psql -U postgres -c \"DROP DATABASE ${PG_DB};\""

create_migration:
	migrate create -ext sql -dir internal/db/migrations -seq init

migrate_up:
	migrate -database "postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=disable" -path internal/db/migrations -verbose up

migrate_down:
	migrate -database "postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=disable" -path internal/db/migrations -verbose down

build:
	go build -o api cmd/api/main.go

run:
	air