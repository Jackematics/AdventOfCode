package main

import (
	"aoc_2024/input_reader"
	"fmt"
	"strconv"
)

type Cell struct {
	Frequency rune
	Antinode bool
}

type Grid [][]Cell

type GridIndex struct {
	Row int
	Col int
}

func (grid Grid) inBounds(index GridIndex) bool {
	return index.Row >= 0 &&
	index.Row < len(grid) &&
	index.Col >= 0 &&
	index.Col < len(grid[0])
}

func (grid Grid) setAntinode(fixed GridIndex, resonant GridIndex) {
	rowDiff := fixed.Row - resonant.Row
	colDiff := fixed.Col - resonant.Col

	antinodeIndex := GridIndex{ Row: resonant.Row - rowDiff, Col: resonant.Col - colDiff }

	if grid.inBounds(antinodeIndex) {
		grid[antinodeIndex.Row][antinodeIndex.Col].Antinode = true
	}
}

func (grid Grid) setResonantAntinodes(fixed GridIndex, resonant GridIndex) {
	rowDiff := fixed.Row - resonant.Row
	colDiff := fixed.Col - resonant.Col
	antinodeIndex := GridIndex{Row: resonant.Row, Col: resonant.Col}

	for i := 1; grid.inBounds(antinodeIndex); i++ {
		grid[antinodeIndex.Row][antinodeIndex.Col].Antinode = true
		antinodeIndex.Row -= rowDiff
		antinodeIndex.Col -= colDiff
	}
; 
}

func (grid Grid) setAntinodes(part string, frequency rune, fixedIndex GridIndex) {
	for row := range grid {
		for col := range grid[row] {
			if row == fixedIndex.Row && col == fixedIndex.Col {
				continue
			}

			if grid[row][col].Frequency == frequency {
				switch part {
				case "one":
					grid.setAntinode(fixedIndex, GridIndex{Row: row, Col: col})
					break
				case "two":
					grid.setResonantAntinodes(fixedIndex, GridIndex{Row: row, Col: col})
					break
				default:
					panic("Invalid part")
				}
			}
		}
	}
}

func (grid Grid) setAllAntinodes(part string) {
	for row := range grid {
		for col := range grid[row] {
			currentFrequency := grid[row][col].Frequency
			if currentFrequency != '.' {
				grid.setAntinodes(part, currentFrequency, GridIndex{ Row: row, Col: col })
			}
		}
	}
}

func solution(input []string, part string) int {
	grid := Grid{}
	for row := range input {
		gridRow := []Cell{} 
		for _, val := range input[row] {
			gridRow = append(gridRow, Cell{ Frequency: val, Antinode: false })
		}
		grid = append(grid, gridRow)
	}

	grid.setAllAntinodes(part)

	antinodeSum := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col].Antinode {
				antinodeSum++
			}
		}
	}

	return antinodeSum
}


func main() {
	exampleOne := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	examplePartOne := solution(exampleOne, "one")
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := solution(exampleOne, "two")
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := solution(input, "one")
	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))

	partTwoResult := solution(input, "two")
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
