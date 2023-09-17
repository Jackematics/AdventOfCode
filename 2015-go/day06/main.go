package main

import (
	"aoc_2015/day06/part_1_light_configurer"
	"aoc_2015/day06/part_2_light_configurer"
	"aoc_2015/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := part_1_light_configurer.CountLightsOn(raw_input)
	part_two_result := part_2_light_configurer.CalculateTotalBrightness(raw_input)

	fmt.Println("Part 1 solution: ", part_one_result)
	fmt.Println("Part 2 solution: ", part_two_result)
}
