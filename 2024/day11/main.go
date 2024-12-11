package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)


func partOne(input string, blinks int) int {
	rawStones := strings.Split(input, " ")
	stones := []int{}
	for _, stone := range rawStones {
		stones = append(stones, utils.ToInt(stone))
	}

	for blink := 1; blink <= blinks; blink++ {
		
		for i := len(stones) - 1; i >= 0; i-- {
			stone := stones[i]

			if stone == 0 {
				stones[i] = 1
				continue
			}

			stoneStr := strconv.Itoa(stone) 
			if len(stoneStr) % 2 == 0 {
				leftStr := stoneStr[:(len(stoneStr) / 2)]
				rightStr := stoneStr[(len(stoneStr) / 2):]

				replaced := []int{ utils.ToInt(leftStr), utils.ToInt(rightStr)}

				stones = append(stones[:i], stones[i+1:]...)
				stones = append(stones[:i], append(replaced, stones[i:]...)...)
				continue
			}

			stones[i] *= 2024
		}
	}

	return len(stones)
}


func main() {
	exampleOne := "125 17"

	exampleOneBlink := partOne(exampleOne, 1)
	fmt.Println("Example Result Part One One Blink: " + strconv.Itoa(exampleOneBlink))
	
	exampleTwoBlinks := partOne(exampleOne, 2)
	fmt.Println("Example Result Part One Two Blinks: " + strconv.Itoa(exampleTwoBlinks))

	exampleThreeBlinks := partOne(exampleOne, 3)
	fmt.Println("Example Result Part One Three Blinks: " + strconv.Itoa(exampleThreeBlinks))

	exampleFourBlinks := partOne(exampleOne, 4)
	fmt.Println("Example Result Part One Four Blinks: " + strconv.Itoa(exampleFourBlinks))

	exampleFiveBlinks := partOne(exampleOne, 5)
	fmt.Println("Example Result Part One Five Blinks: " + strconv.Itoa(exampleFiveBlinks))

	exampleSixBlinks := partOne(exampleOne, 6)
	fmt.Println("Example Result Part One Six Blinks: " + strconv.Itoa(exampleSixBlinks))

	exampleTwentyFiveBlinks := partOne(exampleOne, 25)
	fmt.Println("Example Result Part One Twenty Five Blinks: " + strconv.Itoa(exampleTwentyFiveBlinks))

	// examplePartTwo := partTwo(exampleOne)
	// fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input[0], 25)
	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))

	partTwoResult := partOne(input[0], 75)
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
