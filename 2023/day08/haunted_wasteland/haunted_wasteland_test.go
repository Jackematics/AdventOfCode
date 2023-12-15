package haunted_wasteland

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_cases = []struct {
	input    []string
	expected int
}{
	{[]string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}, 2},
	{[]string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}, 6},
}

func TestTotalSteps(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.expected, TotalSteps(test_case.input))
	}
}

func TestTotalSimultaneousSteps(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	expected := 6

	assert.Equal(t, expected, TotalSimultaneousSteps(input))
}
