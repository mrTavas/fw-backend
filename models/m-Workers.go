package models

// Workers Worker's table
type Workers struct {
	ID   int    `sql:",pk"`
	UUID string `sql:", unique" json:"uuid"`

	CurrentBalance int `sql:"default:0"`

	Phone     int    `sql:",unique" json:"phone"`
	Password  string `sql:",notnull" json:"pass"`
	Initials  string `sql:",notnull" json:"initials"`
	Сarpenter bool   `sql:"default:false" json:"carpenter"`
	Grinder   bool   `sql:"default:false" json:"grinder"`
	Painter   bool   `sql:"default:false" json:"painter"`
	Collector bool   `sql:"default:false" json:"collector"`
	ImageLink string `json:"image_link"`
}
