postgres:
	docker run --name bulbsocial -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:17-alpine

createdb:
	docker exec -it bulbsocial createdb --username=root --owner=root bulb_dev

dropdb:
	docker exec -it bulbsocial dropdb bulb_dev

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
