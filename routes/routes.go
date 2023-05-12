package routes

import (
	"net/http"

	"go_warung-laundry/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}
	//register & login uuser
	e.POST("/users/login", controllers.LoginUserController)
	e.POST("/users/register", controllers.RegisterUserController)

	//jenis laundry
	e.GET("/jenis_laundry", controllers.GetAllJenisLaundry)
	e.GET("/jenis_laundry/:id", controllers.GetJenisLaundryByID)
	e.POST("/jenis_laundry", controllers.CreateJenisLaundry)
	e.DELETE("/jenis_laundry/:id", controllers.DeleteJenisLaundryByID)
	e.PUT("/jenis_laundry/:id", controllers.UpdateJenisLaundryByID)

	//user_collection
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUserByID)
	e.PUT("/users/:id", controllers.UpdateUserByID)
	e.DELETE("/users/:id", controllers.DeleteUserByID)

	//pelanggan_collection
	e.POST("/pelanggan", controllers.CreatePelanggan)
	e.GET("/pelanggan", controllers.GetAllPelanggan)
	e.GET("/pelanggan/:id", controllers.GetPelangganByID)
	e.PUT("/pelanggan/:id", controllers.UpdatePelangganByID)
	e.DELETE("/pelanggan/:id", controllers.DeletePelangganByID)
	e.GET("/pelanggan/pembayaran", controllers.PelangganGetAllPembayaran)
	e.GET("/pelanggan/pembayaran/:id", controllers.PelangganGetPembayaranByID)
	e.POST("/pelanggan/pembayaran", controllers.PelangganCreatePembayaran)
	e.PUT("/pelanggan/pembayaran/:id", controllers.PelangganUpdatePembayaranByID)
	e.DELETE("/pelanggan/pembayaran/:id", controllers.PelangganDeletePembayaranByID)
	e.POST("/pelanggan/transaksi", controllers.PelangganCreateTransaksi)
	e.GET("/pelanggan/transaks", controllers.PelangganGetAllTransaksi)
	e.GET("/pelanggan/transaksi/:id", controllers.PelangganGetTransaksiByID)
	e.PUT("/pelanggan/transaksi/:id", controllers.PelangganUpdateTransaksiByID)
	e.DELETE("/pelanggan/transaksi", controllers.PelangganDeleteTransaksiByID)
}
