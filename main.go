package main

import (
	"go-framework/bootstrap"
	"go-framework/framework/logging"
)

func main() {
	app := bootstrap.NewApp()

	port := app.Config.Get("APP_PORT", "3000")
	logging.Info("Server starting on port %s", port)

	if err := app.Run(":" + port); err != nil {
		logging.Fatal("Failed to start server: %v", err)
	}
}
