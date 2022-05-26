postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRESPASSWORD=admin123 -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropdb: 
	docker exec -it postgres dropdb --username=root --owner=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc

