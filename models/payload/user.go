package payload

type (
	CreateUserRequest struct {
		Name     string `json:"nama" form:"nama" validate:"required,max=20"`
		Username string `json:"username" form:"username" validate:"required,username"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	CreateUserResponse struct {
		UserID uint   `json:"user_id"`
		Token  string `json:"token"`
	}

	LoginUserRequest struct {
		Email    string `json:"username" form:"username" validate:"required,username"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	LoginUserResponse struct {
		UserID uint   `json:"user_id"`
		Token  string `json:"token"`
	}
)
