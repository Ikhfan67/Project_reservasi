package home

import(
	"net/http"
	// "path"
	// "fmt"
	"html/template"
	

)


func HomeHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/user/home/home.html",
		"views/user/atribut/head.html",
		"views/user/atribut/top.html",
		"views/user/atribut/script.html",
		"views/user/atribut/footer.html",
		"views/user/section/about.html",
		"views/user/section/number.html",
		"views/user/section/service.html",
		"views/user/section/menu.html",
		"views/user/section/testimoni.html",
		"views/user/section/instagram.html",
	))

	var err = tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}
