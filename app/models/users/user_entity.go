package users_model

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
