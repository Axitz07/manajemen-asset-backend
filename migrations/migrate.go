package migrations

import (
	"asset-management/models"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println("Running migrations...")
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.AutoMigrate(
		&models.Employee{},
		&models.Category{},
		&models.Asset{},
		&models.AssetLoan{},
		&models.Maintenance{},
		&models.AssetHistory{},
	)
	if err != nil {
		return err
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	fmt.Println("Migration success")
	return nil
}

func DropAll(db *gorm.DB) error {
	fmt.Println("Dropping tables...")
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.Migrator().DropTable(
		&models.AssetHistory{},
		&models.Maintenance{},
		&models.AssetLoan{},
		&models.Asset{},
		&models.Category{},
		&models.Employee{},
	)
	if err != nil {
		return err
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	fmt.Println("Drop success")
	return nil
}
