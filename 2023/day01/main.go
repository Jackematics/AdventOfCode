package main

import (
	"aoc_2023/day01/trebuchet_calibrator"
	"aoc_2023/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := trebuchet_calibrator.SumCalibrationValuesPartOne(raw_input)
	part_two_result := trebuchet_calibrator.SumCalibrationValuesPartTwo(raw_input)

	fmt.Println("Part 1 solution", part_one_result)
	fmt.Println("Part 2 solution", part_two_result)
}
