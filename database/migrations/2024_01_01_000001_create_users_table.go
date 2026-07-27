package migrations

import (
	"go-framework/app/models"
	"go-framework/framework/database"
	"go-framework/framework/logging"
)

type CreateUsersTable struct{}

func (m *CreateUsersTable) Signature() string {
	return "2024_01_01_000001_create_users_table"
}

func (m *CreateUsersTable) Description() string {
	return "Create users table"
}

func (m *CreateUsersTable) Up() error {
	logging.Info("Running migration: %s", m.Description())
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&models.User{})
}

func (m *CreateUsersTable) Down() error {
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.Migrator().DropTable("users")
}
