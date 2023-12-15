package main

import (
	"aoc_2023/day09/mirage_maintenance"
	"aoc_2023/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := mirage_maintenance.SumNextHistoryValues(raw_input)
	part_two_result := mirage_maintenance.SumPrevHistoryValues(raw_input)

	fmt.Println("Part one solution", part_one_result)
	fmt.Println("Part two solution", part_two_result)
}
