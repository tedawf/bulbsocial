package main

import (
	"log"

	"github.com/tedawf/tradebulb/internal/db"
	"github.com/tedawf/tradebulb/internal/env"
	"github.com/tedawf/tradebulb/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/tradebulb_local?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
