package test

import (
	"DomainTemplate/src/config"
	"DomainTemplate/src/logs"
	"DomainTemplate/src/services"
	"context"
	"go.uber.org/fx"
	"os"
	"testing"
)

var (
	gConfig *config.GConfig
	logger  *logs.ConsoleLogger
)

// TestMain executes before all tests
func TestMain(m *testing.M) {
	app := fx.New(
		config.Module,
		logs.Module,
		services.Module,
		fx.Populate(&gConfig, &logger),
	)

	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}
