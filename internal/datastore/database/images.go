package database

import (
	"github.com/Sutheres/property-service/internal/models"
	"github.com/pkg/errors"
	"time"
)

type imageQueries interface {
	InsertImage(image models.Image) error
	GetImages() ([]models.Image, error)
}

func (d *database) InsertImage(i models.Image) error {

	now := time.Now()
	i.CreatedAt = now
	i.UpdatedAt = now

	tx := d.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO images (title, url, created_at, updated_at, property_id) " +
		"VALUES (:title, :url, :created_at, :updated_at, :property_id)", &i)

	if err != nil {
		return errors.Wrap(err, "NamedExec")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Commit")
	}
	return nil
}

func (d *database) GetImages() ([]models.Image, error) {
	return nil, nil
}
