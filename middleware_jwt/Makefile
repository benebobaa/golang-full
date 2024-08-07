dsn:
	export DB_DRIVER="postgres"
	export DB_DSN="postgresql://root:root@localhost:5432/jwt?sslmode=disable"

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root jwt

dropdb:
	docker exec -it postgres16 dropdb -U root jwt

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/jwt?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/jwt?sslmode=disable" -verbose up

sqlc:
	sqlc generate

.PHONY: build run test createdb dropdb migratecreate migrateup migratedown sqlc dsn

