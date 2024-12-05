.PHONY: postgres
postgres:
	docker run --name bulbsocial -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:17-alpine

.PHONY: createdb
createdb:
	docker exec -it bulbsocial createdb --username=root --owner=root bulb_dev

.PHONY: dropdb
dropdb:
	docker exec -it bulbsocial dropdb bulb_dev

.PHONY: migrateup
migrateup:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose up
	
.PHONY: migrateup1
migrateup1:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose up 1

.PHONY: migratedown
migratedown:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose down

.PHONY: migratedown1
migratedown1:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose down 1

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: server
server:
	go run cmd/main.go

.PHONY: mock
mock:
	mockgen -package=mockdb -destination internal/db/mock/store.go github.com/tedawf/bulbsocial/internal/db Store,Querier
