package routes

import (
	"net/http"

	"github.com/edwintrumpet/mercado-libre-test/src/schema"
	"github.com/edwintrumpet/mercado-libre-test/src/services"
	"github.com/gin-gonic/gin"
)

// Set function that define all routes
func Set(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "X-men evaluator works!"})
	})

	router.POST("/mutant", func(c *gin.Context) {
		var dna schema.Dna
		error := c.ShouldBind(&dna)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": error.Error()})
			return
		}
		response := services.MutantService(dna.Dna)
		c.JSON(int(response.Status), gin.H{"message": response.Message})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "This route does not exist"})
	})
}
