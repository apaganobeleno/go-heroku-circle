package models

//Gopher is a member of the BAQ Golang meetup
type Gopher struct {
	ID      int    `gorm:"primary_key;column:id"`
	Name    string `fako:"full_name"`
	Company string `fako:"company"`
}

//TableName defines our table name for GORM which in
//this case is not the default
func (Gopher) TableName() string {
	return "baq_gophers"
}
