package models

// Clients Table of clients
type Clients struct {
	ID       int    `sql:", pk"`
	Phone    int    `sql:", unique, notnull" json:"phone"`
	Password int    `sql:",notnull"`
	Initials string `sql:",notnull" json:"initials"`
	Score    int    `sql:"default:1"`
}
