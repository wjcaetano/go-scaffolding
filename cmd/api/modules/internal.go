package modules

import (
	"go-scaffolding/internal/router"
	"go-scaffolding/internal/server"
	"go.uber.org/fx"
)

var internalFactory = fx.Provide(
	router.NewRouter,
)

var InternalModule = fx.Options(
	internalFactory,
	fx.Invoke(
		server.StartHTTPServer,
	),
)
