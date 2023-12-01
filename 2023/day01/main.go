package main

import (
	"aoc_2015/day01/trebuchet_calibrator"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := trebuchet_calibrator.SumCalibrationValues(raw_input)

	fmt.Println("Part 1 solution", part_one_result)
}
