package service

import "github.com/Sutheres/property-service/internal/models"

func (s service) AddProperty(property models.Property) error {

	err := s.db.AddProperty(property)
	if err != nil {
		return err
	}
	return nil
}
