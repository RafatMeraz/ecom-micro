package handler

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"health": "ok",
	})
}
