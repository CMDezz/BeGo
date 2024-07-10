postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root BeGo

dropdb:
	docker exec -it postgres16 dropdb BeGo

migrateInit:
	migrate create -ext sql -dir dto/migrations -seq migrate_schema

migrateUp:
	migrate -path dto/migrations -database "postgresql://root:secret@localhost:5432/BeGo?sslmode=disable" -verbose up

migrateDown:
	migrate -path dto/migrations -database "postgresql://root:secret@localhost:5432/BeGo?sslmode=disable" -verbose down

start:
	go run main.go

.PHONY: postgres createdb dropdb migrateInit migrateUp migrateDown start