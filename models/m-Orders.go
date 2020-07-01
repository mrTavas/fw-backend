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
	DataOfficeStart time.Time `sql:"default:now()" json:"data_office_start"`
	DataOfficeEnd   time.Time `sql:"default:now()" json:"data_office_end"`

	DataManufacturingStart time.Time `sql:"default:now()" json:"data_manufacturing_start"`
	DataManufacturingEnd   time.Time `sql:"default:now()" json:"data_manufacturing_end"`

	DataGrindingStart time.Time `sql:"default:now()" json:"data_grinding_start"`
	DataGrindingEnd   time.Time `sql:"default:now()" json:"data_grinding_end"`

	DataPrintingStart time.Time `sql:"default:now()" json:"data_printing_start"`
	DataPrintingEnd   time.Time `sql:"default:now()" json:"data_printing_end"`

	DataCollectingStart time.Time `sql:"default:now()" json:"data_collecting_start"`
	DataCollectingEnd   time.Time `sql:"default:now()" json:"data_collecting_end"`

	DataReady time.Time `sql:"default:now()" json:"data_ready"`

	StatusOfficeStart bool `sql:",notnull, default:true" json:"status_office_start"`
	StatusOfficeEnd   bool `sql:",notnull, default:false" json:"status_office_end"`

	StatusManufacturingStart bool `sql:",notnull, default:false" json:"status_manufacturing_start"`
	StatusManufacturingEnd   bool `sql:",notnull, default:false" json:"status_manufacturing_end"`

	StatusGrindingStart bool `sql:",notnull, default:false" json:"status_grinding_start"`
	StatusGrindingEnd   bool `sql:",notnull, default:false" json:"status_grinding_end"`

	StatusPrintingStart bool `sql:",notnull, default:false" json:"status_printing_start"`
	StatusPrintingEnd   bool `sql:",notnull, default:false" json:"status_printing_end"`

	StatusCollectingStart bool `sql:",notnull, default:false" json:"status_collecting_start"`
	StatusCollectingEnd   bool `sql:",notnull, default:false" json:"status_collecting_end"`

	StatusReady bool `sql:",notnull, default:false" json:"status_ready"`
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

	CurrentWorkerID       int    `sql:"default:Null" json:"current_worker_id"`
	CurrentWorkerInitials string `sql:"default:Null" json:"current_worker_initials"`
	CurrentWorkerPhone    int    `json:"current_worker_phone"`

	Color        string `json:"color"`
	Patina       string `json:"patina"`
	FasadArticle string `json:"fasad_article"`
	Material     string `json:"material"`

	CostCarpenter int `json:"cost_carpenter"`
	CostGrinder   int `json:"cost_grinder"`
	CostPainter   int `json:"cost_painter"`
	CostCollector int `json:"cost_collector"`

	Params []OrdersParam `json:"params"`
}

// SavedOrders - saves orders when status chenged
type SavedOrders struct {
	ID      int         `sql:",pk"`
	OrderID int         `json:"order_id"`
	Title   string      `json:"title"`
	Date    time.Time   `sql:"default:now()"`
	Status  OrderStatus `json:"status"`

	ClientID       int    `sql:"default:Null" json:"current_worker_id"`
	ClientInitials string `sql:",notnull" json:"client_initials"`
	ClientPhone    int    `json:"client_phone"`

	CurrentWorkerID       int    `sql:"default:Null"`
	CurrentWorkerInitials string `sql:"default:Null" json:"current_worker_initials"`
	CurrentWorkerPhone    int    `json:"current_worker_phone"`

	Color        string `json:"color"`
	Patina       string `json:"patina"`
	FasadArticle string `json:"fasad_article"`
	Material     string `json:"material"`

	CostCarpenter int `sql:",notnull" json:"cost_carpenter"`
	CostGrinder   int `sql:",notnull" json:"cost_grinder"`
	CostPainter   int `sql:",notnull" json:"cost_painter"`
	CostCollector int `sql:",notnull" json:"cost_collector"`

	Params []OrdersParam `json:"params"`
}

// OrdersChangesLogs - logs about changes in orders (/editOrder)
type OrdersChangesLogs struct {
	ID        int       `sql:",pk"`
	OrderID   int       `json:"order_id"`
	Date      time.Time `sql:"default:now()"`
	ManagerID int       `json:"manager_id"`
	Initials  string    `json:"initials"`

	Changes string `json:"changes"`
}
