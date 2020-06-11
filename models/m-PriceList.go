package models

// PriceList - list
type PriceList struct {
	// ID    int    `sql:", pk"`
	NAME  string `sql:", unique" json:"name"`
	PRICE int    `sql:", notnull" json:"price"`
}
