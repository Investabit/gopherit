package gostrings

import (
	"testing"
)

func TestFilter(t *testing.T) {
	{
		result := FilterCharactersExcluding(
			"this-is-a-test", "abcdef", ReplaceEmpty,
		)

		if result != "ae" {
			t.Fail()
		}
	}

	{
		result := FilterCharactersIncluding(
			"this-is-a-test", "abcdef", ReplaceEmpty,
		)

		if result != "this-is--tst" {
			t.Fail()
		}
	}
}
