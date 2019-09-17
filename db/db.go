package db

import (
	"github.com/amacneil/dbmate/pkg/dbmate"
	"github.com/pkg/errors"
	"net/url"
)

type db struct {
	m *dbmate.DB
}

func NewDb() (*db, error) {
	dbURL, err := url.Parse("db_url")

	if err != nil {
		return nil, errors.Wrap(err, "url.Parse")
	}
	return &db{
		m: dbmate.New(dbURL),
	}, nil
}

func (d *db) Migrate() error {
	return errors.Wrap(d.m.CreateAndMigrate(), "CreateAndMigrate")
}