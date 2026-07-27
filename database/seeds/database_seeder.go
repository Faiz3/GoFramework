package seeds

import "go-framework/framework/logging"

type DatabaseSeeder struct{}

func (s *DatabaseSeeder) Run() error {
	logging.Info("Running database seeder")

	userSeeder := &UserSeeder{}
	if err := userSeeder.Run(); err != nil {
		return err
	}

	return nil
}
