package service

import "github.com/property/internal/models"

func (s service) GetProperties() ([]models.Property, error) {
	p, err := s.db.GetProperties()
	return p, err
}
