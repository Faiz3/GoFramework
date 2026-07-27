package commands

import (
	"fmt"

	"go-framework/bootstrap"
	"go-framework/framework/console"
)

type ServeCommand struct{}

func (c *ServeCommand) Signature() string {
	return "serve"
}

func (c *ServeCommand) Description() string {
	return "Start the development server"
}

func (c *ServeCommand) Handle(args []string) error {
	app := bootstrap.NewApp()
	parsedArgs := console.ParseArgs(args)

	port := parsedArgs["port"]
	if port == "" {
		port = app.Config.Get("APP_PORT", "3000")
	}

	fmt.Printf("Starting server on :%s\n", port)
	return app.Run(":" + port)
}
