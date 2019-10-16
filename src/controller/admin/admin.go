package admin

import(
	"net/http"
	"html/template"
)

func AdminHandler(w http.ResponseWriter,r *http.Request){
	http.FileServer(http.Dir("assets/admin"))

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Kelompok 2",
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/admin/admin.html",
	))

	var err = tmpl.ExecuteTemplate(w, "admin", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	// if r.Method == "POST" {
    //     var tmpl = template.Must(template.New("admin").ParseFiles("view.html"))

    //     if err := r.ParseForm(); err != nil {
    //         http.Error(w, err.Error(), http.StatusInternalServerError)
    //         return
    //     }

    //     var name = r.FormValue("name")
    //     var message = r.Form.Get("message")

    //     var data = map[string]string{"name": name, "message": message}

    //     if err := tmpl.Execute(w, data); err != nil {
    //         http.Error(w, err.Error(), http.StatusInternalServerError)
    //     }
    //     return
    // }

    // http.Error(w, "", http.StatusBadRequest)
}