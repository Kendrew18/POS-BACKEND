# How TO Use API Stock Masuk
__________
##  Input Stock Masuk

Link: http://project-server.us.to:38600/SM/stock-masuk

Method: POST

Controllers:

    Request.Tanggal = c.FormValue("tanggal_stock_masuk")
	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Kode = c.FormValue("kode_supplier")
	Request.Nama_penanggung_jawab = c.FormValue("nama_penanggung_jawab")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Tanggal_kadalurasa = c.FormValue("tanggal_kadalurasa")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Harga_pokok = c.FormValue("harga_pokok")

nb: 
- kode_stock : String Separator
- tanggal_kadaluarsa : String Separator
- jumlah_barang: String Separator (untuk jumlah angkanya bisa koma (double))
- Harga_pokok: string separator

##  Dropdown Nama Supplier

Link: http://project-server.us.to:38600/SP/drop-down-nama-sup

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")

## Dropdown Barang Supplier

Link: http://project-server.us.to:38600/SP/drop-down-barang-sup
Method: GET

Controllers:

    Request.Kode_supplier = c.FormValue("kode_supplier")

## Read Stock Masuk

Link: http://project-server.us.to:38600

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")
	Request_filter.Kode_supplier = c.FormValue("kode_supplier")
	Request_filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_filter.Tanggal_2 = c.FormValue("tanggal_2")

## Update Barang Stock Masuk

Link: http://project-server.us.to:38600/SM/stock-masuk

Method: PUT

Controllers:

    Request_kode.Kode_barang_keluar_masuk = c.FormValue("kode_barang_keluar_masuk")
	Request.Tanggal_kadaluarsa = c.FormValue("tanggal_kadaluarsa")
	Request.Jumlah_barang, _ = strconv.ParseFloat(c.FormValue("jumlah_barang"), 64)
	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

## Delete Stock Masuk

Link: http://project-server.us.to:38600/SM/stock-masuk

Method: DELETE

Controllers:

    Request_kode.Kode_barang_keluar_masuk = c.FormValue("kode_barang_keluar_masuk")