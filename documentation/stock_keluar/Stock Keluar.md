# How TO Use API Stock Keluar
__________
##  Input Stock Keluar

Link: http://project-server.us.to:38600/SK/stock-keluar

Method: POST

Controllers:

    Request.Tanggal = c.FormValue("tanggal_stock_keluar")
	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Kode = c.FormValue("kode_toko")
	Request.Nama_penanggung_jawab = c.FormValue("penanggung_jawab")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Harga_jual = c.FormValue("harga_jual")

nb: 
- kode_stock : String Separator
- jumlah_barang: String Separator (untuk jumlah angkanya bisa koma (double))
- harga_jual: string separator

##  Dropdown Nama Toko

Link: http://project-server.us.to:38600/TK/dropdown-toko

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")

##  Dropdown Stock Barang

Link: http://project-server.us.to:38600/ST/dropdown-stock

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")

##  Read Stock Keluar

Link: http://project-server.us.to:38600/SK/stock-keluar

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")
	Request_filter.Kode_toko = c.FormValue("kode_toko")
	Request_filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_filter.Tanggal_2 = c.FormValue("tanggal_2")

