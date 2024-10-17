package main

import (
	"log"

	"github.com/tedawf/tt4d/internal/db"
	"github.com/tedawf/tt4d/internal/env"
	"github.com/tedawf/tt4d/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/tt4d_local?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
