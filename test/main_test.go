package test

import (
	"DomainTemplate/src/config"
	"DomainTemplate/src/logs"
	"DomainTemplate/src/services"
	"context"
	"go.uber.org/fx"
)

func Inject() {
	app := fx.New(
		config.Module,
		logs.Module,
		services.Module,
	)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
