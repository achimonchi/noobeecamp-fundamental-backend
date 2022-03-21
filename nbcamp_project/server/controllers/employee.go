package controllers

import (
	"database/sql"
	"html/template"
	"log"
	"nbcamp_project/server/apps/web"
	params "nbcamp_project/server/params/employee"
	"nbcamp_project/server/services"
	"nbcamp_project/server/utils"
	"net/http"
	"path"

	"github.com/NooBeeID/go-logging/messages"
)

type EmployeeController interface {
	Index(rw http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type employeeController struct {
	db *sql.DB
}

func NewEmployeeController(db *sql.DB) EmployeeController {
	return &employeeController{db}
}

func (e *employeeController) Index(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/index.html"), utils.LayoutMaster)
	if err != nil {
		messages.SysError(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	emps := services.NewEmployeeServices(e.db).GetAllEmployees()

	response := web.RenderWeb{
		Title: "Halaman Employee",
		Data:  emps,
	}

	err = tmpl.Execute(rw, response)
	if err != nil {
		messages.SysError(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e *employeeController) Add(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		form(w, "add.html", nil, "Halaman Tambah Pegawai")
	} else if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		request := params.EmployeeCreate{
			NIP:     r.Form.Get("nip"),
			Name:    r.Form.Get("name"),
			Address: r.Form.Get("address"),
		}

		empServices := services.NewEmployeeServices(e.db)

		add := empServices.CreateNewEmployee(&request)

		msg := ""
		if add {
			msg = `
			<script>
				alert("Tambah data pegawai berhasil !")
				window.location.href="../employees"
			</script>
		`
		} else {
			msg = `
				<script>
					alert("Tambah data pegawai gagal !")
					window.location.href="../employees"
				</script>
			`
		}

		w.Write([]byte(msg))

	}
}

func form(rw http.ResponseWriter, formName string, data interface{}, title string) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/"+formName), utils.LayoutMaster)
	if err != nil {
		messages.SysError(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	response := web.RenderWeb{
		Title: title,
		Data:  data,
	}

	err = tmpl.Execute(rw, response)
	if err != nil {
		messages.SysError(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e *employeeController) Update(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	query := r.URL.Query()
	id := query["id"][0]
	empServices := services.NewEmployeeServices(e.db)

	if method == "GET" {
		emp := empServices.FindEmployeeByID(id)
		form(w, "update.html", emp, "Halaman Update Pegawai")
	} else if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		request := params.EmployeeUpdate{
			NIP:     r.Form.Get("nip"),
			Name:    r.Form.Get("name"),
			Address: r.Form.Get("address"),
		}

		add := empServices.UpdateEmployeeByID(id, &request)

		msg := ""
		if add {
			msg = `
			<script>
				alert("Ubah data pegawai berhasil !")
				window.location.href="../employees"
			</script>
		`
		} else {
			msg = `
				<script>
					alert("Ubah data pegawai gagal !")
					window.location.href="../employees"
				</script>
			`
		}

		w.Write([]byte(msg))

	}
}

func (e *employeeController) Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query["id"][0]

	services.NewEmployeeServices(e.db).DeleteEmployeeByID(id)

	msg := `<script>
	alert("Ubah data pegawai berhasil !")
	window.location.href="../employees"
</script>`

	w.Write([]byte(msg))

}
