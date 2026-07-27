package commands

import "fmt"

type RouteListCommand struct{}

func (c *RouteListCommand) Signature() string {
	return "route:list"
}

func (c *RouteListCommand) Description() string {
	return "List all registered routes"
}

func (c *RouteListCommand) Handle(args []string) error {
	fmt.Println("Registered Routes:")
	fmt.Println("  GET    /")
	fmt.Println("  GET    /admin/dashboard")
	fmt.Println("  GET    /users")
	fmt.Println("  POST   /users")
	fmt.Println("  GET    /users/:id")
	fmt.Println("  PUT    /users/:id")
	fmt.Println("  DELETE /users/:id")
	fmt.Println("  POST   /api/auth/login")
	fmt.Println("  POST   /api/auth/register")
	fmt.Println("  GET    /api/protected/profile")
	return nil
}
