package models

import "time"

// Orders Order's table
type Orders struct {
	ID     int       `sql:",pk"`
	Date   time.Time `sql:"default:now()"`
	Status string    `sql:",notnull, default:Office" json:"status"`

	ClientInitials        string `sql:",notnull" json:"client_initials"`
	ClientPhone           int    `json:"client_phone"`
	CurrentWorkerInitials string `sql:",notnull" json:"current_worker_initials"`
	CurrentWorkerPhone    int    `json:"current_worker_phone"`

	CostManufacturing int `sql:",notnull" json:"cost_manufacturing"`
	CostPainting      int `sql:",notnull" json:"cost_painting"`
	CostFinishing     int `sql:",notnull" json:"cost_finishing"`
	CostFull          int `sql:",notnull" json:"cost_full"`
}
