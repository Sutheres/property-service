package service

import "github.com/property/internal/models"

func (s service) AddProperty(property models.Property) error {

	err := s.db.AddProperty(property)
	if err != nil {
		return err
	}
	return nil
}
