package models

// Managers Table of managers
type Managers struct {
	ID       int    `sql:", pk"`
	UUID     string `sql:", unique" json:"uuid"`
	Phone    int    `sql:", unique, notnull" json:"phone"`
	Password string `sql:",notnull" json:"pass"`
	Initials string `sql:",notnull" json:"initials"`
}
