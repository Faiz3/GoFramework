package routes

import (
	"go-framework/app/http/controllers"
	"go-framework/app/http/middleware"
	"go-framework/framework/routing"

	"github.com/gofiber/fiber/v2"
)

func Web(router *routing.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome", fiber.Map{
			"title": "Welcome to GoFramework",
		})
	})

	router.Group("/admin", func(r *routing.Router) {
		r.Get("/dashboard", &controllers.AdminController{})
	}, middleware.AdminAuth)

	router.Resource("users", &controllers.UserController{})
}
