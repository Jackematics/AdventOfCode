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

type Guard struct {
	Position GridIndex
	Direction rune
}

func (guard Guard) indexInFront() GridIndex {
	if guard.Direction == '^' {
		return GridIndex{Row: guard.Position.Row - 1, Col: guard.Position.Col}
	}

	if guard.Direction == '>' {
		return GridIndex{Row: guard.Position.Row, Col: guard.Position.Col + 1}
	}

	if guard.Direction == 'v' {
		return GridIndex{Row: guard.Position.Row + 1, Col: guard.Position.Col}
	}

	if guard.Direction == '<' {
		return GridIndex{Row: guard.Position.Row , Col: guard.Position.Col - 1}
	}

	panic("Invalid guard direction")
}

func (guard *Guard) turnRight() {
	if guard.Direction == '^' {
		(*guard).Direction = '>'
		return
	}

	if guard.Direction == '>' {
		(*guard).Direction = 'v'
		return
	}

	if guard.Direction == 'v' {
		(*guard).Direction = '<'
		return
	}

	if guard.Direction == '<' {
		(*guard).Direction = '^'
		return
	}
}

type Grid []string

func (grid *Grid) initGuard() Guard {
	for row := range (*grid) {
		for col := range (*grid)[row] {
			if (*grid)[row][col] == '^' {
				position := GridIndex{Row: row, Col: col}
				return Guard{ Position: position, Direction: '^' }
			}
		}
	}

	panic("No guard found")
}

func (grid *Grid) inBounds(gridIndex GridIndex) bool {
	return gridIndex.Row >= 0 &&
	gridIndex.Row < len(*grid) &&
	gridIndex.Col >= 0 &&
	gridIndex.Col < len((*grid)[0])
}

func (grid *Grid) obstacleAhead(guard Guard) bool {
	indexInFront := guard.indexInFront()
	if (!grid.inBounds(indexInFront)) {
		return false
	}

	if (*grid)[indexInFront.Row][indexInFront.Col] == '#' {
		return true
	}

	return false
}

func (grid *Grid) markIndex(row int, col int, char string) {
		(*grid)[row] = (*grid)[row][:col] + char + (*grid)[row][col + 1:]
}

func (grid Grid) distinctPositions() int {
	distinctPositions := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'X' {
				distinctPositions++
			}
		}
	}

	return distinctPositions
}

func partOne(input Grid) int {
	guard := input.initGuard()

	for input.inBounds(guard.Position) {
		if input.obstacleAhead(guard) {
			guard.turnRight()
			continue
		}

		input.markIndex(guard.Position.Row, guard.Position.Col, "X")
		guard.Position = guard.indexInFront()
	}

	return input.distinctPositions()
}

type GuardImages []Guard

func (images GuardImages) looping(guard Guard) bool {
	for i := 0; i < len(images); i++ {
		if guard.Position.Row == images[i].Position.Row &&
		guard.Position.Col == images[i].Position.Col &&
		guard.Direction == images[i].Direction {
			return true
		}
	}

	return false
}

func partTwo(input Grid) int {
	originalGuard := input.initGuard()
	possibleObstructions := 0
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '#' || input[row][col] == '^' {
				continue
			}

			input.markIndex(row, col, "#")
			guard := originalGuard
			guardImages := GuardImages{}

			for {
				if guardImages.looping(guard) {
					possibleObstructions++
					break
				}

				if !input.inBounds(guard.Position) {
					break
				}

				if input.obstacleAhead(guard) {
					guard.turnRight()
					continue
				}

				imagePos := GridIndex{ Row: guard.Position.Row, Col: guard.Position.Col }
				image := Guard{ Position: imagePos, Direction: guard.Direction }
				guardImages = append(guardImages, image)

				guard.Position = guard.indexInFront()
			}

			input.markIndex(row, col, ".")
		}
	}

	return possibleObstructions
}

func main() {
	exampleOne := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	examplePartOne := partOne(append([]string{}, exampleOne...))
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := partTwo(append([]string{}, exampleOne...))
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(append([]string{}, input...))
	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))

	partTwoResult := partTwo(append([]string{}, input...))
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
