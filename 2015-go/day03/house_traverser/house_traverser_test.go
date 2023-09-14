package house_traverser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var traverse_houses_test_cases = []struct {
	input  string
	result int
}{
	{"^v", 3},
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
}

func TestTraverseHouses(t *testing.T) {
	for _, test_case := range traverse_houses_test_cases {
		assert.Equal(t, test_case.result, TraverseHouses(test_case.input))
	}
}
