package nice_string_validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneCountNiceStrings(t *testing.T) {
	input := []string{
		"ugknbfddgicrmopn",
		"aaa",
		"jchzalrnumimnmhp",
		"haegwjzuvuyypxyu",
		"dvszwmarrgswjxmb",
	}
	expected := 2
	assert.Equal(t, expected, PartOneCountNiceStrings(input))
}

func TestPartTwoCountNiceStrings(t *testing.T) {
	input := []string{
		"qjhvhtzxzqqjkmpb",
		"xxyxx",
		"jchzalrnumimnmhp",
		"ieodomkazucvgmuy",
	}
	expected := 2
	assert.Equal(t, expected, PartTwoCountNiceStrings(input))
}
