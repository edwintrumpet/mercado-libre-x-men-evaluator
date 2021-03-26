package isMutant_test

import (
	"testing"

	isMutant "github.com/edwintrumpet/mercado-libre-test/src/functions/is_mutant"
)

func TestIsMutant(t *testing.T) {
	// No mutant test
	dnaSecuence := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}

	got := isMutant.IsMutant(dnaSecuence)
	if got != false {
		t.Errorf("Response should be false but receives %t", got)
		t.Fail()
	} else {
		t.Log("No mutant DNA passed")
	}

	// Mutant test with horizontal match
	dnaSecuence = []string{
		"ATGCGA",
		"CAGTGC",
		"TTTTTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}

	got = isMutant.IsMutant(dnaSecuence)
	if got != true {
		t.Errorf("Response should be true but receives %t", got)
		t.Fail()
	} else {
		t.Log("Mutant DNA with horizontal match passed")
	}

	// Mutant test with vertical match
	dnaSecuence = []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCATCA",
		"TCACTG",
	}

	got = isMutant.IsMutant(dnaSecuence)
	if got != true {
		t.Errorf("Response should be true but receives %t", got)
		t.Fail()
	} else {
		t.Log("Mutant DNA with vertical match passed")
	}

	// Mutant test with diagonal match
	dnaSecuence = []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGAAGG",
		"GCGTCA",
		"TCACTG",
	}

	got = isMutant.IsMutant(dnaSecuence)
	if got != true {
		t.Errorf("Response should be true but receives %t", got)
		t.Fail()
	} else {
		t.Log("Mutant DNA with diagonal match passed")
	}

	// Mutant test with multiple matches
	dnaSecuence = []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	got = isMutant.IsMutant(dnaSecuence)
	if got != true {
		t.Errorf("Response should be true but receives %t", got)
		t.Fail()
	} else {
		t.Log("Mutant DNA with multiple matches passed")
	}
}
