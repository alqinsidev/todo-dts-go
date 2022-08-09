package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": message})
}

func Success(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}
