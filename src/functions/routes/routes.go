package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Set(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Gin works!"})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "This route does not exist"})
	})
}
