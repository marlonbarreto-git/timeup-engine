package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	time "timeup-engine/cmd/api/internal/domain/http/handler"
)

func StartGin() error {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/traceability", time.GetTraceability)
		api.POST("/traceability", time.AddTraceability)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router.Run(":8000")
}
