# How TO Use API Pre Order
__________
##  Input Pre Order

Link: http://project-server.us.to:38600/PO/pre-order

Method: POST

Controllers:

    Request.Tanggal = c.FormValue("tanggal_pre_order")
	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Nama_penanggung_jawab = c.FormValue("nama_penanggung_jawab")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Tanggal_kadalurasa = c.FormValue("tanggal_kadalurasa")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Harga_pokok = c.FormValue("harga_pokok")

##  Read Pre Order

Link: http://project-server.us.to:38600/PO/pre-order

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")
	Request_filter.Kode_supplier = c.FormValue("kode_supplier")
	Request_filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_filter.Tanggal_2 = c.FormValue("tanggal_2")

##  Update Barang Pre Order

Link: http://project-server.us.to:38600/PO/pre-order

Method: PUT

Controllers:

    Request_kode.Kode_barang_pre_order = c.FormValue("kode_barang_pre_order")
	Request.Tanggal_kadaluarsa = c.FormValue("tanggal_kadaluarsa")
	Request.Jumlah_barang, _ = strconv.ParseFloat(c.FormValue("jumlah_barang"), 64)
	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

##  Delete Barang Pre Order

Link: http://project-server.us.to:38600/PO/pre-order

Method: DELETE

Controllers:

    Request_kode.Kode_barang_pre_order = c.FormValue("kode_barang_pre_order")


##  Update Status Pre Order

Link: http://project-server.us.to:38600/PO/status-pre-order

Method: PUT

Controllers:

	Request_kode.Kode_pre_order = c.FormValue("kode_pre_order")
	Request.Status, _ = strconv.Atoi(c.FormValue("status"))