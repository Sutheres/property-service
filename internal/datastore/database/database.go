package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/property/internal/models"
)

type database struct {
	q   *sqlx.DB
	url string
}
type Database interface{
	GetProperties() ([]models.Property, error)
}

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
