package main

import (
	"aoc_2015/day07/bitwise_operations"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := bitwise_operations.ExecuteBitwiseInstructions(raw_input)

	fmt.Println("Part 1 solution: ", part_one_result)
}
