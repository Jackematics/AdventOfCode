package main

import (
	"aoc_2023/day04/scratchcards"
	"aoc_2023/input_reader"
	"fmt"
)

func main() {
	raw_numbers, _ := input_reader.ReadInput("input.txt")

	part_one_result := scratchcards.SumScratchardsPoints(raw_numbers)
	part_two_result := scratchcards.FindTotalScratchcards(raw_numbers)

	fmt.Println("Part one result: ", part_one_result)
	fmt.Println("Part two result: ", part_two_result)
}
