package main

import(
	"fmt"
	"net/http"
	h "project_reservasi/src/controller/home"
	a "project_reservasi/src/controller/about"
	// conn "project_reservasi/src/config"
)

func main(){
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/home", h.HomeHandler)
	http.HandleFunc("/about", a.AboutHandler)

	http.Handle("/css/", 
	http.StripPrefix("/css/", 
		http.FileServer(http.Dir("assets/css"))))
	
	http.Handle("/js/", 
	http.StripPrefix("/js/", 
		http.FileServer(http.Dir("assets/js"))))
	
	http.Handle("/image/", 
	http.StripPrefix("/image/", 
		http.FileServer(http.Dir("assets/images"))))

	fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
