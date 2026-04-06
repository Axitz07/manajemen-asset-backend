package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(page, limit int) ([]models.Employee, int64, error)
	FindByID(id int) (models.Employee, error)
	Create(employee *models.Employee) error
	Update(employee *models.Employee) error
	Delete(id int) error
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepo{db}
}

func (r *employeeRepo) FindAll(page, limit int) ([]models.Employee, int64, error) {
	var data []models.Employee
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&models.Employee{}).Count(&total)
	err := r.db.Limit(limit).Offset(offset).Find(&data).Error

	return data, total, err
}

func (r *employeeRepo) FindByID(id int) (models.Employee, error) {
	var data models.Employee
	err := r.db.First(&data, id).Error
	return data, err
}

func (r *employeeRepo) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepo) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepo) Delete(id int) error {
	return r.db.Delete(&models.Employee{}, id).Error
}
