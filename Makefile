DB_CONN = "postgresql://root:admin123@localhost:5432/simple_bank?sslmode=disable"
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -d postgres
pgadmin: 
	docker run -p 443:443 -e 'PGADMIN_DEFAULT_EMAIL=ariel.serato@megatrans.com.ar' -e 'PGADMIN_DEFAULT_PASSWORD=admin123' -d dpage/pgadmin4
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropdb: 
	docker exec -it postgres dropdb --username=root simple_bank
migrateup:
	migrate -path db/migration -database $(DB_CONN) -verbose up
migrateup1:
	migrate -path db/migration -database $(DB_CONN) -verbose up 1	
migratedown:
	migrate -path db/migration -database $(DB_CONN) -verbose down
migratedown1:
	migrate -path db/migration -database $(DB_CONN) -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Arielcito/simple-bank-go/db/sqlc Store 
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock pgadmin migratedown1 migrateup1