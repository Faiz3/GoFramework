package commands

import (
	"go-framework/database/seeds"
	"go-framework/framework/logging"
)

type SeedCommand struct{}

func (c *SeedCommand) Signature() string {
	return "db:seed"
}

func (c *SeedCommand) Description() string {
	return "Seed the database with records"
}

func (c *SeedCommand) Handle(args []string) error {
	logging.Info("Seeding database...")

	seeder := &seeds.DatabaseSeeder{}
	return seeder.Run()
}
