package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Token    string `gorm:"unique;not null" json:"-"`
}
