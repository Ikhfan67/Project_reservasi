package auth

import (
	"net/http"
	"html/template"
	"log"
	// "fmt"
	"database/sql"
	// "path"
	
	// "golang.org/x/crypto/bcrypt"

	// q "project_reservasi/src/controller/query"
	// mo "project_reservasi/src/model"
	conn "project_reservasi/src/config"
	

)

var db *sql.DB
var err error

func RegistrasiHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets/login"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}
	var tmpl = template.Must(template.ParseFiles(
		"views/auth/register.html",
	))

	var err = tmpl.ExecuteTemplate(w, "registrasi", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Register(w http.ResponseWriter,r *http.Request){
	db := conn.Connect()
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")
        insForm, err := db.Prepare("INSERT INTO user(username, password) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(username, password)
        log.Println("INSERT: username: " + username + " | password: " + password)
    }
    defer db.Close()
    http.Redirect(w, r, "/loginadmin", 301)
}
