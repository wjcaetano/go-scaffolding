package modules

import "go.uber.org/fx"

func NewApp() *fx.App {
	options := []fx.Option{
		InternalModule,
	}

	return fx.New(
		fx.Options(options...),
	)
}
