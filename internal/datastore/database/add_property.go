package database

import "github.com/Sutheres/property-service/internal/models"

func (d database) AddProperty(property models.Property) error {

	tx := d.db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO property (address, turn_key, note, bedrooms, bathrooms) VALUES (:address, :turn_key, :note, :bedrooms, :bathrooms)", &property)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
