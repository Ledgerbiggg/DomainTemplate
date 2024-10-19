package main

import (
	"DomainTemplate/src/config"
	"DomainTemplate/src/logs"
	"DomainTemplate/src/services"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		logs.Module,
		services.Module,
	)
	app.Run()
}
