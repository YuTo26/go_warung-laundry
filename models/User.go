package models

type User struct {
<<<<<<< Updated upstream
	ID       uint   `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	JK       string `json:"jk"`
	Alamat   string `json:"alamat"`
	UserTelp string `json:"usertelp"`
	UserFoto string `json:"userfoto"`
	Level    string `json:"level"`
	Token    string `gorm:"-"`
=======
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Token    string `gorm:"unique;not null" json:"-"`
>>>>>>> Stashed changes
}
