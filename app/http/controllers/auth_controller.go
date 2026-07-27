package controllers

import (
	"time"

	"go-framework/framework/auth"
	fiberhttp "go-framework/framework/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	fiberhttp.BaseController
}

func (c *AuthController) Handle(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "AuthController"})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	if err := ctx.BodyParser(input); err != nil {
		return c.Error(ctx, "Invalid input", fiber.StatusBadRequest)
	}

	hashedPassword, _ := auth.HashPassword(input.Password)
	_ = hashedPassword

	token, err := auth.New("secret").GenerateToken(1, 24*time.Hour)
	if err != nil {
		return c.Error(ctx, "Failed to generate token", fiber.StatusInternalServerError)
	}

	return c.Success(ctx, fiber.Map{
		"token": token,
	})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	return c.Success(ctx, fiber.Map{"message": "Register endpoint"})
}
