package config

import(
	"net/http"
	// "path"
	"html/template"
)

func AboutHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets"))
	// var filepath = path.Join("views/home", "home.html")
	// var tmpl, err = template.ParseFiles(filepath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/about/about.html",
		"views/head.html",
	))

	var err = tmpl.ExecuteTemplate(w, "about", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}