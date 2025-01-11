package models

type Product struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`
}
