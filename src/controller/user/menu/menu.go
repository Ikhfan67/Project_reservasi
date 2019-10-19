package menu

import(
	"net/http"
	// "path"
	"html/template"
)

func MenuHandler(w http.ResponseWriter,r *http.Request){
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
		"views/user/menu/menu.html",
		"views/user/atribut/head.html",
		"views/user/atribut/top.html",
		"views/user/atribut/script.html",
		"views/user/atribut/footer.html",
		"views/user/section/menu.html",
	))

	var err = tmpl.ExecuteTemplate(w, "menu", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}