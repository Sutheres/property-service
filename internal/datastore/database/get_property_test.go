package database

import "testing"

func TestDatabase_GetProperty(t *testing.T) {
	//prepareTestDatabase()
	_, err := d.GetProperty(1)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
}
