package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"nbcamp_project/server/helper"
	params "nbcamp_project/server/params/menu"
	"nbcamp_project/server/services"
	"net/http"
)

type MenuController interface {
	GetListMenu(w http.ResponseWriter, r *http.Request)
	CreateNewMenu(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	db *sql.DB
}

func NewMenuController(db *sql.DB) MenuController {
	return &menuController{db}
}

func (m *menuController) GetListMenu(w http.ResponseWriter, r *http.Request) {
	menus, err := services.NewMenuService(m.db).GetAllMenu()

	if err != nil {
		// handle not found
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))

		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	// handle not found
	if len(*menus) == 0 {
		helper.HandleNotFound(w, errors.New("NO DATA"))
		return
	}

	helper.HandleSuccess(w, menus)
}

func (m *menuController) CreateNewMenu(w http.ResponseWriter, r *http.Request) {
	var menu params.MenuCreate

	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	if menu.Category == "" || menu.Desc == "" || menu.Name == "" {

		helper.HandleBadRequest(w, err)
		return
	}

	_, err = services.NewMenuService(m.db).CreateMenu(&menu)
	if err != nil {
		helper.HandleInternalServerError(w, err)
		return
	}

	helper.HandleCreateSuccess(w, "Create new menu success !")
}
