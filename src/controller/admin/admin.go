package admin

import (
	"net/http"
	"html/template"
	"fmt"

	conn "project_reservasi/src/config"
	m "project_reservasi/src/model"

)

func AdminHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets/admin"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/admin/admin.html",
		"views/admin/atribut/head.html",
		"views/admin/atribut/script.html",

	))

	var err = tmpl.ExecuteTemplate(w, "admin", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func ReservasiDashboard(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets/admin"))

	var tmpl = template.Must(template.ParseFiles(
		"views/admin/ad_reservasi.html",
		"views/admin/atribut/head.html",
		"views/admin/atribut/script.html",

	))

	switch r.Method {
	case "GET":
		db := conn.Connect()

		selDB, err := db.Query("SELECT id, name, email, phone, date, time, person FROM booking ORDER BY id ASC")
		if err != nil {
			panic(err.Error())
		}
		emp := m.Reservasi{}
		res := []m.Reservasi{}
		for selDB.Next() {
			var id,phone int
			var name, email, date, time, person string
			err = selDB.Scan(&id, &name, &email, &phone, &date, &time, &person)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.Email = email
			emp.Phone = phone
			emp.Date = date
			emp.Time = time
			emp.Person = person
			res = append(res, emp)
		}
		// fmt.Println(emp.Name)
		tmpl.ExecuteTemplate(w, "ad_reservasi", res)
		defer db.Close()

	case "POST":
        fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	default:
        fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
	
}

func CommentsDashboard(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets/admin"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/admin/ad_comments.html",
		"views/admin/atribut/head.html",
		"views/admin/atribut/script.html",

	))

	var err = tmpl.ExecuteTemplate(w, "ad_comments", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}