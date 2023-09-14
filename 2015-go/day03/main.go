package main

import (
	"aoc_2015/day03/house_traverser"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	result := house_traverser.TraverseHouses(raw_input[0])
	// part_two_result := present_wrapping_calculator.CalculateAllRibbonNeeded(raw_input)

	fmt.Println("Solution: ", result)
	// fmt.Println("Part 2 solution: ", part_two_result)
}
