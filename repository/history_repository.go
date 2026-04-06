package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	FindAll(page, limit int) ([]models.AssetHistory, int64, error)
	Create(history *models.AssetHistory) error
}

type historyRepo struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepo{db}
}

func (r *historyRepo) FindAll(page, limit int) ([]models.AssetHistory, int64, error) {
	var data []models.AssetHistory
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&models.AssetHistory{}).Count(&total)
	err := r.db.Order("changed_at desc").Limit(limit).Offset(offset).Find(&data).Error

	return data, total, err
}

func (r *historyRepo) Create(history *models.AssetHistory) error {
	return r.db.Create(history).Error
}
