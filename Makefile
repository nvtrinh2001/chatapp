postgresstart:
	docker start postgres15

postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

migratecreate: 
	migrate create -ext sql -dir db/migrations add_users_table

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@127.0.0.1:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@127.0.0.1:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown
