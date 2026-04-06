package seeders

import (
	"asset-management/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	fmt.Println("Seeding database...")

	category := models.Category{
		Name: "Electronics",
	}
	if err := db.Create(&category).Error; err != nil {
		return err
	}

	employee := models.Employee{
		Name:  "Deril",
		Email: "deril@mail.com",
		Phone: "08123456789",
	}
	if err := db.Create(&employee).Error; err != nil {
		return err
	}

	assets := []models.Asset{
		{
			Code:         "AST001",
			Name:         "Laptop Asus",
			Status:       "In Use",
			Condition:    "Baik",
			CategoryID:   category.CategoryID,
			PurchaseYear: 2024,
			QRCode:       "QR-AST001",
		},
		{
			Code:         "AST002",
			Name:         "Printer Canon",
			Status:       "Maintenance",
			Condition:    "Rusak Ringan",
			CategoryID:   category.CategoryID,
			PurchaseYear: 2023,
			QRCode:       "QR-AST002",
		},
		{
			Code:         "AST003",
			Name:         "Monitor Dell",
			Status:       "Available",
			Condition:    "Baik",
			CategoryID:   category.CategoryID,
			PurchaseYear: 2025,
			QRCode:       "QR-AST003",
		},
	}

	for i := range assets {
		if err := db.Create(&assets[i]).Error; err != nil {
			return err
		}
	}

	loanDate := time.Now().AddDate(0, 0, -3)
	loan := models.AssetLoan{
		AssetID:    assets[0].AssetID,
		EmployeeID: employee.EmployeeID,
		LoanDate:   loanDate,
		Status:     "Borrowed",
	}
	if err := db.Create(&loan).Error; err != nil {
		return err
	}

	maintenance := models.Maintenance{
		AssetID:           assets[1].AssetID,
		IssueDescription:  "Printer cartridge tidak terbaca",
		MaintenanceDate:   time.Now().AddDate(0, 0, -1),
		MaintenanceStatus: "Repairing",
	}
	if err := db.Create(&maintenance).Error; err != nil {
		return err
	}

	histories := []models.AssetHistory{
		{
			AssetID:   assets[0].AssetID,
			OldStatus: "Available",
			NewStatus: "In Use",
			Note:      "Asset Laptop Asus dipinjam",
			ChangedAt: loanDate,
		},
		{
			AssetID:   assets[1].AssetID,
			OldStatus: "Available",
			NewStatus: "Maintenance",
			Note:      "Asset Printer Canon masuk maintenance",
			ChangedAt: maintenance.MaintenanceDate,
		},
	}

	for i := range histories {
		if err := db.Create(&histories[i]).Error; err != nil {
			return err
		}
	}

	fmt.Println("Seeding success ✅")
	return nil
}
