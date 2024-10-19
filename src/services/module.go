package services

import (
	"DomainTemplate/src/logs"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(),
	fx.Invoke(func(c *logs.ConsoleLogger) {
		c.Info("service init success")
	}),
)
