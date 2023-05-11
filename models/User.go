package models

type User struct {
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
}
