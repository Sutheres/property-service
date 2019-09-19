package database

import (
	"github.com/Sutheres/property-service/internal/models"
	"testing"
)

type addPropertyTest struct {
	input models.Property
	expected bool
}

func TestDatabase_GetProperty(t *testing.T) {
	prepareTestDatabase()
	_, err := d.GetProperty("property_id")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
}

func TestDatabase_GetProperties(t *testing.T) {
	prepareTestDatabase()
	_, err := d.GetProperties()
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
}

func TestDatabase_AddProperty(t *testing.T) {
	p1 := models.Property{
		PropertyID:     "property1",
		StreetNumber:   "2933",
		Street:         "Montana",
		StreetSuffix:   "Ave",
		City:           "Cincinnati",
		State:          "OH",
		Zip:            "45211",
		Bedrooms:       4.5,
		Bathrooms:      2.5,
		Price:          "40,000",
		RealEstateType: "Residential",
		PropertyType:   "Single-family",
	}
	p2 := models.Property{
		PropertyID:     "property2",
		StreetNumber:   "2933",
		Street:         "Montana",
		StreetSuffix:   "Ave",
		City:           "Cincinnati",
		State:          "OH",
		Zip:            "45211",
		Bedrooms:       4.5,
		Bathrooms:      2.5,
		Price:          "40,000",
		RealEstateType: "dfkajfhad",
		PropertyType:   "Single-family",
	}
	p3 := models.Property{
		PropertyID:     "property3",
		StreetNumber:   "2933",
		Street:         "Montana",
		StreetSuffix:   "Ave",
		City:           "Cincinnati",
		State:          "OH",
		Zip:            "45211",
		Bedrooms:       0,
		Bathrooms:      2.5,
		Price:          "40,000",
		RealEstateType: "Residential",
		PropertyType:   "Single-family",
	}
	p4 := models.Property{
		PropertyID:     "property4",
		StreetNumber:   "2933",
		Street:         "Montana",
		StreetSuffix:   "Ave",
		City:           "Cincinnati",
		State:          "OH",
		Zip:            "45211",
		Bedrooms:       4.5,
		Bathrooms:      2.5,
		Price:          "40,000",
		RealEstateType: "Residential",
		PropertyType:   "Single-dfadsf",
	}

	properties := []addPropertyTest{
		{p1, true},
		{p2, false},
		{p3, false},
		{p4, false},
	}

	for _, p :=range properties {
		err := d.AddProperty(p.input)
		isValid := true
		if err != nil {
			isValid = false
		}
		if isValid != p.expected {
			t.Errorf("AddProperty(%s): expected (%t), actual (%t) ", p.input.PropertyID, p.expected, isValid)
		}
	}
}
