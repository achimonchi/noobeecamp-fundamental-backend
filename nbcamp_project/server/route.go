package server

import (
	"database/sql"
	"nbcamp_project/server/controllers"
	"net/http"

	"github.com/NooBeeID/go-logging/messages"
	"github.com/gorilla/mux"
)

func StartServer(router *mux.Router, db *sql.DB, port string) {
	buildRoute(router, db)

	// ini berfungsi menghandle sebuah folder yg isinya adalah
	// static file. Biasanya ini bersifat asset asset
	fileServer := http.FileServer(http.Dir("static/assets"))

	// ini proses untuk handle URL nya.
	// jadi saat di lakukan http://localhost:port/static/
	// maka akan menampilkan isi folder dari /static/assets/
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	messages.SysInfo("Server running at port %s", port)
	http.ListenAndServe(port, router)

}

func buildRoute(router *mux.Router, db *sql.DB) {
	homeRoute(router)
	employeeRoute(router, db)
	menuRoute(router, db)
}

func homeRoute(router *mux.Router) {
	homeController := controllers.NewHomeController()
	router.HandleFunc("/", homeController.Index)
}

func employeeRoute(router *mux.Router, db *sql.DB) {
	employeeController := controllers.NewEmployeeController(db)
	router.HandleFunc("/employees", employeeController.Index)
	router.HandleFunc("/employees/add", employeeController.Add)
	router.HandleFunc("/employees/update", employeeController.Update)
	router.HandleFunc("/employees/delete", employeeController.Delete)
}

func menuRoute(router *mux.Router, db *sql.DB) {
	menuController := controllers.NewMenuController(db)

	router.HandleFunc("/api/menus", menuController.GetListMenu).Methods("GET")
	router.HandleFunc("/api/menus", menuController.CreateNewMenu).Methods("POST")

}
