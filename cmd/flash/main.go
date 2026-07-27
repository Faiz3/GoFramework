package main

import "go-framework/app/console"

func main() {
	kernel := console.NewKernel()
	kernel.Run()
}
