package handlers

import (
	"asset-management/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	service service.HistoryService
}

func NewHistoryHandler(s service.HistoryService) *HistoryHandler {
	return &HistoryHandler{s}
}

func (h *HistoryHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1000"))

	data, total, err := h.service.GetAll(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": data, "total": total})
}
