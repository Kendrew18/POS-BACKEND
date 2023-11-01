package routes

import (
	"POS-BACKEND/controllers/jenis_barang"
	"POS-BACKEND/controllers/satuan_barang"
	"POS-BACKEND/controllers/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	US := e.Group("/US")

	//user management
	US.GET("/login", user.Login)

	JB := e.Group("/JB")

	//Jenis Barang
	JB.POST("/jenis-barang", jenis_barang.InputJenisBarang)
	JB.GET("/jenis-barang", jenis_barang.ReadJenisBarang)

	SB := e.Group("/SB")

	//Satuan Barang
	SB.POST("/satuan-barang", satuan_barang.InputSatuanBarang)
	SB.GET("/satuan-barang", satuan_barang.ReadSatuanBarang)

	//Satuan Barang
	//Stock Barang

	return e
}
