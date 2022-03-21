package controllers

import (
	"html/template"
	"nbcamp_project/server/utils"
	"net/http"
	"path"

	"github.com/NooBeeID/go-logging/messages"
)

type HomeController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct {
}

func NewHomeController() HomeController {
	return &homeController{}
}

func (*homeController) Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(path.Join("static", "pages/home/index.html"), utils.LayoutMaster)
	if err != nil {
		messages.SysError(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		messages.SysError(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
