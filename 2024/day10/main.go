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

type GridIndexes []GridIndex

type Trail struct {
	VisitedNines []GridIndex
	Path []GridIndex
}

type TopographicalMap [][]int

func initTopographicalMap(input []string) TopographicalMap {
	topographicalMap := TopographicalMap{}
	for i := range input {
		row := []int{}
		for _, line := range []rune(input[i]) {
			row = append(row, int(line - '0'))
		}
		topographicalMap = append(topographicalMap, row)
	}

	return topographicalMap
}

func (topographicalMap TopographicalMap) trailheadIndexes() []GridIndex {
	trailheadIndexes := []GridIndex{}
	for row := range topographicalMap {
		for  col := range topographicalMap[row] {
			if topographicalMap[row][col] == 0 {
				trailheadIndexes = append(trailheadIndexes, GridIndex{Row: row, Col: col})
			}
		}
	}

	return trailheadIndexes
}

func (topographicalMap TopographicalMap) inBounds(index GridIndex) bool {
	return index.Row >= 0 &&
	index.Row < len(topographicalMap) && 
	index.Col >= 0 &&
	index.Col < len(topographicalMap[0])
}

func contains(indexes []GridIndex, index GridIndex) bool {
	for _, currentIndex := range indexes {
		if index.Row == currentIndex.Row && index.Col == currentIndex.Col {
			return true
		}
	}

	return false
}

func (topographicalMap TopographicalMap) getViablePathsPartOne(viablePaths []GridIndex) []GridIndex {
	deltas := []GridIndex{
		{ Row: 1, Col: 0 },
		{ Row: -1, Col: 0 },
		{ Row: 0, Col: 1 },
		{ Row: 0, Col: -1 },
	}

	nextViablePaths := []GridIndex{}
	for _, path := range viablePaths {
		for _, delta := range deltas {
			adjacentIndex := GridIndex{Row: path.Row + delta.Row, Col: path.Col + delta.Col}
	
			if !topographicalMap.inBounds(adjacentIndex)  {
				continue
			}
	
			if topographicalMap[adjacentIndex.Row][adjacentIndex.Col] == topographicalMap[path.Row][path.Col] + 1 {
				nextViablePaths = append(nextViablePaths, adjacentIndex)
			}
		}
	}

	return nextViablePaths
}

func (topographicalMap TopographicalMap) getVisitedNines(visitedNines []GridIndex, viablePaths []GridIndex) []GridIndex {
	for len(viablePaths) != 0 {
		for i := len(viablePaths) - 1; i >= 0; i-- {
			path := viablePaths[i]
			if contains(visitedNines, path) {
				viablePaths = append(viablePaths[:i], viablePaths[i+1:]...)
				continue
			}

			if topographicalMap[path.Row][path.Col] == 9 {
				visitedNines = append(visitedNines, path)
				viablePaths = append(viablePaths[:i], viablePaths[i+1:]...)
				continue
			}
		}
		
		viablePaths = topographicalMap.getViablePathsPartOne(viablePaths)
	}

	return visitedNines
}

func partOne(input []string) int {
	topographicalMap := initTopographicalMap(input)
	trailheadIndexes := topographicalMap.trailheadIndexes()
	
	sum := 0
	for _, trailheadIndex := range trailheadIndexes {
		sum += len(topographicalMap.getVisitedNines([]GridIndex{}, []GridIndex{trailheadIndex}))
	}

	return sum
}

func (topographicalMap TopographicalMap) getViablePathsPartTwo(trailTail GridIndex, visitedNines GridIndexes) []GridIndex {
	viablePaths := []GridIndex{}
	deltas := []GridIndex{
		{ Row: 1, Col: 0 },
		{ Row: -1, Col: 0 },
		{ Row: 0, Col: 1 },
		{ Row: 0, Col: -1 },
	}

	for _, delta := range deltas {
		adjacentIndex := GridIndex{Row: trailTail.Row + delta.Row, Col: trailTail.Col + delta.Col}

		if !topographicalMap.inBounds(adjacentIndex) || contains(visitedNines, adjacentIndex)  {
			continue
		}

		if topographicalMap[adjacentIndex.Row][adjacentIndex.Col] == topographicalMap[trailTail.Row][trailTail.Col] + 1 {
			viablePaths = append(viablePaths, adjacentIndex)
		}
	}

	return viablePaths
}

func (topographicalMap TopographicalMap) getAllTrails(trails [][]GridIndex, trail []GridIndex, visitedNines []GridIndex) [][]GridIndex {
	currentTail := trail[len(trail) - 1]
	if topographicalMap[currentTail.Row][currentTail.Col] == 9 {
		visitedNines = append(visitedNines, currentTail)
	}
	viableTails := topographicalMap.getViablePathsPartTwo(currentTail, visitedNines)

	if len(viableTails) == 0 {
		return append(trails, trail)
	}

	for _, tail := range viableTails {
		trails = topographicalMap.getAllTrails(trails, append(trail, tail), visitedNines)
	}

	return trails
}

func (topographicalMap TopographicalMap) partTwoScore(trailheadIndex GridIndex) int {
	trails := topographicalMap.getAllTrails([][]GridIndex{}, []GridIndex{trailheadIndex}, []GridIndex{})
	score := 0
	for _, trail := range trails {
		if len(trail) == 10 {
			score++
		}
	}

	return score
}

func partTwo(input []string) int {
	topographicalMap := initTopographicalMap(input)
	trailheadIndexes := topographicalMap.trailheadIndexes()
	
	sum := 0
	for _, trailheadIndex := range trailheadIndexes {
		sum += topographicalMap.partTwoScore(trailheadIndex)
	}

	return sum
}

func main() {
	exampleOne := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	examplePartOne := partOne(exampleOne)
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := partTwo(exampleOne)
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input)
	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))

	partTwoResult := partTwo(input)
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
