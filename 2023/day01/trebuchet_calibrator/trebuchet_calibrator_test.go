package trebuchet_calibrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneSumCalibrationValues(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := 142

	assert.Equal(t, expected, SumCalibrationValuesPartOne(input))
}

var part_two_test_cases = []struct {
	input  []string
	result int
}{
	{[]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}, 281},
	{[]string{"sevenine"}, 79},
	{[]string{"eighthree"}, 83},
	{[]string{"9fgsixzkbscvbxdsfive6spjfhzxbzvgbvrthreeoneightn"}, 98},
	{[]string{"two1two"}, 22},
	{[]string{"twoneight"}, 28},
}

func TestPartTwoSumCalibrationValues(t *testing.T) {
	for _, test_case := range part_two_test_cases {
		assert.Equal(t, test_case.result, SumCalibrationValuesPartTwo(test_case.input))
	}
}
