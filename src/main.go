package main

import(
	"fmt"
	"net/http"
	"database/sql"


	r "project_reservasi/src/route"
	loc "project_reservasi/src/loc_file"
)

var db *sql.DB

func main(){
	r.RouteView()
	r.RouteAuth()
	r.RouteAdmin()
	loc.LocFile()
	

	fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
