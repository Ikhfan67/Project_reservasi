package auth

import (
	"net/http"
	// "fmt"
	// "log"
	"html/template"
	// "database/sql"

	"github.com/kataras/go-sessions"
	// "github.com/kataras/go-sessions"
	"golang.org/x/crypto/bcrypt"
	// q "project_reservasi/src/controller/query"
	conn "project_reservasi/src/config"
	mo "project_reservasi/src/model"

)

func LoginHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets/login"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}
	var tmpl = template.Must(template.ParseFiles(
		"views/auth/login.html",
	))

	var err = tmpl.ExecuteTemplate(w, "login", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter,r *http.Request){
	db := conn.Connect()
    session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 {
		http.Redirect(w, r, "/", 302)
	}
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")

	// fmt.Println(username)
	// fmt.Println(password)

	var users = mo.User{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		password 
		FROM user WHERE username=?
		`, username).
		Scan(
			&users.Id,
			&users.Username,
			&users.Password,
		)

	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_tes == nil {
		//login success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		http.Redirect(w, r, "/loginadmin", 302)
	} else {
		//login failed
		http.Redirect(w, r, "/admin", 302)
	}

}
