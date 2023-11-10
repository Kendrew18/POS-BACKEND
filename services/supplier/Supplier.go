package supplier

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"net/http"
	"strconv"
)

func Input_Supplier(Request request.Input_Supplier_Request, Request_Barang request.Input_Barang_Supplier_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("supplier")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_supplier = "SP-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_supplier", "nama_supplier", "nomor_telpon", "kode_gudang").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	con_brg := db.CreateConGorm().Table("barang_supplier")

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock)
	Request_Barang.Kode_supplier = Request.Kode_supplier

	for i := 0; i < len(kode_stock); i++ {
		var Request_Barang_Input request.Input_Barang_Supplier_Request

		co = 0

		err = con_brg.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_Barang_Input.Co = co + 1
		Request_Barang_Input.Kode_barang_supplier = "SPB-" + strconv.Itoa(Request_Barang_Input.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = co
			return res, err.Error
		}

		err = con_brg.Select("co", "kode_barang_supplier", "kode_supplier", "kode_stock").Create(&Request_Barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

	}

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}

	return res, nil
}

func Read_Supplier(Request request.Read_Supplier_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Supplier_Response
	var obj_data response.Read_Supplier_Response

	con := db.CreateConGorm().Table("supplier")

	rows, err := con.Select("kode_supplier", "nama_supplier", "nomor_telpon").Where("kode_gudang = ?", Request.Kode_gudang).Rows()

	defer rows.Close()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		con_barang := db.CreateConGorm().Table("barang_supplier")
		var detail_data []response.Read_Barang_Supplier_Response
		rows.Scan(&obj_data.Kode_supplier, &obj_data.Nama_supplier, &obj_data.Nomor_telpon)

		err := con_barang.Select("kode_stock", "nama_barang").Joins("join stock on barang_supplier.kode_stock = stock.kode_stock").Where("kode_supplier = ?", obj_data.Kode_supplier).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		obj_data.Barang_supplier = detail_data

		data = append(data, obj_data)
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = data
	}

	return res, nil
}

func Dropdown_Nama_Supplier(Request request.Read_Supplier_Request) (response.Response, error) {

	var res response.Response
	var Nama_supplier []response.Read_Dropdown_Nama_Supplier_Response

	con := db.CreateConGorm().Table("supplier")

	err := con.Select("kode_supplier", "nama_supplier").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&Nama_supplier).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Nama_supplier
		return res, err
	}

	if Nama_supplier == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Nama_supplier

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Nama_supplier
	}

	return res, nil
}
