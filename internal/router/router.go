package router

import (
	"github.com/gin-gonic/gin"
	"go-scaffolding/internal/routes"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	routes.RegisterBaseRoutes(router)

	return router
}
