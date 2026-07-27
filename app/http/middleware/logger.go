package middleware

import (
	"time"

	"go-framework/framework/logging"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)
	logging.Info("[%s] %s %s - %dms",
		c.Method(),
		c.Path(),
		c.IP(),
		duration.Milliseconds(),
	)

	return err
}
