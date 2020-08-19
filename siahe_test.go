package siahe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAndFind(t *testing.T) {
	primeEvilsList := New()

	// Fill Siahe with some random data
	primeEvilsList.Index("key-1", "Diablo")
	primeEvilsList.Index("key-1", "Lord")
	primeEvilsList.Index("key-1", "Terror")

	primeEvilsList.Index("key-2", "Mephisto")
	primeEvilsList.Index("key-2", "Lord")
	primeEvilsList.Index("key-2", "Hatred")
	primeEvilsList.Index("key-2", "Spirit")
	primeEvilsList.Index("key-2", "Love")

	primeEvilsList.Index("key-3", "Baal")
	primeEvilsList.Index("key-3", "Lord")
	primeEvilsList.Index("key-3", "Destruction")

	primeEvilsList.Index("key-1", "Diavlo")

	expectedResults := map[string][]string{
		"L":      {"key-1", "key-2", "key-3"},
		"Lo":     {"key-1", "key-2", "key-3"},
		"Lov":    {"key-2"},
		"Love":   {"key-2"},
		"B":      {"key-3"},
		"Loved":  {},
		"Lord":   {"key-1", "key-2", "key-3"},
		"D":      {"key-1", "key-3"},
		"Diablo": {"key-1"},
		"Des":    {"key-3"},
		"Du":     {},
		"Duriel": {},
		"":       {"key-1", "key-2", "key-3"},
	}

	for query, expectedResult := range expectedResults {
		result := primeEvilsList.Find(query)

		assert.Equal(t, len(expectedResult), len(result), "Find should return right number of results when searching for '%s'", query)
		assert.Equal(t, expectedResult, result, "Find should return right pairs when searching for '%s'", query)
	}
}
