package controllers

import (
	"net/http"

<<<<<<< Updated upstream
	"go_warung-laundry/models"
=======
	"go_warung-laundry/models/payload"
>>>>>>> Stashed changes
	"go_warung-laundry/usecase"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
<<<<<<< Updated upstream
	user := models.User{}
	c.Bind(&user)

	err := usecase.LoginUser(&user)
=======
	req := new(payload.LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := usecase.LoginUser(req)
>>>>>>> Stashed changes
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
<<<<<<< Updated upstream
		"status":   "success login",
		"username": user,
=======
		"status":  "success login",
		"name":    resp.Name,
		"message": "hello" + resp.Username,
>>>>>>> Stashed changes
	})
}
