package commands

import (
	"fmt"

	"go-framework/database/migrations"
	"go-framework/framework/logging"
)

type MigrateCommand struct{}

func (c *MigrateCommand) Signature() string {
	return "migrate"
}

func (c *MigrateCommand) Description() string {
	return "Run database migrations"
}

func (c *MigrateCommand) Handle(args []string) error {
	logging.Info("Running migrations...")

	migrationList := []interface{ Up() error; Down() error; Signature() string; Description() string }{
		&migrations.CreateUsersTable{},
		&migrations.CreateCategoriesTable{},
		&migrations.CreateProductsTable{},
	}

	for _, m := range migrationList {
		fmt.Printf("  Running: %s\n", m.Description())
		if err := m.Up(); err != nil {
			return fmt.Errorf("failed to run migration %s: %w", m.Signature(), err)
		}
		fmt.Printf("  Done: %s\n", m.Description())
	}

	logging.Info("All migrations completed successfully")
	return nil
}
