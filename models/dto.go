package models

type CreateAssetRequest struct {
	Code         string `json:"asset_code" validate:"required"`
	Name         string `json:"asset_name" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Condition    string `json:"condition" validate:"required"`
	CategoryID   int    `json:"category_id" validate:"required"`
	PurchaseYear int    `json:"purchase_year"`
	QRCode       string `json:"qr_code"`
	Image        string `json:"asset_image"`
}

type CreateCategoryRequest struct {
	Name string `json:"category_name" validate:"required,min=3"`
}

type CreateEmployeeRequest struct {
	Name  string `json:"employee_name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,min=8"`
}

type LoanRequest struct {
	AssetID    int    `json:"asset_id" validate:"required"`
	EmployeeID int    `json:"employee_id" validate:"required"`
	LoanDate   string `json:"loan_date"`
}

type CreateMaintenanceRequest struct {
	AssetID           int    `json:"asset_id" validate:"required"`
	IssueDescription  string `json:"issue_description" validate:"required,min=3"`
	MaintenanceDate   string `json:"maintenance_date"`
	MaintenanceStatus string `json:"maintenance_status" validate:"required"`
}
