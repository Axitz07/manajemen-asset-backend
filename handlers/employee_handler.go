package handlers

import (
	"asset-management/models"
	"asset-management/service"
	"asset-management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(s service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{s}
}

func (h *EmployeeHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1000"))

	data, total, err := h.service.GetAll(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": data, "total": total})
}

func (h *EmployeeHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(data)
}

func (h *EmployeeHandler) Create(c *fiber.Ctx) error {
	var req models.CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}
	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	employee := models.Employee{Name: req.Name, Email: req.Email, Phone: req.Phone}
	if err := h.service.Create(&employee); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employee)
}

func (h *EmployeeHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req models.CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	employee := models.Employee{Name: req.Name, Email: req.Email, Phone: req.Phone}
	if err := h.service.Update(id, &employee); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employee)
}

func (h *EmployeeHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}
