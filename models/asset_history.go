package models

import "time"

type AssetHistory struct {
	HistoryID int       `gorm:"primaryKey;autoIncrement" json:"history_id"`
	AssetID   int       `json:"asset_id"`
	OldStatus string    `json:"old_status"`
	NewStatus string    `json:"new_status"`
	Note      string    `json:"note"`
	ChangedAt time.Time `json:"changed_at"`
}
