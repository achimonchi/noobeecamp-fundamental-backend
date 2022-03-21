package repositories

import "nbcamp_project/server/models"

type EmployeeRepository interface {
	// ini berfungsi untuk menyimpan data ke database
	Save(employee *models.Employee) error
	FindAll() (*[]models.Employee, error)
	FindByID(id string) *models.Employee
	UpdateByID(id string, employee *models.Employee) bool
	DeleteByID(id string)
}
