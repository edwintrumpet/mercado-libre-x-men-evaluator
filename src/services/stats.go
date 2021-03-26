package services

import (
	"github.com/edwintrumpet/mercado-libre-test/src/database"
	"github.com/edwintrumpet/mercado-libre-test/src/schema"
)

// Stats search in database the number of times that
// was provided an specific type of dna, it returns
// these values and a calculated ratio
func Stats() (schema.Stats, error) {
	var stats schema.Stats
	response, err := database.StatsGet()
	if err != nil {
		return stats, err
	}

	m := make(map[string]int)
	for _, v := range response {
		m[v.Name] = v.Value
	}
	stats.CountHumanDna = float32(m["count_human_dna"])
	stats.CountMutantDna = float32(m["count_mutant_dna"])
	if stats.CountHumanDna != 0 {
		stats.Ratio = stats.CountMutantDna / stats.CountHumanDna
	} else {
		stats.Ratio = stats.CountMutantDna
	}

	return stats, nil
}
