package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`
}

type UserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
