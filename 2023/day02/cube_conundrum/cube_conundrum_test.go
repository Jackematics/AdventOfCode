package cube_conundrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_cases = []struct {
	input           []string
	part_one_result int
	part_two_result int
}{
	{[]string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, 1, 48},
	{[]string{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, 2, 12},
	{[]string{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, 0, 1560},
	{[]string{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, 0, 630},
	{[]string{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, 5, 36},
	{[]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}, 8, 2286},
}

func TestSumValidIds(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.part_one_result, SumValidIds(test_case.input))
	}
}

func TestSumSetPower(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.part_two_result, SumSetPower(test_case.input))
	}
}
