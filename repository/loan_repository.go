package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type LoanRepository interface {
	FindAll(page, limit int) ([]models.AssetLoan, int64, error)
	FindByID(id int) (models.AssetLoan, error)
	Create(loan *models.AssetLoan) error
	FindActiveLoan(assetID int) (models.AssetLoan, error)
	Update(loan *models.AssetLoan) error
	Delete(id int) error
}

type loanRepo struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepo{db}
}

func (r *loanRepo) FindAll(page, limit int) ([]models.AssetLoan, int64, error) {
	var data []models.AssetLoan
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&models.AssetLoan{}).Count(&total)
	err := r.db.Order("created_at desc").Limit(limit).Offset(offset).Find(&data).Error

	return data, total, err
}

func (r *loanRepo) FindByID(id int) (models.AssetLoan, error) {
	var loan models.AssetLoan
	err := r.db.First(&loan, id).Error
	return loan, err
}

func (r *loanRepo) Create(loan *models.AssetLoan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepo) FindActiveLoan(assetID int) (models.AssetLoan, error) {
	var loan models.AssetLoan
	err := r.db.Where("asset_id = ? AND status = ?", assetID, "Borrowed").
		First(&loan).Error
	return loan, err
}

func (r *loanRepo) Update(loan *models.AssetLoan) error {
	return r.db.Save(loan).Error
}

func (r *loanRepo) Delete(id int) error {
	return r.db.Delete(&models.AssetLoan{}, id).Error
}
