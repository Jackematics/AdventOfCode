package gear_ratios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestSumPartNumbers(t *testing.T) {
	expected := 4361
	assert.Equal(t, expected, SumPartNumbers(input))
}

func TestSumGearRatios(t *testing.T) {
	expected := 467835
	assert.Equal(t, expected, SumGearRatios(input))
}
