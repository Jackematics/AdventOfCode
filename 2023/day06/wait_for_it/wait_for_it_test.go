package wait_for_it

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestMultiplyWinningStrategyCount(t *testing.T) {
	assert.Equal(t, 288, MultiplyWinningStrategyCount(test_input))
}

func TestTotalWinningStrategyCount(t *testing.T) {
	assert.Equal(t, 71503, TotalWinningStrategyCount(test_input))
}
