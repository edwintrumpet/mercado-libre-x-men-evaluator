package main

import (
	"github.com/edwintrumpet/mercado-libre-test/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Set(router)

	router.Run(":80")
}
