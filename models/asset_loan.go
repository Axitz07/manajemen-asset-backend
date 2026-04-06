package models

import "time"

type AssetLoan struct {
	LoanID     int        `gorm:"primaryKey;autoIncrement" json:"loan_id"`
	AssetID    int        `json:"asset_id"`
	EmployeeID int        `json:"employee_id"`
	LoanDate   time.Time  `json:"loan_date"`
	ReturnDate *time.Time `gorm:"default:null" json:"return_date"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
}
