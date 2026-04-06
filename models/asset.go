package models

import "time"

type Asset struct {
	AssetID      int       `gorm:"primaryKey;autoIncrement" json:"asset_id"`
	Code         string    `gorm:"column:asset_code;type:varchar(100);unique;not null" json:"asset_code"`
	Name         string    `gorm:"column:asset_name;type:varchar(255);not null" json:"asset_name"`
	CategoryID   int       `gorm:"column:category_id;not null" json:"category_id"`
	PurchaseYear int       `json:"purchase_year"`
	Condition    string    `gorm:"type:varchar(50)" json:"condition"`
	Status       string    `gorm:"type:varchar(50)" json:"status"`
	QRCode       string    `gorm:"type:text" json:"qr_code"`
	Image        string    `gorm:"column:asset_image;type:text" json:"asset_image"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
