package reservasi

import(
	"net/http"
	"html/template"
	"log"
	"fmt"

	conn "project_reservasi/src/config"
	m "project_reservasi/src/model"
)

func ReservasiHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets"))

	// var data = map[string]interface{}{
	// 	"title": "Learning Golang Web",
	// 	"name":  "Kelompok 2",
	// }

	// var err = tmpl.ExecuteTemplate(w, "reservasi", data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	var tmpl = template.Must(template.ParseFiles(
		"views/user/reservasi/reservasi.html",
		"views/user/atribut/head.html",
		"views/user/atribut/top.html",
		"views/user/atribut/script.html",
		"views/user/atribut/footer.html",
	))
	
	switch r.Method {
	case "GET":
		db := conn.Connect()

		selDB, err := db.Query("SELECT id, name, date, time FROM booking ORDER BY id ASC")
		if err != nil {
			panic(err.Error())
		}
		emp := m.Reservasi{}
		res := []m.Reservasi{}
		for selDB.Next() {
			var id int
			var name, date, time string
			err = selDB.Scan(&id, &name, &date, &time)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.Date = date
			emp.Time = time
			res = append(res, emp)
		}
		// fmt.Println(emp.Name)
		tmpl.ExecuteTemplate(w, "reservasi", res)
		defer db.Close()

	case "POST":
        fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	default:
        fmt.Fprintf(w, "Sorry, only GET methods are supported.")
    }

}

func Reservasi(w http.ResponseWriter,r *http.Request){
	db := conn.Connect()
    if r.Method == "POST" {
        name := r.FormValue("name")
        email := r.FormValue("email")
        phone := r.FormValue("phone")
        date := r.FormValue("date")
        time := r.FormValue("time")
		person := r.FormValue("person")
        insForm, err := db.Prepare("INSERT INTO booking(name, email, phone, date, time, person) VALUES(?,?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, email, phone, date, time, person)
        log.Println("INSERT: name: " + name + " sukses")
    }
    defer db.Close()
    http.Redirect(w, r, "/reservasi", 301)
}