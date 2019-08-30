package models

type Property struct {
	Address   string `db:"address"`
	TurnKey   bool   `db:"turn_key"`
	Note      string `db:"note"`
	Bedrooms  int64  `db:"bedrooms"`
	Bathrooms int64  `db:"bathrooms"`
}
