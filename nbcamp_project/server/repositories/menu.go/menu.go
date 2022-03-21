package repositories

import "nbcamp_project/server/models"

type MenuRepository interface {
	GetAll() (*[]models.Menu, error)
	CreateMenu(menu *models.Menu) error
}
