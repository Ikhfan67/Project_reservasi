package loc_file

import(
	"net/http"
)

func LocFile(){
	
	http.Handle("/css/", 
	http.StripPrefix("/css/", 
		http.FileServer(http.Dir("assets/css"))))
	
	http.Handle("/js/", 
	http.StripPrefix("/js/", 
		http.FileServer(http.Dir("assets/js"))))
	
	http.Handle("/image/", 
	http.StripPrefix("/image/", 
		http.FileServer(http.Dir("assets/images"))))
	
	http.Handle("/icon/", 
	http.StripPrefix("/icon/", 
		http.FileServer(http.Dir("assets/icon"))))

	http.Handle("/vidio/", 
	http.StripPrefix("/vidio/", 
		http.FileServer(http.Dir("assets/vidio"))))

	http.Handle("/login/css/", 
	http.StripPrefix("/login/css/", 
		http.FileServer(http.Dir("assets/login/css"))))

	http.Handle("/login/vendor/", 
	http.StripPrefix("/login/vendor/", 
		http.FileServer(http.Dir("assets/login/vendor"))))

	http.Handle("/login/fonts/", 
	http.StripPrefix("/login/fonts/", 
		http.FileServer(http.Dir("assets/login/fonts"))))

	http.Handle("/login/images/", 
	http.StripPrefix("/login/images/", 
		http.FileServer(http.Dir("assets/login/images"))))

	http.Handle("/login/js/", 
	http.StripPrefix("/login/js/", 
		http.FileServer(http.Dir("assets/login/js"))))

	http.Handle("/admin/js/", 
	http.StripPrefix("/admin/js/", 
		http.FileServer(http.Dir("assets/admin/js"))))

	http.Handle("/admin/css/", 
	http.StripPrefix("/admin/css/", 
		http.FileServer(http.Dir("assets/admin/css"))))
}