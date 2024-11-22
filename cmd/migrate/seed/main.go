package main

import (
	"log"

	"github.com/tedawf/bulbsocial/internal/db"
	"github.com/tedawf/bulbsocial/internal/env"
	"github.com/tedawf/bulbsocial/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:admin@localhost/bulbsocial_dev?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
