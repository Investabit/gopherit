package gostrings_test

import (
	"testing"

	"github.com/Investabit/gopherit/gostrings"
	"github.com/stretchr/testify/assert"
)

func TestAtomicReplace(t *testing.T) {
	type Case struct {
		// Independant variables
		input        string
		replacements map[string]string
		// Dependant variables
		count  map[string]int
		output string
	}
	cases := []Case{
		// Case with longer replacement
		Case{
			input:  "fuzzy wuzzy wuzza bear",
			output: "fuzzy becausey becausea bear",
			count: map[string]int{
				"wuzz": 2,
			},
			replacements: map[string]string{
				"wuzz": "because",
			},
		},
		// Case with shorter replacement
		Case{
			input:  "fuzzy wuzzy wuzza bear",
			output: "fuzzy wuzy wuza bear",
			count: map[string]int{
				"wuzz": 2,
			},
			replacements: map[string]string{
				"wuzz": "wuz",
			},
		},
	}

	for _, cse := range cases {
		count, output := stringutil.AtomicReplace(cse.input, cse.replacements)
		assert.EqualValues(t, cse.output, output, "output must match")
		assert.EqualValues(t, cse.count, count, "count must match")
	}
}
