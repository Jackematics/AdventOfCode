package main

import (
	"aoc_2015/day04/scratchcards"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_numbers, _ := input_reader.ReadInput("input.txt")

	part_one_result := scratchcards.SumScratchardsPoints(raw_numbers)

	fmt.Println("Part one result: ", part_one_result)
}
