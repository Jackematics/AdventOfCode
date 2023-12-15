package mirage_maintenance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_cases = []struct {
	input             []string
	part_one_expected int
	part_two_expected int
}{
	{[]string{"0 3 6 9 12 15"}, 18, -3},
	{[]string{"1 3 6 10 15 21"}, 28, 0},
	{[]string{"10 13 16 21 30 45"}, 68, 5},
	{[]string{"0 3 6 9 12 15", "1 3 6 10 15 21", "10 13 16 21 30 45"}, 114, 2},
}

func TestSumNextHistoryValues(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.part_one_expected, SumNextHistoryValues(test_case.input))
	}
}

func TestSumPrevHistoryValues(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.part_two_expected, SumPrevHistoryValues(test_case.input))
	}
}
