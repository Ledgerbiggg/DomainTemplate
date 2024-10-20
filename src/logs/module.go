package logs

import (
	"go.uber.org/fx"
)

var Module = fx.Module("log",
	fx.Provide(NewConsoleLogger),
	fx.Invoke(func(l *ConsoleLogger) {
		l.Info("init console logger")
	}),
)
