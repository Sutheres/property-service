package database

import (
	"github.com/property/internal/models"
)

func (d database) GetProperties() ([]models.Property, error) {
	var properties []models.Property
	err := d.q.Select(&properties, "SELECT * FROM property")
	var p []models.Property
	for _, props := range properties {
		p2 := props
		p = append(p, p2)
	}
	return p, err
}
