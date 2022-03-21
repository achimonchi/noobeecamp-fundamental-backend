package services

import (
	"database/sql"
	"nbcamp_project/server/models"
	params "nbcamp_project/server/params/menu"
	repositories "nbcamp_project/server/repositories/menu.go"
)

type MenuServices struct {
	MenuRepository repositories.MenuRepository
	DB             *sql.DB
}

func NewMenuService(db *sql.DB) *MenuServices {
	repository := repositories.NewMenuRepository(db)
	return &MenuServices{
		MenuRepository: repository,
		DB:             db,
	}
}

func (m *MenuServices) GetAllMenu() (*[]params.MenuSingleView, error) {
	menus, err := m.MenuRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return makeMenuListView(menus), nil
}

func makeMenuSingleView(model *models.Menu) *params.MenuSingleView {
	return &params.MenuSingleView{
		ID:        model.ID,
		Name:      model.Name,
		Desc:      model.Desc,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func makeMenuListView(models *[]models.Menu) *[]params.MenuSingleView {
	var menuListview []params.MenuSingleView
	for _, model := range *models {
		menuListview = append(menuListview, *makeMenuSingleView(&model))
	}

	return &menuListview
}

func (m *MenuServices) CreateMenu(request *params.MenuCreate) (bool, error) {
	model := request.ParseToModel()
	err := m.MenuRepository.CreateMenu(model)
	if err != nil {
		return false, err
	}

	return true, nil
}
