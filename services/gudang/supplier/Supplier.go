package supplier

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"fmt"
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
		Request_Barang_Input.Kode_supplier = Request.Kode_supplier
		Request_Barang_Input.Kode_stock = kode_stock[i]

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = co
			return res, err.Error
		}

		err = con_brg.Select("co", "kode_barang_supplier", "kode_supplier", "kode_stock").Create(&Request_Barang_Input)

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

	rows, err := con.Select("kode_supplier", "nama_supplier", "nomor_telpon").Where("kode_gudang = ?", Request.Kode_gudang).Order("supplier.co ASC").Rows()

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

		err := con_barang.Select("barang_supplier.kode_stock", "nama_barang").Joins("join stock on barang_supplier.kode_stock = stock.kode_stock").Where("kode_supplier = ?", obj_data.Kode_supplier).Scan(&detail_data).Error

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

func Delete_Supplier(Request request.Delete_Supplier_Request) (response.Response, error) {
	var res response.Response

	var supplier []string
	var PO_supplier []string
	var refund []string
	sup := ""

	con_masuk := db.CreateConGorm().Table("stock_keluar_masuk")

	err := con_masuk.Select("kode").Joins("JOIN barang_stock_keluar_masuk bkm on bkm.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk").Where("kode = ? AND kode_stock = ?", Request.Kode_supplier, Request.Kode_stock).Scan(&supplier).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	fmt.Println(supplier)

	con_PO := db.CreateConGorm().Table("pre_order")

	err = con_PO.Select("kode_supplier").Joins("JOIN barang_pre_order bpo on bpo.kode_pre_order = pre_order.kode_pre_order").Where("kode_supplier = ? AND kode_stock = ?", Request.Kode_supplier, Request.Kode_stock).Scan(&PO_supplier).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	fmt.Println(PO_supplier)

	con_refund := db.CreateConGorm().Table("refund")

	err = con_refund.Select("kode_supplier").Joins("JOIN barang_refund bpo on bpo.kode_refund = refund.kode_refund").Where("kode_supplier = ? AND kode_stock=?", Request.Kode_supplier, Request.Kode_stock).Scan(&refund).Error

	fmt.Println(refund)

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	con_sup := db.CreateConGorm().Table("supplier")

	err = con_sup.Select("kode_supplier").Where("kode_supplier = ?", Request.Kode_supplier).Scan(&sup).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if supplier == nil && PO_supplier == nil && sup != "" && err == nil {

		var barang_supplier []string

		con_barang := db.CreateConGorm().Table("barang_supplier")

		err := con_barang.Where("kode_supplier=? AND kode_stock = ?", Request.Kode_supplier, Request.Kode_stock).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		con := db.CreateConGorm()

		err = con.Table("barang_supplier").Select("kode_barang_supplier").Where("kode_supplier=?", Request.Kode_supplier).Scan(&barang_supplier)

		fmt.Println(barang_supplier)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		if barang_supplier == nil {

			err = con.Table("supplier").Where("kode_supplier=?", Request.Kode_supplier).Delete("")

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

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Erorr karena ada condition yang tidak terpenuhi"
		res.Data = Request
		return res, err
	}

	return res, nil
}

func Dropdown_Barang_Supplier(Request request.Read_Barang_Supplier_Request) (response.Response, error) {
	var res response.Response
	var Barang_Supplier []response.Read_Barang_Supplier_Response

	con := db.CreateConGorm().Table("barang_supplier")

	err := con.Select("barang_supplier.kode_stock", "nama_barang").Joins("JOIN stock s on s.kode_stock = barang_supplier.kode_stock").Where("kode_supplier = ?", Request.Kode_supplier).Scan(&Barang_Supplier).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang_Supplier
		return res, err
	}

	if Barang_Supplier == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang_Supplier

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Barang_Supplier
	}

	return res, nil
}
