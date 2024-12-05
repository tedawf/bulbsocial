package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tedawf/bulbsocial/internal/api"
	"github.com/tedawf/bulbsocial/internal/config"
	"github.com/tedawf/bulbsocial/internal/db"
	"go.uber.org/zap"
)

func main() {
	// config
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	// logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to initialize logger: ", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// db
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		sugar.Fatal("cannot connect to db: ", err)
	}
	defer conn.Close()
	store := db.NewSQLStore(conn)

	// server
	server := api.NewServer(store, sugar)
	if err = server.Start(cfg.ServerAddress); err != nil {
		sugar.Fatal("cannot start server: ", err)
	}
}
