package services

import "github.com/edwintrumpet/mercado-libre-test/src/schema"

func Stats() schema.Stats {
	var stats schema.Stats
	stats.CountMutantDna = 40
	stats.CountHumanDna = 100
	stats.Ratio = stats.CountMutantDna / stats.CountHumanDna
	return stats
}
