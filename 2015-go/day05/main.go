package main

import (
	"aoc_2015/day05/nice_string_validator"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := nice_string_validator.PartOneCountNiceStrings(raw_input)
	part_two_result := nice_string_validator.PartTwoCountNiceStrings(raw_input)

	fmt.Println("Part 2 solution: ", part_one_result)
	fmt.Println("Part 2 solution: ", part_two_result)
}
