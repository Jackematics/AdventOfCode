package main

import (
	"aoc_2024/input_reader"
	"fmt"
	"strconv"
)

type GridIndex struct {
	Row int
	Col int
}

func outOfBounds(input []string, centreRow int, centreCol int, rowDelta int, colDelta int) bool {
	return centreRow + 3*rowDelta < 0 || 
	centreRow + 3*rowDelta >= len(input) ||
	centreCol + 3*colDelta < 0 ||
	centreCol + 3*colDelta >= len(input[0])
}

func countXmasDirection(input []string, centreRow int, centreCol int, rowDelta int, colDelta int) int {
	if outOfBounds(input, centreRow, centreCol, rowDelta, colDelta) {
		return 0
	}

	if input[centreRow + rowDelta][centreCol + colDelta] == 'M' &&
	input[centreRow + 2*rowDelta][centreCol + 2*colDelta] == 'A' &&
	input[centreRow + 3*rowDelta][centreCol + 3*colDelta] == 'S' {
		return 1
	}

	return 0
}

func countXmasPartOne(input []string, centreRow int, centreCol int) int {
	xmasCount := 0

	deltas := []GridIndex{{-1, -1}, { -1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	for _, delta := range deltas {
		xmasCount += countXmasDirection(input, centreRow, centreCol, delta.Row, delta.Col)
	}

	return xmasCount
}


func partOne(input []string) int {
	xmasCount := 0
	for row := range input {
		for col := range input[row] {
			if input[row][col] == 'X' {
				xmasCount += countXmasPartOne(input, row, col)
			}
		}
	}

	return xmasCount 
}

func countDiagonalMAndS(input []string, diagonal []GridIndex) (int, int) {
	mCount := 0
	sCount := 0
	for _, gridIndex := range diagonal {
		if input[gridIndex.Row][gridIndex.Col] == 'M' {
			mCount++
		}

		if input[gridIndex.Row][gridIndex.Col] == 'S' {
			sCount++
		}
	}

	return mCount, sCount
}

func countXmasPartTwo(input []string, centreRow int, centreCol int) int {
	xmasCount := 0

	diagOne := []GridIndex{{Row: centreRow - 1, Col: centreCol - 1}, {Row: centreRow + 1, Col: centreCol + 1}}
	diagTwo := []GridIndex{{Row: centreRow - 1, Col: centreCol + 1}, {Row: centreRow + 1, Col: centreCol - 1}}

	mCountOne, sCountOne := countDiagonalMAndS(input, diagOne)
	mCountTwo, sCountTwo := countDiagonalMAndS(input, diagTwo)

	if mCountOne == 1 && sCountOne == 1 && mCountTwo == 1 && sCountTwo == 1 {
		xmasCount++
	}
	

	return xmasCount
}

func partTwo(input []string) int {
	xmasCount := 0
	for row := 1; row < len(input) - 1; row++ {
		for col := 1; col < len(input[0]) - 1; col++ {
			if input[row][col] == 'A' {
				xmasCount += countXmasPartTwo(input, row, col)
			}
		}
	}

	return xmasCount
}


func main() {
	exampleOne := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	examplePartOne := partOne(exampleOne)
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := partTwo(exampleOne)
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input)
	partTwoResult := partTwo(input)

	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
