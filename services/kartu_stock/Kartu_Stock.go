package kartu_stock

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
	"time"
)

func Read_Kartu_Stock(Request request.Read_Kartu_Stock_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("audit")

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	bulanthn := date.Format("2006-01")

	state := "'" + bulanthn + "-%'"

	tanggal := ""

	err := con.Select("MAX(tanggal)").Where("kode_gudang = ? && tanggal like ?", Request.Kode_gudang, state).Scan(&tanggal).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	fmt.Println(tanggal)

	return res, nil
}
