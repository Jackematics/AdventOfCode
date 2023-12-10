package main

import (
	"aoc_2015/day08/haunted_wasteland"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := haunted_wasteland.TotalSteps(raw_input)

	fmt.Println("Part 1 solution", part_one_result)
}
