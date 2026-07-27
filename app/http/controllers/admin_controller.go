package controllers

import (
	"go-framework/framework/http"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	http.BaseController
}

func (ac *AdminController) Handle(c *fiber.Ctx) error {
	return ac.View(c, "admin/dashboard", fiber.Map{
		"title": "Admin Dashboard",
	})
}
