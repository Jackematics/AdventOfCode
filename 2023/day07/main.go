package main

import (
	"aoc_2023/day07/camel_cards"
	"aoc_2023/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := camel_cards.TotalWinnings(raw_input, "part_one")
	part_two_result := camel_cards.TotalWinnings(raw_input, "part_two")

	fmt.Println("Part 1 solution", part_one_result)
	fmt.Println("Part 2 solution", part_two_result)
}
