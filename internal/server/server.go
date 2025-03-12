package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"
)

const (
	readHeaderTimeout = 5 * time.Second
)

func StartHTTPServer(lc fx.Lifecycle, router *gin.Engine) {
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: readHeaderTimeout,
		Handler:           router,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					log.Println("starting server on :8080")
					if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.Fatalf("could not start server: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("stopping server")
				return server.Shutdown(ctx)
			},
		})
}
