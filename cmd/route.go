package main

import (
	"article-tags/handler"
	"article-tags/pkg/response"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers routes and handlers and return router
func RegisterRoutes(app *handler.Application) *gin.Engine {
	r := gin.Default()

	r.NoRoute(NoRoute)

	r.POST("/tags")
	r.PUT("/tags")
	r.GET("/tags/:publication")
	r.GET("/tags/popular")
	r.DELETE("/tags")

	r.GET("/health", HealthCheck)

	return r
}

func NoRoute(c *gin.Context) {
	response.NotFound(c, "route not found", nil)
}

func HealthCheck(c *gin.Context) {
	response.Success(c, "ok", nil)
}
