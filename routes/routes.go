package routes

import (
	"asset-management/handlers"
	"asset-management/repository"
	"asset-management/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	assetRepo := repository.NewAssetRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	loanRepo := repository.NewLoanRepository(db)
	historyRepo := repository.NewHistoryRepository(db)
	maintenanceRepo := repository.NewMaintenanceRepository(db)

	assetSvc := service.NewAssetService(assetRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	employeeSvc := service.NewEmployeeService(employeeRepo)
	loanSvc := service.NewLoanService(loanRepo, assetRepo, historyRepo)
	historySvc := service.NewHistoryService(historyRepo)
	maintenanceSvc := service.NewMaintenanceService(maintenanceRepo, assetRepo, historyRepo)

	assetHandler := handlers.NewAssetHandler(assetSvc)
	categoryHandler := handlers.NewCategoryHandler(categorySvc)
	employeeHandler := handlers.NewEmployeeHandler(employeeSvc)
	loanHandler := handlers.NewLoanHandler(loanSvc)
	historyHandler := handlers.NewHistoryHandler(historySvc)
	maintenanceHandler := handlers.NewMaintenanceHandler(maintenanceSvc)

	asset := app.Group("/api/assets")
	asset.Get("/", assetHandler.GetAll)
	asset.Get("/:id", assetHandler.GetByID)
	asset.Post("/", assetHandler.Create)
	asset.Put("/:id", assetHandler.Update)
	asset.Delete("/:id", assetHandler.Delete)

	category := app.Group("/api/categories")
	category.Get("/", categoryHandler.GetAll)
	category.Get("/:id", categoryHandler.GetByID)
	category.Post("/", categoryHandler.Create)
	category.Put("/:id", categoryHandler.Update)
	category.Delete("/:id", categoryHandler.Delete)

	employee := app.Group("/api/employees")
	employee.Get("/", employeeHandler.GetAll)
	employee.Get("/:id", employeeHandler.GetByID)
	employee.Post("/", employeeHandler.Create)
	employee.Put("/:id", employeeHandler.Update)
	employee.Delete("/:id", employeeHandler.Delete)

	loan := app.Group("/api/loans")
	loan.Get("/", loanHandler.GetAll)
	loan.Post("/borrow", loanHandler.Borrow)
	loan.Post("/return/:asset_id", loanHandler.Return)
	loan.Delete("/:id", loanHandler.Delete)

	maintenance := app.Group("/api/maintenances")
	maintenance.Get("/", maintenanceHandler.GetAll)
	maintenance.Get("/:id", maintenanceHandler.GetByID)
	maintenance.Post("/", maintenanceHandler.Create)
	maintenance.Put("/:id", maintenanceHandler.Update)
	maintenance.Delete("/:id", maintenanceHandler.Delete)

	history := app.Group("/api/histories")
	history.Get("/", historyHandler.GetAll)
}
