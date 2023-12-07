package main

import (
	"aoc_2015/day05/seed_mapper"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := seed_mapper.FindLowestLocationNumber(raw_input, "one")
	part_two_result := seed_mapper.FindLowestLocationNumber(raw_input, "two")

	fmt.Println("Part one result: ", part_one_result)
	fmt.Println("Part two result: ", part_two_result)
}
