package comments

import(
	"net/http"
	// "html/template"
	"log"
	// "fmt"

	conn "project_reservasi/src/config"
)

func Comments(w http.ResponseWriter,r *http.Request){
	db := conn.Connect()
    if r.Method == "POST" {
        name := r.FormValue("name")
        comment := r.FormValue("comment")
        insForm, err := db.Prepare("INSERT INTO comment(name, comments) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, comment)
        log.Println("INSERT: comments: " + comment + " sukses")
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}