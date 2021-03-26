package services

import (
	"fmt"

	"github.com/edwintrumpet/mercado-libre-test/src/database"
	"github.com/edwintrumpet/mercado-libre-test/src/schema"
)

func Stats() schema.Stats {
	var stats schema.Stats
	response, err := database.StatsGet()
	if err != nil {
		fmt.Println(err)
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

	return stats
}
