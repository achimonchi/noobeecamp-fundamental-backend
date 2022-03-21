package services

import (
	"database/sql"
	"fmt"
	"nbcamp_project/server/helper"
	"nbcamp_project/server/models"
	params "nbcamp_project/server/params/employee"
	repositories "nbcamp_project/server/repositories/employee"
)

type EmployeeServices struct {
	EmployeeRepository repositories.EmployeeRepository
	DB                 *sql.DB
}

func NewEmployeeServices(db *sql.DB) *EmployeeServices {
	repositories := repositories.NewEmployeeRepository(db)
	return &EmployeeServices{
		EmployeeRepository: repositories,
		DB:                 db,
	}
}

func (e *EmployeeServices) CreateNewEmployee(request *params.EmployeeCreate) bool {
	defer helper.HandleError()
	model := request.ParseToModel()
	err := e.EmployeeRepository.Save(model)
	helper.HandlePanicIfError(err)

	return true
}

func (e *EmployeeServices) GetAllEmployees() *[]params.EmployeeSingleView {
	employees, err := e.EmployeeRepository.FindAll()
	defer helper.HandleError()
	helper.HandlePanicIfError(err)

	var empParams []params.EmployeeSingleView

	for _, emp := range *employees {
		empParams = append(empParams, *parseModelToParams(&emp))
	}
	return &empParams
}

func (e *EmployeeServices) FindEmployeeByID(id string) *params.EmployeeSingleView {
	employee := e.EmployeeRepository.FindByID(id)

	return parseModelToParams(employee)
}

func (e *EmployeeServices) UpdateEmployeeByID(id string, request *params.EmployeeUpdate) bool {
	model, err := request.ParseToModel(id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return e.EmployeeRepository.UpdateByID(id, model)
}

func parseModelToParams(employee *models.Employee) *params.EmployeeSingleView {
	return &params.EmployeeSingleView{
		ID:      employee.ID.String(),
		NIP:     employee.NIP,
		Name:    employee.Name,
		Address: employee.Address,
	}
}

func (e *EmployeeServices) DeleteEmployeeByID(id string) {
	e.EmployeeRepository.DeleteByID(id)
}
