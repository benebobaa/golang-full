.PHONY: build run test createdb dropdb migratecreate migrateup migratedown sqlc

# build:
# 	go build -o myapp ./cmd/api
#
# run:
# 	go run ./cmd/api

test:
	go test ./...

createdb:
	sudo docker exec -it postgres16 createdb --username=root --owner=root sqlc

dropdb:
	sudo docker exec -it postgres16 dropdb sqlc

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/sqlc?sslmode=disable" -verbose down 1

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/sqlc?sslmode=disable" -verbose up

sqlc:
	sqlc generate
