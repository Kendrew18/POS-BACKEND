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

##  Dropdown Jenis Barang

Link: http://project-server.us.to:38600/JB/dropdown-jenis-barang

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")