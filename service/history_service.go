package service

import (
	"asset-management/models"
	"asset-management/repository"
)

type HistoryService interface {
	GetAll(page, limit int) ([]models.AssetHistory, int64, error)
}

type historyService struct {
	repo repository.HistoryRepository
}

func NewHistoryService(r repository.HistoryRepository) HistoryService {
	return &historyService{repo: r}
}

func (s *historyService) GetAll(page, limit int) ([]models.AssetHistory, int64, error) {
	return s.repo.FindAll(page, limit)
}
