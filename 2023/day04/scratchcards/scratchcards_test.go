package scratchcards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var whole_set = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

var test_cases = []struct {
	input           []string
	part_one_result int
}{
	{[]string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, 8},
	{[]string{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"}, 2},
	{[]string{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"}, 2},
	{[]string{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"}, 1},
	{[]string{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"}, 0},
	{[]string{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}, 0},
	{whole_set, 13},
}

func TestSumScratchardsPoints(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.part_one_result, SumScratchardsPoints(test_case.input))
	}
}

func TestFindTotalScratchcards(t *testing.T) {
	assert.Equal(t, 30, FindTotalScratchcards(whole_set))
}
