package config

import(
	"net/http"
)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	var message = "welcome"

	w.Write([]byte(message))

	db, err := Connect()
	if err != nil{
		fmt.Println(err.Error())
	}

	defer db.Close()
	
}