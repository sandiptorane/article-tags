package routes

import (
	"article-tags/internal/handler"
	"article-tags/pkg/response"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers routes and handlers and return router
func RegisterRoutes(app *handler.Application) *gin.Engine {
	r := gin.Default()

	r.NoRoute(NoRoute)

	r.POST("/tags", app.AddTag)
	r.GET("/tags/:publication", app.GetFollowedTags)
	r.GET("/tags/:publication/popular", app.GetPopularTags)
	r.DELETE("/tags/:publication", app.DeleteTag)

	r.GET("/health", HealthCheck)

	return r
}

func NoRoute(c *gin.Context) {
	response.NotFound(c, "route not found", nil)
}

func HealthCheck(c *gin.Context) {
	response.Success(c, "ok", nil)
}
