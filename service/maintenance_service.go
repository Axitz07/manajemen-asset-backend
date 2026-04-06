package service

import (
	"asset-management/models"
	"asset-management/repository"
	"fmt"
	"time"
)

type MaintenanceService interface {
	GetAll(page, limit int) ([]models.Maintenance, int64, error)
	GetByID(id int) (models.Maintenance, error)
	Create(item *models.Maintenance) error
	Update(id int, item *models.Maintenance) error
	Delete(id int) error
}

type maintenanceService struct {
	repo        repository.MaintenanceRepository
	assetRepo   repository.AssetRepository
	historyRepo repository.HistoryRepository
}

func NewMaintenanceService(m repository.MaintenanceRepository, a repository.AssetRepository, h repository.HistoryRepository) MaintenanceService {
	return &maintenanceService{repo: m, assetRepo: a, historyRepo: h}
}

func (s *maintenanceService) GetAll(page, limit int) ([]models.Maintenance, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *maintenanceService) GetByID(id int) (models.Maintenance, error) {
	return s.repo.FindByID(id)
}

func (s *maintenanceService) Create(item *models.Maintenance) error {
	asset, err := s.assetRepo.FindByID(item.AssetID)
	if err != nil {
		return fmt.Errorf("asset not found")
	}

	oldStatus := asset.Status
	if item.MaintenanceStatus == "Done" {
		asset.Status = "Available"
	} else {
		asset.Status = "Maintenance"
	}

	if err := s.assetRepo.Update(&asset); err != nil {
		return err
	}

	if item.MaintenanceDate.IsZero() {
		item.MaintenanceDate = time.Now()
	}
	if err := s.repo.Create(item); err != nil {
		return err
	}

	note := fmt.Sprintf("%s masuk maintenance", asset.Name)
	if item.MaintenanceStatus == "Done" {
		note = fmt.Sprintf("%s selesai maintenance", asset.Name)
	}

	return s.historyRepo.Create(&models.AssetHistory{
		AssetID:   asset.AssetID,
		OldStatus: oldStatus,
		NewStatus: asset.Status,
		Note:      note,
		ChangedAt: time.Now(),
	})
}

func (s *maintenanceService) Update(id int, item *models.Maintenance) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	item.MaintenanceID = existing.MaintenanceID
	item.CreatedAt = existing.CreatedAt
	if item.MaintenanceDate.IsZero() {
		item.MaintenanceDate = existing.MaintenanceDate
	}

	if err := s.repo.Update(item); err != nil {
		return err
	}

	asset, err := s.assetRepo.FindByID(item.AssetID)
	if err != nil {
		return nil
	}

	oldStatus := asset.Status
	if item.MaintenanceStatus == "Done" {
		asset.Status = "Available"
	} else {
		asset.Status = "Maintenance"
	}

	if err := s.assetRepo.Update(&asset); err != nil {
		return err
	}

	if oldStatus != asset.Status {
		note := fmt.Sprintf("%s selesai maintenance", asset.Name)
		if asset.Status == "Maintenance" {
			note = fmt.Sprintf("%s masuk maintenance", asset.Name)
		}

		return s.historyRepo.Create(&models.AssetHistory{
			AssetID:   asset.AssetID,
			OldStatus: oldStatus,
			NewStatus: asset.Status,
			Note:      note,
			ChangedAt: time.Now(),
		})
	}

	return nil
}

func (s *maintenanceService) Delete(id int) error {
	return s.repo.Delete(id)
}
