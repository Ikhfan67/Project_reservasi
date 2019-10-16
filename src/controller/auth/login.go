package auth

import(
	"net/http"
	"fmt"
	"html/template"
	"database/sql"

	// "github.com/kataras/go-sessions"
	// "golang.org/x/crypto/bcrypt"
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

	// session := sessions.Start(w, r)
	// if len(session.GetString("username")) != 0 {
	// 	http.Redirect(w, r, "/admin", 302)
	// }
	// if r.Method != "POST" {
	// 	http.ServeFile(w, r, "views/auth/login.html")
	// 	fmt.Println("wrong")
	// 	return
	// }
	// username := r.FormValue("username")
	// password := r.FormValue("password")

	// users := QueryUser(username)

	// //deskripsi dan compare password
	// var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	// if password_tes == nil {
	// 	//login success
	// 	session := sessions.Start(w, r)
	// 	session.Set("username", users.Username)
	// 	http.Redirect(w, r, "/admin", 302)
	// } else {
	// 	//login failed
	// 	http.Redirect(w, r, "/login", 302)
	// }
	
}

var db *sql.DB
var err error

func QueryUser(username string) mo.User {
	var users = mo.User{}
	err := db.QueryRow(`
		SELECT id, 
		username,
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.Id,
			&users.Username,
			&users.Password,
		)
		if err != nil {
			fmt.Println(err)
		}
	return users
}