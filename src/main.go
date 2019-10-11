package main

import(
	"fmt"
	i "project_reservasi/src/controller/home"
	conn "project_reservasi/src/config"
)

func main(){
	http.HandleFunc("/", i.HomeHandler)
	http.HandleFunc("/home", i.HomeHandler)
}
