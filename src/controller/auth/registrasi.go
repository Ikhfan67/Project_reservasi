package auth

import(
	"net/http"
	"html/template"
)

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