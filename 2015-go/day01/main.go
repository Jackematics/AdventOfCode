package main

import (
	"aoc_2015/day01/floor_finder"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := floor_finder.FindFloor(raw_input[0])
	part_two_result, _ := floor_finder.FindBasementEntryPosition(raw_input[0])

	fmt.Println("Part 1 solution: ", part_one_result)
	fmt.Println("Part 2 solution: ", part_two_result)
}
