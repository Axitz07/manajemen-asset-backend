package models

import "time"

type Maintenance struct {
	MaintenanceID     int       `gorm:"primaryKey;autoIncrement" json:"maintenance_id"`
	AssetID           int       `json:"asset_id"`
	IssueDescription  string    `gorm:"type:text" json:"issue_description"`
	MaintenanceDate   time.Time `json:"maintenance_date"`
	MaintenanceStatus string    `gorm:"type:varchar(50)" json:"maintenance_status"`
	CreatedAt         time.Time `json:"created_at"`
}
