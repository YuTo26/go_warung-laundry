package controllers

import (
	"net/http"

	"go_warung-laundry/models/payload"
	"go_warung-laundry/usecase"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	req := new(payload.LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := usecase.LoginUser(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success login",
		"name":    resp.Name,
		"message": "hello" + resp.Username,
	})
}
