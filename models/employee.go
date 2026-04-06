package models

import "time"

type Employee struct {
	EmployeeID int       `gorm:"primaryKey;autoIncrement" json:"employee_id"`
	Name       string    `gorm:"column:employee_name" json:"employee_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
}
