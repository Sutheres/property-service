package database

import "github.com/property/internal/models"

func (d database) GetProperty(propertyID int64) (models.Property, error)  {
	p := models.Property{}
	err := d.db.Get(&p, "SELECT * FROM property WHERE id=$1", propertyID)
	return p, err
}