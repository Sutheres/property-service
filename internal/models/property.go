package models

import "github.com/Sutheres/property-service/gen/models"

type Property struct {
	ID        int64 `db:"id"`
	Address   string `db:"address"`
	TurnKey   bool   `db:"turn_key"`
	Note      string `db:"note"`
	Bedrooms  int64  `db:"bedrooms"`
	Bathrooms int64  `db:"bathrooms"`
}

func (p Property) ToProperty() models.Property {
	return models.Property{
		Address:   &p.Address,
		Bathrooms: p.Bathrooms,
		Bedrooms:  p.Bedrooms,
		ID:        p.ID,
		Note:      p.Note,
		TurnKey:   p.TurnKey,
	}
}
