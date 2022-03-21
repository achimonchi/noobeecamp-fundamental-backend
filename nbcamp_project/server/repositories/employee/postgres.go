package repositories

import (
	"database/sql"
	"nbcamp_project/server/helper"
	"nbcamp_project/server/models"
)

type employeeRepo struct {
	DB *sql.DB
}

// proses inject DB ke repositories
func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepo{
		DB: db,
	}
}

// proses implement method

func (e *employeeRepo) Save(employee *models.Employee) error {
	query := `
		INSERT INTO employees (
			id, nip, name, address, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		employee.ID, employee.NIP, employee.Name,
		employee.Address, employee.CreatedAt, employee.UpdatedAt,
	)

	return err
}

func (e *employeeRepo) FindAll() (*[]models.Employee, error) {
	query := `
		SELECT id, nip, name, address, created_at, updated_at
		FROM employees
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer helper.HandleError()

	defer stmt.Close()

	rows, err := stmt.Query()
	helper.HandlePanicIfError(err)

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee

		err = rows.Scan(
			&employee.ID, &employee.NIP, &employee.Name,
			&employee.Address, &employee.CreatedAt, &employee.UpdatedAt,
		)

		helper.HandlePanicIfError(err)

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (e *employeeRepo) UpdateByID(id string, employee *models.Employee) bool {
	query := `
		UPDATE employees
		SET nip=$2, name=$3, address=$4, updated_at=$5
		WHERE id=$1
	`

	stmt, err := e.DB.Prepare(query)
	defer helper.HandleError()

	helper.HandlePanicIfError(err)
	defer stmt.Close()

	_, err = stmt.Exec(id, employee.NIP, employee.Name, employee.Address, employee.UpdatedAt)
	helper.HandlePanicIfError(err)

	return true
}

func (e *employeeRepo) FindByID(id string) *models.Employee {
	defer helper.HandleError()
	query := `
		SELECT id, nip, name, address, created_at, updated_at
		FROM employees
		WHERE id=$1
		LIMIT 1
	`

	stmt, err := e.DB.Prepare(query)
	helper.HandlePanicIfError(err)

	row := stmt.QueryRow(id)

	var emp models.Employee
	err = row.Scan(
		&emp.ID, &emp.NIP, &emp.Name,
		&emp.Address, &emp.CreatedAt, &emp.UpdatedAt,
	)

	helper.HandlePanicIfError(err)

	return &emp
}

func (e *employeeRepo) DeleteByID(id string) {
	query := `
		DELETE FROM employees
		WHERE id=$1
	`

	defer helper.HandleError()

	stmt, err := e.DB.Prepare(query)
	helper.HandlePanicIfError(err)

	defer stmt.Close()

	_, err = stmt.Exec(id)
	helper.HandlePanicIfError(err)
}
