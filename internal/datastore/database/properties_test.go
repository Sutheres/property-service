package database

import "testing"

func TestDatabase_GetProperty(t *testing.T) {
	prepareTestDatabase()
	_, err := d.GetProperty("property_id")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
}
