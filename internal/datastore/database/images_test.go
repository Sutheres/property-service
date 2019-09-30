package database

import (
	"github.com/Sutheres/property-service/internal/models"
	"testing"
)

type addImageTest struct {
	input models.Image
	expected bool
}

func TestDatabase_InsertImage(t *testing.T) {

	tests := []addImageTest{
		{
			models.Image {
				ID:         "1",
				Title:      "fake title",
				URL:        "https://isvr.acceleragent.com/usr/1549750369/CustomPages/images/3.jpg",
				PropertyID: "1",
			}, true,
		},
	}

	for _, i :=range tests {
		err := d.InsertImage(i.input)
		isValid := true
		if err != nil {
			isValid = false
		}
		if isValid != i.expected {
			t.Errorf("AddImage(%s): expected (%t), actual (%t) " + err.Error(), i.input.PropertyID, i.expected, isValid)
		}

	}

}
