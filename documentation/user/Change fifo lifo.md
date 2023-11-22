# How TO Use API Change Fifo Lifo
__________
##  Change Fifo Lifo

Link: http://project-server.us.to:38600/US/change-lifo-fifo

Method: PUT

Controllers:

    Request.Status, _ = strconv.Atoi(c.FormValue("status"))
	Kode_gudang := c.FormValue("kode_gudang")

nb: status adalah int , status 0 = lifo dan status 1 = fifo