package models

import "time"

// OrdersParam list
type OrdersParam struct {
	Title   string `json:"title"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	Filenka string `json:"filenka"`
}

// OrderStatus and Data list for orders
type OrderStatus struct {
	DataOffice        time.Time `sql:"default:now()" json:"data_office"`
	DataManufacturing time.Time `sql:"default:now()" json:"data_manufacturing"`
	DataGrinding      time.Time `sql:"default:now()" json:"data_grinding"`
	DataPrinting      time.Time `sql:"default:now()" json:"data_printing"`
	DataReady         time.Time `sql:"default:now()" json:"data_ready"`

	StatusOffice        bool `sql:",notnull, default:true" json:"status_office"`
	StatusManufacturing bool `sql:",notnull, default:false" json:"status_manufacturing"`
	StatusGrinding      bool `sql:",notnull, default:false" json:"status_grinding"`
	StatusPrinting      bool `sql:",notnull, default:false" json:"status_printing"`
	StatusReady         bool `sql:",notnull, default:false" json:"status_ready"`
}

// Orders Order's table
type Orders struct {
	ID     int         `sql:",pk" json:"id"`
	Title  string      `json:"title"`
	Date   time.Time   `sql:"default:now()"`
	Status OrderStatus `json:"status"`

	ClientID       int    `sql:"default:Null" json:"client_id"`
	ClientInitials string `sql:",notnull" json:"client_initials"`
	ClientPhone    int    `json:"client_phone"`

	CurrentWorkerID       int    `sql:",notnull" json:"current_worker_id"`
	CurrentWorkerInitials string `sql:",notnull" json:"current_worker_initials"`
	CurrentWorkerPhone    int    `json:"current_worker_phone"`

	Color        string `json:"color"`
	Patina       string `json:"patina"`
	FasadArticle string `json:"fasad_article"`
	Material     string `json:"material"`

	CostManufacturing int `sql:",notnull" json:"cost_manufacturing"`
	CostPainting      int `sql:",notnull" json:"cost_painting"`
	CostFinishing     int `sql:",notnull" json:"cost_finishing"`
	CostFull          int `sql:",notnull" json:"cost_full"`

	Params []OrdersParam `json:"params"`
}

// SavedOrders - saves orders when status chenged
type SavedOrders struct {
	ID      int `sql:",pk"`
	OrderID int
	Title   string      `json:"title"`
	Date    time.Time   `sql:"default:now()"`
	Status  OrderStatus `json:"status"`

	ClientID       int    `sql:"default:Null"`
	ClientInitials string `sql:",notnull" json:"client_initials"`
	ClientPhone    int    `json:"client_phone"`

	CurrentWorkerID       int    `sql:",notnull"`
	CurrentWorkerInitials string `sql:",notnull" json:"current_worker_initials"`
	CurrentWorkerPhone    int    `json:"current_worker_phone"`

	Color        string `json:"color"`
	Patina       string `json:"patina"`
	FasadArticle string `json:"fasad_article"`
	Material     string `json:"material"`

	CostManufacturing int `sql:",notnull" json:"cost_manufacturing"`
	CostPainting      int `sql:",notnull" json:"cost_painting"`
	CostFinishing     int `sql:",notnull" json:"cost_finishing"`
	CostFull          int `sql:",notnull" json:"cost_full"`

	Params []OrdersParam `json:"params"`
}
