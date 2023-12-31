postgres:
	docker run --name postgres-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=bank -d postgres:latest
createdb:
	docker exec -it postgres-bank createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres-bank dropdb simplae_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrateup1:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -destination db/mock/store.go github.com/hitesh25kumar/db/sqlc Store
.PHONY:  postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock