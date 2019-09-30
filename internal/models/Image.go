package models

import (
	"github.com/Sutheres/property-service/gen/models"
	"time"
)

type Image struct {
	ID    string `db:"id"`
	Title string `db:"title"`
	URL   string `db:"url"`
	PropertyID string `db:"property_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (i Image) toImage() models.Image {
	return models.Image{
		ID:    i.ID,
		Title: i.Title,
		URL:   i.URL,
	}
}