package handlers

import (
	"asset-management/models"
	"asset-management/service"
	"asset-management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssetHandler struct {
	service service.AssetService
}

func NewAssetHandler(s service.AssetService) *AssetHandler {
	return &AssetHandler{s}
}

func (h *AssetHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1000"))

	data, total, err := h.service.GetAll(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": data, "total": total})
}

func (h *AssetHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(data)
}

func (h *AssetHandler) Create(c *fiber.Ctx) error {
	var req models.CreateAssetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	asset := models.Asset{
		Code:         req.Code,
		Name:         req.Name,
		Status:       req.Status,
		Condition:    req.Condition,
		CategoryID:   req.CategoryID,
		PurchaseYear: req.PurchaseYear,
		QRCode:       req.QRCode,
		Image:        req.Image,
	}

	if err := h.service.Create(&asset); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(asset)
}

func (h *AssetHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req models.CreateAssetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	asset := models.Asset{
		Code:         req.Code,
		Name:         req.Name,
		Status:       req.Status,
		Condition:    req.Condition,
		CategoryID:   req.CategoryID,
		PurchaseYear: req.PurchaseYear,
		QRCode:       req.QRCode,
		Image:        req.Image,
	}

	if err := h.service.Update(id, &asset); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(asset)
}

func (h *AssetHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}
