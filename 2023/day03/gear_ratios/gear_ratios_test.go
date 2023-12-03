package gear_ratios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPartNumbers(t *testing.T) {
	input := [][]string{
		{"467..114.."},
		{"...*......"},
		{"..35..633."},
		{"......#..."},
		{"617*......"},
		{".....+.58."},
		{"..592....."},
		{"......755."},
		{"...$.*...."},
		{".664.598.."},
	}

	expected := 4361
	assert.Equal(t, expected, SumPartNumbers(input))
}
