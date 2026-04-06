package handlers

import (
	"asset-management/models"
	"asset-management/service"
	"asset-management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	service service.LoanService
}

func NewLoanHandler(s service.LoanService) *LoanHandler {
	return &LoanHandler{s}
}

func (h *LoanHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1000"))

	data, total, err := h.service.GetAll(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": data, "total": total})
}

func (h *LoanHandler) Borrow(c *fiber.Ctx) error {
	var req models.LoanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.Borrow(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "asset borrowed successfully"})
}

func (h *LoanHandler) Return(c *fiber.Ctx) error {
	assetID, err := strconv.Atoi(c.Params("asset_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid asset id"})
	}
	if err := h.service.Return(assetID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "asset returned successfully"})
}

func (h *LoanHandler) Delete(c *fiber.Ctx) error {
	loanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid loan id"})
	}
	if err := h.service.Delete(loanID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "loan deleted successfully"})
}
