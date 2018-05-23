package gostrings_test

import (
	"testing"

	"github.com/Investabit/investabit-go/stringutil"
	"github.com/stretchr/testify/assert"
)

func TestIndexAll(t *testing.T) {
	type Case struct {
		input   string
		substr  string
		indices []int
	}
	cases := []Case{
		// Case with one instance of a substring
		Case{
			input:   "fuzzy wuzzy wuzza bear",
			substr:  "wuzzy",
			indices: []int{6},
		},
		// Case with two instances of a substring
		Case{
			input:   "fuzzy wuzzy wuzza bear",
			substr:  "wuzz",
			indices: []int{6, 12},
		},
		// Case with overlapping instances of a substring
		Case{
			input:   "fuzzy aabbaabbaa bear",
			substr:  "aabbaa",
			indices: []int{6},
		},
		// Case to make sure index isn't incremented too far
		Case{
			input:   "fuzzy aabbaabbaaaabbaa bear",
			substr:  "aabbaa",
			indices: []int{6, 16},
		},
	}

	for _, cse := range cases {
		result := stringutil.IndexAll(cse.input, cse.substr)
		assert.EqualValues(t, cse.indices, result, "indices must match")
	}
}
