package trebuchet_calibrator

import (
	"slices"
	"strconv"
)

var integers = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func SumCalibrationValues(raw_calibration_values []string) int {
	sum := 0
	for _, raw_calibration_value := range raw_calibration_values {
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
		sum += concatenated
	}

	return sum
}
