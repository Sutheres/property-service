package service

import "github.com/Sutheres/property-service/internal/models"

func (s service) GetProperties() ([]models.Property, error) {
	p, err := s.db.GetProperties()
	return p, err
}
