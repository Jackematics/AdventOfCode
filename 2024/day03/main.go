package main

import (
	"aoc_2024/input_reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


func partOne(input string) int {
	mulRegexp := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	muls := mulRegexp.FindAllString(input, -1)

	sum := 0
	mulValRegex := regexp.MustCompile("[0-9]{1,3}")
	for _, mul := range muls {
		mulVals := mulValRegex.FindAllString(mul, -1)
		val1, _ := strconv.Atoi(mulVals[0])
		val2, _ := strconv.Atoi(mulVals[1])

		sum += (val1 * val2)
	}


	return sum 
}

func partTwo(input string) int {
	commandRegExp := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don\'t\(\)`)

	commands := commandRegExp.FindAllString(input, -1)

	sum := 0
	mulValRegex := regexp.MustCompile("[0-9]{1,3}")
	enabled := true
	for _, command := range commands {
		if command == "don't()" {
			enabled = false
		}

		if command == "do()" {
			enabled = true
		}

		if strings.Contains(command, "mul") && enabled {
			mulVals := mulValRegex.FindAllString(command, -1)
			val1, _ := strconv.Atoi(mulVals[0])
			val2, _ := strconv.Atoi(mulVals[1])
			
			sum += (val1 * val2)
		}
	}


	return sum 
}

func main() {
	exampleOne := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	exampleTwo := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	examplePartOne := partOne(exampleOne)
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))

	examplePartTwo := partTwo(exampleTwo)
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input[0])
	partTwoResult := partTwo(input[0])

	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
