package reservasi

import(
	"net/http"
	"html/template"
)

func ReservasiHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/user/reservasi/reservasi.html",
		"views/user/atribut/head.html",
		"views/user/atribut/top.html",
		"views/user/atribut/script.html",
		"views/user/atribut/footer.html",
		"views/user/section/reservasi.html",
	))

	var err = tmpl.ExecuteTemplate(w, "reservasi", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}