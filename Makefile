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
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose down

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	go test -v -cover ./...
