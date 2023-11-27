package routes

import (
	"POS-BACKEND/controllers/audit"
	"POS-BACKEND/controllers/jenis_barang"
	"POS-BACKEND/controllers/pre_order"
	"POS-BACKEND/controllers/refund"
	"POS-BACKEND/controllers/satuan_barang"
	"POS-BACKEND/controllers/stock"
	"POS-BACKEND/controllers/stock_keluar"
	"POS-BACKEND/controllers/stock_masuk"
	"POS-BACKEND/controllers/supplier"
	"POS-BACKEND/controllers/toko"
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
	US.PUT("/change-lifo-fifo", user.Change_Fifo_Lifo)

	//Jenis Barang
	JB := e.Group("/JB")
	JB.POST("/jenis-barang", jenis_barang.InputJenisBarang)
	JB.GET("/jenis-barang", jenis_barang.ReadJenisBarang)
	JB.DELETE("/jenis-barang", jenis_barang.DeleteJenisBarang)
	JB.GET("/dropdown-jenis-barang", jenis_barang.DropdownJenisBarang)

	//Satuan Barang
	SB := e.Group("/SB")
	SB.POST("/satuan-barang", satuan_barang.InputSatuanBarang)
	SB.GET("/satuan-barang", satuan_barang.ReadSatuanBarang)
	SB.DELETE("/satuan-barang", satuan_barang.DeleteSatuanBarang)

	//Stock Barang
	ST := e.Group("/ST")
	ST.POST("/stock-barang", stock.InputBarang)
	ST.GET("/stock-barang", stock.ReadBarang)
	ST.DELETE("/stock-barang", stock.DeleteBarang)
	ST.GET("/stock", stock.ReadStock)
	ST.GET("/detail-stock", stock.Detailstock)
	ST.GET("/dropdown-stock", stock.DropdownStock)

	//Supplier
	SP := e.Group("/SP")
	SP.POST("/supplier", supplier.InputSupplier)
	SP.GET("/supplier", supplier.ReadSupplier)
	SP.GET("/drop-down-nama-sup", supplier.DropdownNamaSupplier)
	SP.DELETE("/supplier", supplier.DeleteSupplier)
	SP.GET("/drop-down-barang-sup", supplier.DropdownBarangSupplier)

	//Stock Masuk
	SM := e.Group("/SM")
	SM.POST("/stock-masuk", stock_masuk.InputStockMasuk)
	SM.GET("/stock-masuk", stock_masuk.ReadStockMasuk)
	SM.PUT("/stock-masuk", stock_masuk.UpdateBarangStockMasuk)
	SM.DELETE("/stock-masuk", stock_masuk.DeleteBarangStockMasuk)

	//Toko
	TK := e.Group("/TK")
	TK.POST("/toko", toko.InputToko)
	TK.GET("/toko", toko.ReadToko)
	TK.DELETE("/toko", toko.DeleteToko)
	TK.GET("/dropdown-toko", toko.DropdownNamaToko)

	//Stock_keluar
	SK := e.Group("/SK")
	SK.POST("/stock-keluar", stock_keluar.InputStockKeluar)
	SK.GET("/stock-keluar", stock_keluar.ReadStockKeluar)

	//Pre-Order
	PO := e.Group("/PO")
	PO.GET("/pre-order", pre_order.ReadPreOrder)
	PO.POST("/pre-order", pre_order.InputPreOrder)
	PO.PUT("/pre-order", pre_order.UpdatePreOrder)
	PO.DELETE("/pre-order", pre_order.DeletePreOrder)
	PO.PUT("/status-pre-order", pre_order.UpdateStatusPreOrder)

	//Refund
	RF := e.Group("/RF")
	RF.POST("/refund", refund.InputRefundSupplier)
	RF.GET("/refund", refund.ReadRefund)
	RF.PUT("/refund", refund.UpdateBarangRefund)
	RF.DELETE("/refund", refund.DeleteBarangRefund)
	RF.PUT("/update-status", refund.UpdateStatusRefund)

	//Audit
	AU := e.Group("AU")
	AU.GET("/audit", audit.ReadDataAwalAuditStock)

	return e
}
