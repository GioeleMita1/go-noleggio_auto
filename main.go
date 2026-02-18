package main

import (
	"noleggio_auto/pkg/plugin/dependencies"
	"noleggio_auto/pkg/plugin"
	
)

func main() {
	
	dependencies := dependencies.Init()
	plugin.RegisterRoutes(dependencies.AutoService)
}