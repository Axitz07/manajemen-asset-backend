package service

import (
	"asset-management/models"
	"asset-management/repository"
)

type EmployeeService interface {
	GetAll(page, limit int) ([]models.Employee, int64, error)
	GetByID(id int) (models.Employee, error)
	Create(employee *models.Employee) error
	Update(id int, employee *models.Employee) error
	Delete(id int) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: r}
}

func (s *employeeService) GetAll(page, limit int) ([]models.Employee, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *employeeService) GetByID(id int) (models.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *employeeService) Create(employee *models.Employee) error {
	return s.repo.Create(employee)
}

func (s *employeeService) Update(id int, employee *models.Employee) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	employee.EmployeeID = existing.EmployeeID
	employee.CreatedAt = existing.CreatedAt
	return s.repo.Update(employee)
}

func (s *employeeService) Delete(id int) error {
	return s.repo.Delete(id)
}
