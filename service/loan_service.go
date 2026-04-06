package service

import (
	"asset-management/models"
	"asset-management/repository"
	"fmt"
	"time"
)

type LoanService interface {
	GetAll(page, limit int) ([]models.AssetLoan, int64, error)
	Borrow(req models.LoanRequest) error
	Return(assetID int) error
	Delete(loanID int) error
}

type loanService struct {
	loanRepo    repository.LoanRepository
	assetRepo   repository.AssetRepository
	historyRepo repository.HistoryRepository
}

func NewLoanService(l repository.LoanRepository, a repository.AssetRepository, h repository.HistoryRepository) LoanService {
	return &loanService{loanRepo: l, assetRepo: a, historyRepo: h}
}

func (s *loanService) GetAll(page, limit int) ([]models.AssetLoan, int64, error) {
	return s.loanRepo.FindAll(page, limit)
}

func (s *loanService) Borrow(req models.LoanRequest) error {
	asset, err := s.assetRepo.FindByID(req.AssetID)
	if err != nil {
		return fmt.Errorf("asset not found")
	}

	if asset.Status != "Available" {
		return fmt.Errorf("asset tidak tersedia untuk dipinjam")
	}

	loanDate := time.Now()
	if req.LoanDate != "" {
		parsed, parseErr := time.Parse("2006-01-02", req.LoanDate)
		if parseErr == nil {
			loanDate = parsed
		}
	}

	oldStatus := asset.Status
	asset.Status = "In Use"
	if err := s.assetRepo.Update(&asset); err != nil {
		return err
	}

	loan := models.AssetLoan{
		AssetID:    req.AssetID,
		EmployeeID: req.EmployeeID,
		LoanDate:   loanDate,
		Status:     "Borrowed",
	}

	if err := s.loanRepo.Create(&loan); err != nil {
		return err
	}

	return s.historyRepo.Create(&models.AssetHistory{
		AssetID:   req.AssetID,
		OldStatus: oldStatus,
		NewStatus: "In Use",
		Note:      fmt.Sprintf("Asset %s dipinjam", asset.Name),
		ChangedAt: time.Now(),
	})
}

func (s *loanService) Return(assetID int) error {
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		return fmt.Errorf("asset not found")
	}

	loan, err := s.loanRepo.FindActiveLoan(assetID)
	if err != nil {
		return fmt.Errorf("active loan not found")
	}

	returnedAt := time.Now()
	loan.Status = "Returned"
	loan.ReturnDate = &returnedAt
	if err := s.loanRepo.Update(&loan); err != nil {
		return err
	}

	oldStatus := asset.Status
	asset.Status = "Available"
	if err := s.assetRepo.Update(&asset); err != nil {
		return err
	}

	return s.historyRepo.Create(&models.AssetHistory{
		AssetID:   assetID,
		OldStatus: oldStatus,
		NewStatus: "Available",
		Note:      fmt.Sprintf("Asset %s dikembalikan", asset.Name),
		ChangedAt: time.Now(),
	})
}

func (s *loanService) Delete(loanID int) error {
	loan, err := s.loanRepo.FindByID(loanID)
	if err != nil {
		return fmt.Errorf("loan not found")
	}

	if loan.Status == "Borrowed" {
		asset, assetErr := s.assetRepo.FindByID(loan.AssetID)
		if assetErr == nil {
			oldStatus := asset.Status
			asset.Status = "Available"
			if err := s.assetRepo.Update(&asset); err != nil {
				return err
			}

			if err := s.historyRepo.Create(&models.AssetHistory{
				AssetID:   asset.AssetID,
				OldStatus: oldStatus,
				NewStatus: "Available",
				Note:      fmt.Sprintf("Transaksi pinjam asset %s dibatalkan", asset.Name),
				ChangedAt: time.Now(),
			}); err != nil {
				return err
			}
		}
	}

	return s.loanRepo.Delete(loanID)
}
