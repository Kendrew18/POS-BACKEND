package routes

import (
	"POS-BACKEND/controllers/jenis_barang"
	"POS-BACKEND/controllers/satuan_barang"
	"POS-BACKEND/controllers/stock"
	"POS-BACKEND/controllers/stock_masuk"
	"POS-BACKEND/controllers/supplier"
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

	//user management
	US := e.Group("/US")
	US.GET("/login", user.Login)

	//Jenis Barang
	JB := e.Group("/JB")
	JB.POST("/jenis-barang", jenis_barang.InputJenisBarang)
	JB.GET("/jenis-barang", jenis_barang.ReadJenisBarang)

	//Satuan Barang
	SB := e.Group("/SB")
	SB.POST("/satuan-barang", satuan_barang.InputSatuanBarang)
	SB.GET("/satuan-barang", satuan_barang.ReadSatuanBarang)

	//Stock Barang
	ST := e.Group("/ST")
	ST.POST("/stock-barang", stock.InputBarang)
	ST.GET("/stock-barang", stock.ReadBarang)

	//Supplier
	SP := e.Group("/SP")
	SP.POST("/supplier", supplier.InputSupplier)
	SP.GET("/supplier", supplier.ReadSupplier)
	SP.GET("/drop-down-nama-sup", supplier.DropdownNamaSupplier)

	//Stock Masuk
	SM := e.Group("/SM")
	SM.POST("/stock-masuk", stock_masuk.InputStockMasuk)
	SM.GET("/stock-masuk", stock_masuk.ReadStockMasuk)

	return e
}
