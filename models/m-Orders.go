package models

import "time"

//Orders Order's table
type Orders struct {
	ID           int       `sql:",pk"`
	Date         time.Time `sql:"default:now()"`
	Rating       int
	Status       string
	ManagerPhone int `json:"manager_phone"`
	WorkerPhone  int `json:"worker_phone"`

	WorkersID int `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Workers   *Workers

	ManagersID int `sql:"on_delete:RESTRICT, on_update: CASCADE"`
	Managers   *Managers
}
