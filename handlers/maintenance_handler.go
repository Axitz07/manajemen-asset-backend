package handlers

import (
	"asset-management/models"
	"asset-management/service"
	"asset-management/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MaintenanceHandler struct {
	service service.MaintenanceService
}

func NewMaintenanceHandler(s service.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{s}
}

func (h *MaintenanceHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1000"))

	data, total, err := h.service.GetAll(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": data, "total": total})
}

func (h *MaintenanceHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(data)
}

func (h *MaintenanceHandler) Create(c *fiber.Ctx) error {
	var req models.CreateMaintenanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}
	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	maintenanceDate := time.Now()
	if req.MaintenanceDate != "" {
		parsed, err := time.Parse("2006-01-02", req.MaintenanceDate)
		if err == nil {
			maintenanceDate = parsed
		}
	}

	item := models.Maintenance{
		AssetID:           req.AssetID,
		IssueDescription:  req.IssueDescription,
		MaintenanceDate:   maintenanceDate,
		MaintenanceStatus: req.MaintenanceStatus,
	}

	if err := h.service.Create(&item); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(item)
}

func (h *MaintenanceHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req models.CreateMaintenanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	maintenanceDate := time.Now()
	if req.MaintenanceDate != "" {
		parsed, err := time.Parse("2006-01-02", req.MaintenanceDate)
		if err == nil {
			maintenanceDate = parsed
		}
	}

	item := models.Maintenance{
		AssetID:           req.AssetID,
		IssueDescription:  req.IssueDescription,
		MaintenanceDate:   maintenanceDate,
		MaintenanceStatus: req.MaintenanceStatus,
	}

	if err := h.service.Update(id, &item); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(item)
}

func (h *MaintenanceHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}
