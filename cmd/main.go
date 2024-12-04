package main

import (
	"database/sql"
	"log"

	"github.com/tedawf/bulbsocial/internal/api"
	"github.com/tedawf/bulbsocial/internal/db"
	"go.uber.org/zap"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// db
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		sugar.Fatal("cannot connect to db: ", err)
	}
	defer conn.Close()
	store := db.NewStore(conn)

	// server
	server := api.NewServer(store, sugar)
	if err = server.Start(serverAddress); err != nil {
		sugar.Fatal("cannot start server: ", err)
	}
}
