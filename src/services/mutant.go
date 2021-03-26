package services

import (
	"net/http"

	"github.com/edwintrumpet/mercado-libre-test/src/database"
	isMutant "github.com/edwintrumpet/mercado-libre-test/src/functions/is_mutant"
	"github.com/edwintrumpet/mercado-libre-test/src/schema"
)

// MutantService receives a dna secuence and returns an ok message if
// dna belongs to a mutant and false otherwise
func Mutant(dna []string) (schema.Response, error) {
	var response schema.Response
	var t string
	isMutant := isMutant.IsMutant(dna)
	if isMutant {
		response.Status = http.StatusOK
		response.Message = "OK"
		t = "count_mutant_dna"
	} else {
		response.Status = http.StatusForbidden
		response.Message = "Forbidden"
		t = "count_human_dna"
	}
	err := database.IncrementCount(t)
	return response, err
}
