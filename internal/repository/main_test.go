package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bulb_dev?sslmode=disable"
)

var testQueries *sqlc.Queries
var TestDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	TestDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = sqlc.New(TestDB)

	os.Exit(m.Run())
}
