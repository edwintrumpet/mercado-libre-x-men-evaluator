package schema

// Dna is the required strcuture for dna secuences given by
// the post request
type Dna struct {
	Dna []string `binding:"required"`
}
