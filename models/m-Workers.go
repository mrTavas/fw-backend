package models

// Workers Worker's table
type Workers struct {
	ID        int    `sql:",pk"`
	UUID      string `sql:", unique" json:"uuid"`
	Phone     int    `sql:",unique" json:"phone"`
	Password  string `sql:",notnull" json:"pass"`
	Initials  string `sql:",notnull" json:"initials"`
	Сarpenter bool   `sql:",notnull, default:false" json:"carpenter"`
	Grinder   bool   `sql:",notnull, default:false" json:"grinder"`
	Painter   bool   `sql:",notnull, default:false" json:"painter"`
	Collector bool   `sql:",notnull, default:false" json:"collector"`
}
