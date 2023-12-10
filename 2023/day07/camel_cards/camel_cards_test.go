package camel_cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input = []string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}

func TestTotalWinningsPartOne(t *testing.T) {
	expected := 6440
	assert.Equal(t, expected, TotalWinnings(test_input, "part_one"))
}

func TestTotalWinningsPartTwo(t *testing.T) {
	expected := 5905
	assert.Equal(t, expected, TotalWinnings(test_input, "part_two"))
}
