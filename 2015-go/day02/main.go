package main

import (
	"aoc_2015/day02/present_wrapping_calculator"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := present_wrapping_calculator.CalculateAllPaperNeeded(raw_input)
	part_two_result := present_wrapping_calculator.CalculateAllRibbonNeeded(raw_input)

	fmt.Println("Part 1 solution: ", part_one_result)
	fmt.Println("Part 2 solution: ", part_two_result)
}
