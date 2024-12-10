package controllers

import (
	"github.com/ariedotme/ariex-backend/application/usecases"
	"github.com/ariedotme/ariex-backend/domain/entities"
	"github.com/gofiber/fiber/v2"
)

type ContactController struct {
	SendContactUseCase *usecases.SendContactUseCase
}

// SendContact godoc
// @Summary Send a contact message
// @Description Allows a user to send a message through the contact form, which will be sent via email
// @Tags Contact
// @Accept json
// @Produce json
// @Param request body entities.Contact true "Contact Information"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /contact [post]
func (c *ContactController) SendContact(ctx *fiber.Ctx) error {
	var contact entities.Contact
	if err := ctx.BodyParser(&contact); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}
	if err := c.SendContactUseCase.Execute(&contact); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Contact sent successfully"})
}
