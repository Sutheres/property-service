package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/testfixtures.v2"
	"log"
	"os"
	"testing"
	"time"
)

var (
	d        database
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	var err error
	db, err := sql.Open(
		"postgres",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "sql.Open"))
	}
	//

	fixtures, err = testfixtures.NewFolder(
		db, &testfixtures.PostgreSQL{},
		"./testdata/fixtures",
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "testfixtures.NewFolder"))
	}
	//
	dbConn := sqlx.NewDb(db, "postgres")
	dbConn.SetMaxIdleConns(5)
	dbConn.SetMaxOpenConns(5)
	dbConn.SetConnMaxLifetime(15 * time.Minute)
	//
	d = database{
		db: dbConn,
	}
	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
