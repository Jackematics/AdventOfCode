package trebuchet_calibrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumCalibrationValues(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := 142

	assert.Equal(t, expected, SumCalibrationValues(input))
}
