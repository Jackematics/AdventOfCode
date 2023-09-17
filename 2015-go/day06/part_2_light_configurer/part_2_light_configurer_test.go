package part_2_light_configurer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalBrightness(t *testing.T) {
	input := []string{
		"turn on 0,0 through 0,0",
		"toggle 0,0 through 999,999",
	}
	expected := 2000001
	assert.Equal(t, expected, CalculateTotalBrightness(input))
}
