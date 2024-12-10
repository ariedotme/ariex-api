package controllers

import (
	"github.com/ariedotme/ariex-backend/application/services"
	"github.com/ariedotme/ariex-backend/application/usecases"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	LoginUserUseCase *usecases.LoginUserUseCase
	UserService      *services.UserService
}

type CreateUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginUser godoc
// @Summary Log in a user
// @Description Authenticate a user and return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body CreateUserPayload true "User Registration Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (c *UserController) Login(ctx *fiber.Ctx) error {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	token, err := c.LoginUserUseCase.Execute(payload.Username, payload.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

// RegisterUser godoc
// @Summary Register a new user (admin)
// @Description Allows registration of a new admin user
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Basic authentication header"
// @Param request body map[string]string true "User Registration Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /register [post]
func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	var payload CreateUserPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Use a validator to check for missing or invalid fields
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Proceed with user creation if validation passes
	user, err := c.UserService.CreateUser(payload.Username, payload.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully", "user_id": user.ID})
}
