package ltools

import (
	"strconv"
	"testing"
)

type BasicStringer struct {
	value int
}

func (bs BasicStringer) String() string {
	return strconv.Itoa(bs.value)
}

func TestMapKeys(t *testing.T) {
	{
		testA := map[string]struct{}{}
		testA["one"] = struct{}{}
		testA["two"] = struct{}{}

		keys := MapKeyStrings(testA)
		if len(keys) != 2 {
			t.Error("incorrect number of keys reported")
		}
		if !((keys[0] == "one" && keys[1] == "two") ||
			(keys[0] == "two" && keys[1] == "one")) {
			t.Error("incorrect values reported")
		}
	}

	{
		testB := map[BasicStringer]struct{}{}
		testB[BasicStringer{1}] = struct{}{}
		testB[BasicStringer{2}] = struct{}{}

		keys := MapKeyStrings(testB)
		if len(keys) != 2 {
			t.Error("incorrect number of keys reported")
		}
		if !((keys[0] == "1" && keys[1] == "2") ||
			(keys[0] == "2" && keys[1] == "1")) {
			t.Error("incorrect values reported")
		}
	}
}
