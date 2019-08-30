package database

import (
	m "github.com/property/internal/models"
)

func (d database) GetProperties() ([]m.Property, error) {

	var properties []m.Property
	err := d.q.Select(&properties, "SELECT * FROM property")
	var p []m.Property
	for _, props := range properties {
		p2 := props
		p = append(p, p2)
	}
	return p, err
}
