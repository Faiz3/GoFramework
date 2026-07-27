package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyCSRF(c *fiber.Ctx) error {
	if c.Method() == "POST" || c.Method() == "PUT" || c.Method() == "DELETE" || c.Method() == "PATCH" {
		token := c.Get("X-CSRF-TOKEN")
		if token == "" {
			token = c.FormValue("_token")
		}
		if token == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "CSRF token mismatch",
			})
		}
	}
	return c.Next()
}

func CSRF(c *fiber.Ctx) error {
	return c.Next()
}
