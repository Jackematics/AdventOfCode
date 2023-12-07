package main

import (
	"aoc_2015/day06/wait_for_it"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := wait_for_it.MultiplyWinningStrategyCount(raw_input)
	part_two_result := wait_for_it.TotalWinningStrategyCount(raw_input)

	fmt.Println("Part 1 solution", part_one_result)
	fmt.Println("Part 2 solution", part_two_result)
}
