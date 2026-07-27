package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CORS(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-TOKEN")

	if c.Method() == "OPTIONS" {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Next()
}
