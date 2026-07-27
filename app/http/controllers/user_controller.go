package controllers

import (
	"go-framework/framework/http"
	"go-framework/framework/logging"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	http.BaseController
}

func (uc *UserController) Handle(c *fiber.Ctx) error {
	return uc.Index(c)
}

func (uc *UserController) Index(c *fiber.Ctx) error {
	logging.Info("UserController.Index called")
	return uc.Success(c, fiber.Map{
		"users": []string{"John Doe", "Jane Doe"},
	})
}

func (uc *UserController) Store(c *fiber.Ctx) error {
	return uc.Success(c, fiber.Map{"message": "User created"})
}

func (uc *UserController) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	return uc.Success(c, fiber.Map{"user_id": id})
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	return uc.Success(c, fiber.Map{"message": "User updated"})
}

func (uc *UserController) Destroy(c *fiber.Ctx) error {
	return uc.Success(c, fiber.Map{"message": "User deleted"})
}

func (uc *UserController) Profile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	return uc.Success(c, fiber.Map{"user_id": userID})
}
