package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type database struct {
	q   *sqlx.DB
	url string
}
type Database interface{}

func NewDatastore(dbUrl string) (Database, error) {
	db, err := sqlx.Connect(
		"postgres", dbUrl,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Connect")
	}
	return &database{
		q:   db,
		url: dbUrl,
	}, nil
}
