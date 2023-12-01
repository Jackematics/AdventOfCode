package input_reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	result, _ := ReadInput("test_file.txt")

	expectedArrayLength := 9

	assert.Equal(t, len(result), expectedArrayLength)
	assert.Equal(t, result[0], "This")
	assert.Equal(t, result[1], "Is")
	assert.Equal(t, result[2], "A")
	assert.Equal(t, result[3], "Test")
}
