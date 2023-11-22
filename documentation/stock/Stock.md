# How TO Use API Stock
__________
##  Read Stock

Link: http://project-server.us.to:38600/ST/stock

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Kode_jenis_barang = c.FormValue("kode_jenis_barang")

nb: kode_jenis barang bisa kosong, jika menggunakan filter jenis barang maka dapat mengisi kode jenis barangnya

##  Dropdown Jenis Barang

Link: http://project-server.us.to:38600/JB/dropdown-jenis-barang

Method: GET

Controllers:

    Request.Kode_gudang = c.FormValue("kode_gudang")
