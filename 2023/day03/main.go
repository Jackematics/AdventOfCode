package main

import (
	"aoc_2015/day03/gear_ratios"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := gear_ratios.SumPartNumbers(raw_input)
	part_two_result := gear_ratios.SumGearRatios(raw_input)

	fmt.Println("Part One Result: ", part_one_result)
	fmt.Println("Part Two Result: ", part_two_result)
}
