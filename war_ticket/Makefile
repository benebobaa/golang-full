.PHONY: build run test createdb dropdb migratecreate migrateup migratedown sqlc dsn

# build:
# 	go build -o myapp ./cmd/api
#
# run:
# 	go run ./cmd/api

test:
	go test ./...

createdb:
	sudo docker exec -it postgres16 createdb --username=root --owner=root warticket

dropdb:
	sudo docker exec -it postgres16 dropdb -U root warticket

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/warticket?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/warticket?sslmode=disable" -verbose up



sqlc:
	sqlc generate
