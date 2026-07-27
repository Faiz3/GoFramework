package console

import (
	"go-framework/app/console/commands"
	"go-framework/framework/console"
)

type Kernel struct {
	Console *console.Console
}

func NewKernel() *Kernel {
	k := &Kernel{
		Console: console.New(),
	}
	k.registerCommands()
	return k
}

func (k *Kernel) registerCommands() {
	k.Console.Register(&commands.ServeCommand{})
	k.Console.Register(&commands.MigrateCommand{})
	k.Console.Register(&commands.SeedCommand{})
	k.Console.Register(&commands.KeyGenerateCommand{})
	k.Console.Register(&commands.RouteListCommand{})
	k.Console.Register(&commands.MakeCommand{})
}

func (k *Kernel) Run() {
	k.Console.Run()
}
