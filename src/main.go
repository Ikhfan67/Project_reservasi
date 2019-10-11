package main

import(
	"fmt"
	i "reservasi/controller/home"
	conn "reservasi/config"
)

func main(){
	http.HandleFunc("/", i.HomeHandler)
	http.HandleFunc("/home", i.HomeHandler)
}
