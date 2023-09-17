package part_1_light_configurer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLightsOn(t *testing.T) {
	input := []string{
		"turn on 0,0 through 999,999",
		"toggle 0,0 through 999,0",
		"turn off 499,499 through 500,500",
	}
	expected := 998996
	assert.Equal(t, expected, CountLightsOn(input))
}
