package models

import (
	"github.com/Sutheres/property-service/gen/models"
)

type Property struct {
	ID             int64    `db:"id"`
	PropertyID     string   `db:"property_id"`
	StreetNumber   string   `db:"street_number"`
	Street         string   `db:"street"`
	StreetSuffix   string   `db:"street_suffix"`
	City           string   `db:"city"`
	State          string   `db:"state"`
	Zip            string   `db:"zip"`
	Bedrooms       float32  `db:"bedrooms"`
	Bathrooms      float32  `db:"bathrooms"`
	Price          string   `db:"price"`
	SquareFeet     float64  `db:"square_feet"`
	Description    string   `db:"description"`
	Status         string   `db:"status"`
	Note           string   `db:"note"`
	Images         []*Image `db:"images"`
	RealEstateType string   `db:"real_estate_type"`
	PropertyType   string   `db:"property_type"`
	CreatedAt      string   `db:"created_at"`
	UpdatedAt      string   `db:"updated_at"`
}

func (p Property) ToProperty() models.Property {

	var images []*models.Image

	for _, i := range p.Images {
		p := i.toImage()
		images = append(images, &p)
	}

	return models.Property{
		Bathrooms:      p.Bathrooms,
		Bedrooms:       p.Bedrooms,
		City:           p.City,
		Price:          p.Price,
		PropertyID:     p.PropertyID,
		PropertyType:   p.PropertyType,
		RealEstateType: p.RealEstateType,
		State:          p.State,
		Street:         p.Street,
		StreetNumber:   p.StreetNumber,
		StreetSuffix:   p.StreetSuffix,
		Zip:            p.Zip,
		CreatedAt:      p.CreatedAt,
		SquareFeet:     p.SquareFeet,
		Description:    p.Description,
		Note:           p.Note,
		Images:         images,
	}
}
