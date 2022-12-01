package time

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTraceability(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "traceability": nil})
}

func AddTraceability(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "traceability": nil})
}
