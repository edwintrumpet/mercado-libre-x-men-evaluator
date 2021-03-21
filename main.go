package main

import (
	"fmt"

	isMutant "github.com/edwintrumpet/mercado-libre-test/is_mutant"
)

func main() {
	noMutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}
	mutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	noMutantResult := isMutant.IsMutant(noMutantDNA)
	mutantResult := isMutant.IsMutant(mutantDNA)

	fmt.Println("noMutantResult", noMutantResult)
	fmt.Println("mutantResult", mutantResult)
}
