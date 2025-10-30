package modules

import (
	"go-scaffolding/web-api/internal/router"
	"go-scaffolding/web-api/internal/server"
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
