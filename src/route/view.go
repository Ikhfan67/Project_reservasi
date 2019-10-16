package route

import(
	"net/http"
	h "project_reservasi/src/controller/user/home"
	a "project_reservasi/src/controller/user/about"
	m "project_reservasi/src/controller/user/menu"
	c "project_reservasi/src/controller/user/contact"
	r "project_reservasi/src/controller/user/reservasi"
	l "project_reservasi/src/controller/auth"
	ad "project_reservasi/src/controller/admin"
)

func RouteView(){
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/home", h.HomeHandler)
	http.HandleFunc("/about", a.AboutHandler)
	http.HandleFunc("/menu", m.MenuHandler)
	http.HandleFunc("/contact", c.ContactHandler)
	http.HandleFunc("/reservasi", r.ReservasiHandler)
	http.HandleFunc("/admin", ad.AdminHandler)
}

func RouteAuth(){
	http.HandleFunc("/loginadmin", l.LoginHandler)
	http.HandleFunc("/registrasiadmin", l.RegistrasiHandler)
}