package seeds

import (
	"go-framework/app/models"
	"go-framework/framework/auth"
	"go-framework/framework/database"
	"go-framework/framework/logging"
)

type UserSeeder struct{}

func (s *UserSeeder) Run() error {
	logging.Info("Seeding users table")

	hash, _ := auth.HashPassword("password")

	users := []models.User{
		{Name: "Admin", Email: "admin@example.com", Password: hash, Role: "admin"},
		{Name: "User", Email: "user@example.com", Password: hash, Role: "user"},
	}

	db, err := database.DB()
	if err != nil {
		return err
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	logging.Info("Seeded %d users", len(users))
	return nil
}
