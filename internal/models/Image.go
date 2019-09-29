package models

import "github.com/Sutheres/property-service/gen/models"

type Image struct {
	ID    string `db:"id"`
	Title string `db:"title"`
	URL   string `db:"url"`
}

func (i Image) toImage() models.Image {
	return models.Image{
		ID:    i.ID,
		Title: i.Title,
		URL:   i.URL,
	}
}