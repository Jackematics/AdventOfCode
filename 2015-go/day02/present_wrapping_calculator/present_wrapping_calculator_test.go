package present_wrapping_calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAllPaperNeeded(t *testing.T) {
	input := []string{"2x3x4", "1x1x10"}
	expected := 101
	assert.Equal(t, expected, CalculateAllPaperNeeded(input))
}

func TestCalculateAllRibbonNeeded(t *testing.T) {
	input := []string{"2x3x4", "1x1x10"}
	expected := 48

	assert.Equal(t, expected, CalculateAllRibbonNeeded(input))
}
