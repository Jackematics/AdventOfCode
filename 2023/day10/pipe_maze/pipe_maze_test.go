package pipe_maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_cases = []struct {
	input    []string
	expected int
}{
	{[]string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}, 4},
	// {[]string{
	// 	"7-F7-",
	// 	".FJ|7",
	// 	"SJLL7",
	// 	"|F--J",
	// 	"LJ.LJ",
	// }, 8},
}

func TestFindFarthestDistanceFromLoopStart(t *testing.T) {
	for _, test_case := range test_cases {
		assert.Equal(t, test_case.expected, FindFarthestDistanceFromLoopStart(test_case.input))
	}
}
