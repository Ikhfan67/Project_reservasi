package main

import(
	"fmt"
	"net/http"
	"database/sql"

	r "project_reservasi/src/route"
	loc "project_reservasi/src/loc_file"
	conn "project_reservasi/src/config"
)


func main(){
	r.RouteView()
	r.RouteAuth()
	loc.LocFile()

	var db *sql.DB
	conn.Connect()
	defer db.Close()


	fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
