package main

import (
	"aoc_2024/input_reader"
	"fmt"
	"sort"
	"strconv"
	"strings"
)


func abs(val int) int {
	if val < 0 {
		return -val 
	} else {
		return val
	}
}

func separateLists(input []string) ([]int, []int) {
	listOne := []int{}
	listTwo := []int{}

	for _, line := range input {
		splitList := strings.Split(line, "   ")
		listOneItem, _ := strconv.Atoi(splitList[0])
		listTwoItem, _ := strconv.Atoi(splitList[1])

		listOne = append(listOne, listOneItem)
		listTwo = append(listTwo, listTwoItem)
	}

	return listOne, listTwo
}


func partOne(input []string) int {
	listOne, listTwo := separateLists(input)
	
	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i] < listOne[j]
	})

	sort.Slice(listTwo, func(i, j int) bool {
		return listTwo[i] < listTwo[j]
	})

	sum := 0
	for i := range listOne {
		sum += int(abs(listOne[i] - listTwo[i]))
	}

	return int(sum)
}

func partTwo(input []string) int {
	listOne, listTwo := separateLists(input)

	sum := 0
	for _, val1 := range listOne {

		count := 0
		for _, val2 := range listTwo {
			if val1 == val2 {
				count++
			} 
		}

		sum += val1 * count
	}

	return sum
}

func main() {
	example := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	examplePartOne := partOne(example)
	fmt.Println("Example Result: " + strconv.Itoa(examplePartOne))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input)
	partTwoResult := partTwo(input)

	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}