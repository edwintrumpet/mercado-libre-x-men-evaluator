package main

import (
	"github.com/edwintrumpet/mercado-libre-test/src/functions/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// noMutantDNA := []string{
	// 	"ATGCGA",
	// 	"CAGTGC",
	// 	"TTATTT",
	// 	"AGACGG",
	// 	"GCGTCA",
	// 	"TCACTG",
	// }
	// mutantDNA := []string{
	// 	"ATGCGA",
	// 	"CAGTGC",
	// 	"TTATGT",
	// 	"AGAAGG",
	// 	"CCCCTA",
	// 	"TCACTG",
	// }

	router := gin.Default()

	routes.Set(router)

	router.Run()
}
