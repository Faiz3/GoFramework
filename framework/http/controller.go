package http

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Handle(c *fiber.Ctx) error
}

type ResourceController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
}

type BaseController struct{}

func (bc *BaseController) Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func (bc *BaseController) Error(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

func (bc *BaseController) Paginated(c *fiber.Ctx, data interface{}, total, page, perPage int) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
		"meta": fiber.Map{
			"total":    total,
			"page":     page,
			"per_page": perPage,
		},
	})
}

func (bc *BaseController) Redirect(c *fiber.Ctx, url string) error {
	return c.Redirect(url)
}

func (bc *BaseController) View(c *fiber.Ctx, name string, data fiber.Map) error {
	return c.Render(name, data)
}
