package schema

// Stats is the strcuture for requests historial
type Stats struct {
	CountMutantDna float32 `json:"count_mutant_dna"`
	CountHumanDna  float32 `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}
