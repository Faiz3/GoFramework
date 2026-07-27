package console

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Signature() string
	Description() string
	Handle(args []string) error
}

type Console struct {
	commands map[string]Command
}

func New() *Console {
	return &Console{
		commands: make(map[string]Command),
	}
}

func (c *Console) Register(cmd Command) {
	c.commands[cmd.Signature()] = cmd
}

func (c *Console) Run() {
	args := os.Args[1:]
	if len(args) == 0 {
		c.List()
		return
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	cmd, ok := c.commands[cmdName]
	if !ok {
		fmt.Printf("Command \"%s\" not found.\n", cmdName)
		os.Exit(1)
	}

	if err := cmd.Handle(cmdArgs); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func (c *Console) List() {
	fmt.Println("Available commands:")
	for sig, cmd := range c.commands {
		fmt.Printf("  %-20s %s\n", sig, cmd.Description())
	}
}

func ParseArgs(args []string) map[string]string {
	result := make(map[string]string)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg[2:], "=", 2)
			if len(parts) == 2 {
				result[parts[0]] = parts[1]
			} else {
				result[parts[0]] = "true"
			}
		}
	}
	return result
}
