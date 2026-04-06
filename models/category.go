package models

import "time"

type Category struct {
	CategoryID int       `gorm:"primaryKey;autoIncrement" json:"category_id"`
	Name       string    `gorm:"column:category_name;type:varchar(255);not null" json:"category_name"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
