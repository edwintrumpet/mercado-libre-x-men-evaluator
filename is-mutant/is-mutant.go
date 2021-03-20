package isMutant

// IsMutant will receive an array of strings with DNA secuences
// It will return true if DNA secuences represent a mutant and false if not
func IsMutant(dna []string) bool {
	for i, secuence := range dna {
		for j, v := range secuence {
			// horizontal review
			if j < len(secuence)-3 &&
				string(v) == string(secuence[j+1]) &&
				string(v) == string(secuence[j+2]) &&
				string(v) == string(secuence[j+3]) {
				return true
			}

			// vertical review
			if i < len(dna)-3 &&
				string(v) == string(dna[i+1][j]) &&
				string(v) == string(dna[i+2][j]) &&
				string(v) == string(dna[i+3][j]) {
				return true
			}

			// diagonal review
			if j < len(secuence)-3 &&
				i < len(dna)-3 &&
				string(v) == string(dna[i+1][j+1]) &&
				string(v) == string(dna[i+2][j+2]) &&
				string(v) == string(dna[i+3][j+3]) {
				return true
			}

			// diagonal review inverse optional
			// if j < len(secuence)-3 &&
			// 	i >= 3 &&
			// 	string(v) == string(dna[i-1][j+1]) &&
			// 	string(v) == string(dna[i-2][j+2]) &&
			// 	string(v) == string(dna[i-3][j+3]) {
			// 	return true
			// }
		}
	}
	return false
}
