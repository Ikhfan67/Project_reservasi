package config

import(
	"net/http"
	conn "project_reservasi/src/config"
)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	var message = "welcome"

	w.Write([]byte(message))

	var db = conn.Connect()
	

	defer db.Close()
	
}