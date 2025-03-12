package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterBaseRoutes(router *gin.Engine) {
	RegisterEndpoints(router, http.MethodGet, "/", homeHandler)
	RegisterEndpoints(router, http.MethodGet, "health", healthHandler)
}

func homeHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World")
}

func healthHandler(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func RegisterDomainRoutes(router *gin.Engine, registers ...EndpointRegister) {
	for _, register := range registers {
		register.RegisterEndpoint(router)
	}
}

func RegisterEndpoints(router *gin.Engine, method, pattern string, handler gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		router.GET(pattern, handler)
	case http.MethodPost:
		router.POST(pattern, handler)
	case http.MethodPut:
		router.PUT(pattern, handler)
	case http.MethodPatch:
		router.PATCH(pattern, handler)
	case http.MethodDelete:
		router.DELETE(pattern, handler)
	default:
	}
}

type EndpointRegister interface {
	RegisterEndpoint(router *gin.Engine)
}
