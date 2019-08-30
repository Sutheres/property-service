package service

import "github.com/property/internal/models"

func (s service) GetProperty(propertyID int64) (models.Property, error) {
	p, err := s.db.GetProperty(propertyID)
	return p, err
}
