package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"slices"
	"strconv"
)

type DiskMap []int

func createDiskMap(input string) DiskMap {
	inputRunes := []rune(input)
	diskMap := DiskMap{}
	
	for i := range inputRunes  {
		intVal := int(inputRunes[i] - '0')
		
		for k := 0; k < intVal; k++ {
			if i % 2 == 0 {
				diskMap = append(diskMap, i / 2)
			}

			if i % 2 == 1 {
				diskMap = append(diskMap, -1)
			}
		}
	}

	return diskMap
}

func checksum(compressed []int) int {
	checksum := 0
	for i := range compressed {
		if compressed[i] == -1 {
			continue
		}

		checksum += i * compressed[i]
	}

	return checksum
}

func partOne(input string) int {
	diskMap := createDiskMap(input)
	compressed := append(DiskMap{}, diskMap...)

	nextFreeSpaceIndex := slices.Index(compressed, -1)
	for i := len(diskMap) - 1; i > nextFreeSpaceIndex; i-- {
		if diskMap[i] == -1 {
			compressed[i] = -1
			continue
		} 

		compressed[nextFreeSpaceIndex] = diskMap[i]
		compressed[i] = -1
		nextFreeSpaceIndex = slices.Index(compressed, -1)
	}
	
	return checksum(compressed)
}

type Range struct {
	Start int
	End int
}

func (indexRange Range) rangeDiff() int {
	return utils.Abs(indexRange.End - indexRange.Start)
}

func (diskMap DiskMap) getLastId() int {
	for i := len(diskMap) - 1; i > 0; i-- {
		if diskMap[i] != -1 {
			return diskMap[i]
		}
	} 

	panic("there must be an id")
} 

func (diskMap DiskMap) occurrenceRange(val int) Range {
	start := slices.Index(diskMap, val)
	end := start

	for i := start; i < len(diskMap) && diskMap[i] == val; i++ {
		end = i
	}

	return Range{Start: start, End: end}
}

func (diskMap DiskMap) nextAvailableEmptyRange(valRange Range) *Range {
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == -1 {
			emptyRange := DiskMap{}
			for j := i; diskMap[j] == -1 && j < valRange.Start; j++ {
				emptyRange = append(emptyRange, -1)

				if len(emptyRange) > valRange.rangeDiff() {
					return &Range{ Start: i, End: j}
				}
			}

		}
	}

	return nil
}

func partTwo(input string) int {
	diskMap := createDiskMap(input)
	compressed := append(DiskMap{}, diskMap...)

	currentId := diskMap.getLastId()
	for currentId >= 0 {
		currentIdOccurrenceRange := diskMap.occurrenceRange(currentId)
		nextAvailableEmptyRange := compressed.nextAvailableEmptyRange(currentIdOccurrenceRange)

		if nextAvailableEmptyRange == nil {
			currentId--
			continue
		}

		for i := nextAvailableEmptyRange.Start; i <= nextAvailableEmptyRange.End; i++ {
			compressed[i] = currentId
		}

		for i := currentIdOccurrenceRange.Start; i <= currentIdOccurrenceRange.End; i++ {
			compressed[i] = -1
		}

		currentId--
	}
	
	return checksum(compressed)
}


func main() {
	exampleOne := "2333133121414131402"

	examplePartOne := partOne(exampleOne)
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := partTwo(exampleOne)
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input[0])
	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))

	partTwoResult := partTwo(input[0])
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
