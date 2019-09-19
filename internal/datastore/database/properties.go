package database

import (
	"github.com/Sutheres/property-service/internal/models"
	"github.com/pkg/errors"
)

type propertyQueries interface {
	GetProperties() ([]models.Property, error)
	GetProperty(propertyID string) (models.Property, error)
	AddProperty(property models.Property) error
}

func (d *database) GetProperty(propertyID string) (models.Property, error)  {
	p := models.Property{}
	err := d.db.Get(&p, "SELECT * FROM properties WHERE property_id=$1", propertyID)
	return p, errors.Wrap(err, "GetProperty")
}

func (d *database) GetProperties() ([]models.Property, error) {
	var properties []models.Property
	err := d.db.Select(&properties, "SELECT * FROM properties")
	var p []models.Property
	for _, props := range properties {
		p2 := props
		p = append(p, p2)
	}
	return p, errors.Wrap(err, "GetProperties")
}

func (d *database) AddProperty(property models.Property) error {

	tx := d.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO property (address, turn_key, note, bedrooms, bathrooms) VALUES (:address, :turn_key, :note, :bedrooms, :bathrooms)", &property)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
