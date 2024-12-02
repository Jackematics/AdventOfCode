package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func acceptableIncreases(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		difference := utils.Abs(report[i] - report[i+1])
		if report[i] < report[i+1] || difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func acceptableDecreases(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		difference := utils.Abs(report[i] - report[i+1])
		if report[i] > report[i+1] || difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func partOne(input []string) int {
	safeSum := 0
	for _, rawLine := range input {
		splitLine := strings.Split(rawLine, " ")

		report := []int{}
		for i := range splitLine {
			level, _ := strconv.Atoi(splitLine[i])
			report = append(report, level)
		}

		if acceptableIncreases(report) || acceptableDecreases(report) {
			safeSum += 1
		}
	}

	return safeSum
}

func partTwo(input []string) int {
	safeSum := 0
	for _, rawLine := range input {
		splitLine := strings.Split(rawLine, " ")

		report := []int{}
		for i := range splitLine {
			level, _ := strconv.Atoi(splitLine[i])
			report = append(report, level)
		}

		if !(acceptableIncreases(report) || acceptableDecreases(report)) {
			for j := range report {
				faultRemoved := make([]int, 0, len(report)-1)

				faultRemoved = append(faultRemoved, report[:j]...)
				faultRemoved = append(faultRemoved, report[j+1:]...)

				if acceptableIncreases(faultRemoved) || acceptableDecreases(faultRemoved) {
					safeSum++
					break
				}
			}
		} else {
			safeSum++
		}
	}

	return safeSum
}

func main() {
	example := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	// examplePartOne := partOne(example)
	// fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := partTwo(example)
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input)
	partTwoResult := partTwo(input)

	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
