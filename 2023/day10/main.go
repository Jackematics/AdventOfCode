package main

import (
	"aoc_2023/day10/pipe_maze"
	"aoc_2023/input_reader"
	"fmt"
)

func main() {
	raw_input, _ := input_reader.ReadInput("input.txt")

	part_one_result := pipe_maze.FindFarthestDistanceFromLoopStart(raw_input)

	fmt.Println("Part one solution", part_one_result)
}
