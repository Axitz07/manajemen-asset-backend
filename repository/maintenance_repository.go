package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type MaintenanceRepository interface {
	FindAll(page, limit int) ([]models.Maintenance, int64, error)
	FindByID(id int) (models.Maintenance, error)
	Create(item *models.Maintenance) error
	Update(item *models.Maintenance) error
	Delete(id int) error
}

type maintenanceRepo struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) MaintenanceRepository {
	return &maintenanceRepo{db}
}

func (r *maintenanceRepo) FindAll(page, limit int) ([]models.Maintenance, int64, error) {
	var data []models.Maintenance
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&models.Maintenance{}).Count(&total)
	err := r.db.Order("maintenance_date desc").Limit(limit).Offset(offset).Find(&data).Error

	return data, total, err
}

func (r *maintenanceRepo) FindByID(id int) (models.Maintenance, error) {
	var data models.Maintenance
	err := r.db.First(&data, id).Error
	return data, err
}

func (r *maintenanceRepo) Create(item *models.Maintenance) error {
	return r.db.Create(item).Error
}

func (r *maintenanceRepo) Update(item *models.Maintenance) error {
	return r.db.Save(item).Error
}

func (r *maintenanceRepo) Delete(id int) error {
	return r.db.Delete(&models.Maintenance{}, id).Error
}
