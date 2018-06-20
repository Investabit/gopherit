package gostrings

import (
	"strings"

	"github.com/Investabit/gopherit/ltools"
)

type FuncTypeStringReplacer func(in string) string

func ReplaceBlindly(replacement string) FuncTypeStringReplacer {
	return func(in string) string {
		return replacement
	}
}

func ReplaceEmpty(in string) string {
	return ""
}

// FilterCharacters reports a string with the specified set of characters
// replaced using the specified replacer function. The inclusive  parameter
// specifies whether the specified characters are replaced (true), or
// all characters excluding the specified characters are replaced (false).
func FilterCharacters(
	input string,
	list string, replacer FuncTypeStringReplacer, inclusive bool,
) string {
	newString := ""
	for _, c := range input {
		if ltools.Xor(inclusive, strings.Contains(list, string(c))) {
			newString += string(c)
		} else {
			newString += replacer(string(c))
		}
	}
	return newString
}

// FilterCharactersExcluding reports a string with non-whitelisted characters
// replaced using the specified replacer function. To remove all non-whitelisted
// characters, use gostrings.ReplaceEmpty as the second parameter.
func FilterCharactersExcluding(
	input string,
	list string, replacer FuncTypeStringReplacer,
) string {
	return FilterCharacters(input, list, replacer, false)
}

// FilterCharactersIncluding reports a string with blacklisted characters
// replaced using the specified replacer function. To remove all non-whitelisted
// characters, use gostrings.ReplaceEmpty as the second parameter.
func FilterCharactersIncluding(
	input string,
	list string, replacer FuncTypeStringReplacer,
) string {
	return FilterCharacters(input, list, replacer, true)
}
