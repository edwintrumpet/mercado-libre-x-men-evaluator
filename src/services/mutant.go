package services

import (
	"net/http"

	isMutant "github.com/edwintrumpet/mercado-libre-test/src/functions/is_mutant"
	"github.com/edwintrumpet/mercado-libre-test/src/schema"
)

// MutantService receives a dna secuence and returns an ok message if
// dna belongs to a mutant and false otherwise
func Mutant(dna []string) schema.Response {
	var response schema.Response

	isMutant := isMutant.IsMutant(dna)
	if isMutant {
		response.Status = http.StatusOK
		response.Message = "OK"
	} else {
		response.Status = http.StatusForbidden
		response.Message = "Forbidden"
	}
	return response
}
