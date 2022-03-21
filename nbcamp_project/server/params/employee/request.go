package params

import (
	"nbcamp_project/server/models"
	"time"

	"github.com/google/uuid"
)

type EmployeeCreate struct {
	NIP     string
	Name    string
	Address string
}

func (e *EmployeeCreate) ParseToModel() *models.Employee {
	return &models.Employee{
		NIP:       e.NIP,
		Name:      e.Name,
		Address:   e.Address,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type EmployeeUpdate struct {
	NIP     string
	Name    string
	Address string
}

func (e *EmployeeUpdate) ParseToModel(id string) (*models.Employee, error) {
	newID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &models.Employee{
		NIP:       e.NIP,
		Name:      e.Name,
		Address:   e.Address,
		ID:        newID,
		UpdatedAt: time.Now(),
	}, nil
}
