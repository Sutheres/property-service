package database

import (
	"github.com/Sutheres/property-service/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

)

type database struct {
	db   *sqlx.DB
	url string
}
type Database interface{
	GetProperties() ([]models.Property, error)
	GetProperty(propertyID string) (models.Property, error)
	AddProperty(property models.Property) error
}

func NewDatastore(dbUrl string) (Database, error) {
	db, err := sqlx.Connect(
		"postgres", dbUrl,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Connect")
	}
	return &database{
		db:   db,
		url: dbUrl,
	}, nil
}
