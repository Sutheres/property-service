package service

import "github.com/Sutheres/property-service/internal/models"

func (s service) GetProperty(propertyID string) (models.Property, error) {
	p, err := s.db.GetProperty(propertyID)
	return p, err
}
