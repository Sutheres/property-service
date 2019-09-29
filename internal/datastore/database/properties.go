package database

import (
	"github.com/Sutheres/property-service/internal/models"
	"github.com/pkg/errors"
	"time"
)

type propertyQueries interface {
	GetProperties() ([]models.Property, error)
	GetProperty(propertyID string) (models.Property, error)
	AddProperty(property models.Property) error
}

func (d *database) GetProperty(propertyID string) (models.Property, error) {
	p := models.Property{}
	err := d.db.Get(&p, "SELECT * FROM properties WHERE property_id=$1", propertyID)
	return p, errors.Wrap(err, "GetProperty")
}

func (d *database) GetProperties() ([]models.Property, error) {
	var properties []models.Property
	err := d.db.Select(&properties, "SELECT * FROM properties")
	var p []models.Property
	for _, props := range properties {
		p2 := props
		p = append(p, p2)
	}
	return p, errors.Wrap(err, "GetProperties")
}

func (d *database) AddProperty(p models.Property) error {

	//create guid for each property
	if len(p.StreetNumber) < 1 {
		return errors.Wrap(errors.New("Invalid street number"), "Error: ")
	}

	if len(p.Street) < 1 {
		return errors.Wrap(errors.New("Invalid street"), "Error: ")
	}

	if len(p.City) < 1 {
		return errors.Wrap(errors.New("Invalid city"), "Error: ")
	}

	if len(p.State) < 1 {
		return errors.Wrap(errors.New("Invalid state"), "Error: ")
	}

	if len(p.Zip) < 1 {
		return errors.Wrap(errors.New("Invalid zip code"), "Error: ")
	}

	if p.RealEstateType == "Residential" || p.PropertyType == "Apartment building" {
		if float32(p.Bedrooms) < 1 {
			return errors.Wrap(errors.New("Minimum bedrooms: 1"), "Error: ")
		}

		if float32(p.Bathrooms) < .5 {
			return errors.Wrap(errors.New("Minimum bathrooms: .5"), "Error: ")
		}
	}

	if len(p.Price) < 1 {
		return errors.Wrap(errors.New("Invalid price"), "Error: ")
	}

	reTypes := []string{
		"Residential",
		"Commercial",
		"Industrial",
		"Land",
	}

	isValid := false
	for _, t :=range reTypes {
		if p.RealEstateType != t {
			isValid = true
		}
	}
	if !isValid {
		return errors.Wrap(errors.New("Invalid real estate type"), "Error: ")
	}

	propertyTypes := []string{
		"Single-family",
			"Multi-family",
			"Condo",
			"Co-op",
			"Townhouse",
			"Duplex",
			"Triple decker",
			"4-plex",
			"High value home",
			"Generational",
			"Vacation home",
			"Shopping center",
			"Strip mall",
			"Medical building",
			"Education building",
			"Hotel",
			"Office building",
			"Apartment building",
			"Manufacturing building",
			"Warehouse",
			"Other",
			"Vacant land",
			"Working farm",
			"Ranch",
	}

	isValid = false
	for _, pt :=range propertyTypes {
		if p.PropertyType == pt {
			isValid = true
		}
	}
	if !isValid {
		return errors.Wrap(errors.New("Invalid property type"), "Error: ")
	}

	now := time.Now().String()
	p.CreatedAt = now
	p.UpdatedAt = now

	tx := d.db.MustBegin()
	_, err := tx.NamedExec(
		"INSERT INTO properties (property_id, street_number, street, street_suffix, city, state, zip, "+
			"bedrooms, bathrooms, price, real_estate_type, property_type, created_at, updated_at) "+
			"VALUES (:property_id, :street_number, :street, :street_suffix, :city, :state, :zip, "+
			":bedrooms, :bathrooms, :price, :real_estate_type, :property_type, :created_at, :updated_at)", &p)
	if err != nil {
		return errors.Wrap(err, "AddProperty.NamedExec")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "AddProperty.Commit")
	}
	return nil
}
