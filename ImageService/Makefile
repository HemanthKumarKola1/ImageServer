postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root1 -e POSTGRES_PASSWORD=secret1 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root1 --owner=root1 schedules

dropdb:
	docker exec -it postgres12 dropdb schedules

migrateup:
	migrate -path db/migration -database "postgresql://root1:secret1@localhost:5432/schedules?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root1:secret1@localhost:5432/schedules?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc