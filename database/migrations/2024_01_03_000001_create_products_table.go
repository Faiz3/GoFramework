package migrations

import (
	"go-framework/app/models"
	"go-framework/framework/database"
	"go-framework/framework/logging"
)

type CreateProductsTable struct{}

func (m *CreateProductsTable) Signature() string {
	return "2024_01_03_000001_create_products_table"
}

func (m *CreateProductsTable) Description() string {
	return "Create products table"
}

func (m *CreateProductsTable) Up() error {
	logging.Info("Running migration: %s", m.Description())
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&models.Product{})
}

func (m *CreateProductsTable) Down() error {
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.Migrator().DropTable("products")
}
