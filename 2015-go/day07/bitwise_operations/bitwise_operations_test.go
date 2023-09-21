package bitwise_operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteBitwiseInstructions(t *testing.T) {
	input := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	}
	expected := map[string]interface{}{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}

	actual := ExecuteBitwiseInstructions(input)

	for key, expected_value := range expected {
		actual_value := actual[key]

		assert.Equal(t, expected_value, actual_value)
	}
}
