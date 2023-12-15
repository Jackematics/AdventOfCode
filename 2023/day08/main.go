package main

import (
	"aoc_2023/day08/haunted_wasteland"
	"aoc_2023/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := haunted_wasteland.TotalSteps(raw_input)
	part_two_result := haunted_wasteland.TotalSimultaneousSteps(raw_input)

	fmt.Println("Part 1 solution", part_one_result)
	fmt.Println("Part 2 solution", part_two_result)
}
