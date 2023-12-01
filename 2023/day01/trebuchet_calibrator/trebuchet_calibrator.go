package trebuchet_calibrator

import (
	"slices"
	"strconv"
	"strings"
)

func SumCalibrationValuesPartOne(raw_calibration_values []string) int {
	sum := 0
	for _, raw_calibration_value := range raw_calibration_values {
		sum += calculateCalibrationValues(raw_calibration_value)
	}

	return sum
}

func SumCalibrationValuesPartTwo(raw_calibration_values []string) int {
	sum := 0
	for _, raw_calibration_value := range raw_calibration_values {
		deverbosified := deverbosifyCalibrationValue(raw_calibration_value)
		sum += calculateCalibrationValues(deverbosified)
	}

	return sum
}

var integers = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func calculateCalibrationValues(raw_calibration_value string) int {
	firstInteger := "-1"
	lastInteger := "-1"

	for i := range raw_calibration_value {
		if slices.Contains(integers, raw_calibration_value[i]) {
			firstInteger = string(raw_calibration_value[i])
			break
		}
	}

	for i := len(raw_calibration_value) - 1; i >= 0; i-- {
		if slices.Contains(integers, raw_calibration_value[i]) {
			lastInteger = string(raw_calibration_value[i])
			break
		}
	}

	concatenated, _ := strconv.Atoi(firstInteger + lastInteger)
	return concatenated
}

var verbose_integers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var verbose_integer_mappings = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func deverbosifyCalibrationValue(raw_calibration_value string) string {
	deverbosified := raw_calibration_value
	for verboseIntegersInString(deverbosified) {
		for _, integer := range verbose_integers {
			index := strings.Index(deverbosified, integer)

			if index == -1 {
				continue
			}

			deverbosified = deverbosified[:index+1] + verbose_integer_mappings[integer] + deverbosified[index+1:]
		}
	}

	return deverbosified
}

func verboseIntegersInString(input string) bool {
	for _, integer := range verbose_integers {
		if strings.Contains(input, integer) {
			return true
		}
	}

	return false
}
