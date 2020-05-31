package models

import "time"

// OrdersParam list
type OrdersParam struct {
	Title   string `json:"title"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	Filenka string `json:"filenka"`

	Color        string `json:"color"`
	Patina       string `json:"patina"`
	FasadArticle string `json:"fasad_article"`
	Material     string `json:"material"`
}

// OrderStatus and Data list for orders
type OrderStatus struct {
	DataOffice        time.Time `sql:"default:now()" json:"data_office"`
	DataManufacturing time.Time `sql:"default:now()" json:"data_manufacturing"`
	DataGrinding      time.Time `sql:"default:now()" json:"data_grinding  "`
	DataReady         time.Time `sql:"default:now()" json:"data_ready"`

	StatusOffice        bool `sql:",notnull, default:false" json:"status_office"`
	StatusManufacturing bool `sql:",notnull, default:false" json:"status_manufacturing"`
	StatusGrinding      bool `sql:",notnull, default:false" json:"status_grinding"`
	StatusReady         bool `sql:",notnull, default:false" json:"status_ready"`
}

// Orders Order's table
type Orders struct {
	ID     int         `sql:",pk"`
	Date   time.Time   `sql:"default:now()"`
	Status OrderStatus `json:"status"`

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
