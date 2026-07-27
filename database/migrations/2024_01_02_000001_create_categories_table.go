package migrations

import (
	"go-framework/app/models"
	"go-framework/framework/database"
	"go-framework/framework/logging"
)

type CreateCategoriesTable struct{}

func (m *CreateCategoriesTable) Signature() string {
	return "2024_01_02_000001_create_categories_table"
}

func (m *CreateCategoriesTable) Description() string {
	return "Create categories table"
}

func (m *CreateCategoriesTable) Up() error {
	logging.Info("Running migration: %s", m.Description())
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&models.Category{})
}

func (m *CreateCategoriesTable) Down() error {
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.Migrator().DropTable("categories")
}
