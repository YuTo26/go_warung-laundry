package payload

import (
	"go_warung-laundry/models"

	"github.com/go-playground/validator/v10"
)

type (
	CreateUserRequest struct {
		Name     string `json:"nama" form:"nama" validate:"required,max=20"`
		Username string `json:"username" form:"username" validate:"required,alphanum,min=5,max=20"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	CreateUserResponse struct {
		UserID   uint   `json:"user_id"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}

	LoginRequest struct {
		Username string `json:"username" form:"username" validate:"required,alphanum,min=5,max=20"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	LoginResponse struct {
		UserID   uint   `json:"user_id"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Token    string `json:"token"`
	}
)

func (resp *CreateUserResponse) FromDomain(user *models.User) {
	resp.UserID = user.ID
	resp.Username = user.Username
	resp.Token = user.Token
}

func (resp *LoginResponse) FromDomain(user *models.User) {
	resp.UserID = user.ID
	resp.Name = user.Name
	resp.Username = user.Username
	resp.Token = user.Token
}

func ValidateRequest(req interface{}) error {
	validate := validator.New()
	return validate.Struct(req)
}
