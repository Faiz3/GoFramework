package requests

import (
	"go-framework/framework/validation"

	"github.com/gofiber/fiber/v2"
)

type StoreUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *StoreUserRequest) Rules() map[string]string {
	return map[string]string{
		"name":     "required|min:3|max:255",
		"email":    "required|email",
		"password": "required|min:8",
	}
}

func (r *StoreUserRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":     "Nama wajib diisi",
		"name.min":          "Nama minimal 3 karakter",
		"email.required":    "Email wajib diisi",
		"email.email":       "Format email tidak valid",
		"password.required": "Password wajib diisi",
		"password.min":      "Password minimal 8 karakter",
	}
}

func (r *StoreUserRequest) Validate(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	validator := validation.New()
	data := map[string]string{
		"name":     r.Name,
		"email":    r.Email,
		"password": r.Password,
	}

	if !validator.Validate(data, r.Rules()) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validator.Errors(),
		})
	}

	return c.Next()
}
