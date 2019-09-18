package database

import "github.com/Sutheres/property-service/internal/models"

func (d database) GetProperty(propertyID string) (models.Property, error)  {
	p := models.Property{}
	err := d.db.Get(&p, "SELECT * FROM properties WHERE property_id=$1", propertyID)
	return p, err
}