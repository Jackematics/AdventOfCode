package floor_finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var find_floor_test_cases = []struct {
	input  string
	result int
}{
	{"(())", 0},
	{"()()", 0},
	{"(((", 3},
	{"(()(()(", 3},
	{"))(((((", 3},
	{"())", -1},
	{"))(", -1},
	{")))", -3},
	{")())())", -3},
}

func TestFindFloor(t *testing.T) {
	for _, test_case := range find_floor_test_cases {
		assert.Equal(t, test_case.result, FindFloor(test_case.input))
	}
}

var basement_entry_test_cases = []struct {
	input  string
	result int
}{
	{")", 1},
	{"()())", 5},
}

func TestFindBasementEntryPosition(t *testing.T) {
	for _, test_case := range basement_entry_test_cases {
		pos, _ := FindBasementEntryPosition(test_case.input)
		assert.Equal(t, test_case.result, pos)
	}
}
