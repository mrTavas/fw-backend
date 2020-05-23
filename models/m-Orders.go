package models

import "time"

// OrdersParam list
type OrdersParam struct {
	Title   string `json:"title"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	Filenka string `json:"filenka"`
}

// Orders Order's table
type Orders struct {
	ID     int       `sql:",pk"`
	Date   time.Time `sql:"default:now()"`
	Status string    `sql:",notnull" json:"status"`

	ClientInitials        string `sql:",notnull" json:"client_initials"`
	ClientPhone           int    `json:"client_phone"`
	CurrentWorkerInitials string `sql:",notnull" json:"current_worker_initials"`
	CurrentWorkerPhone    int    `json:"current_worker_phone"`

	CostManufacturing int `sql:",notnull" json:"cost_manufacturing"`
	CostPainting      int `sql:",notnull" json:"cost_painting"`
	CostFinishing     int `sql:",notnull" json:"cost_finishing"`
	CostFull          int `sql:",notnull" json:"cost_full"`

	Params []OrdersParam `json:"params"`
}
