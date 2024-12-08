package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	Value int
	Operators []int
}

var OperationsPartOne = map[rune]func(int, int) int{
	'+': func(a int, b int) int { return a + b },
	'*': func(a int, b int) int { return a * b },
}

var OperationsPartTwo = map[rune]func(int, int) int{
	'+': func(a int, b int) int { return a + b },
	'*': func(a int, b int) int { return a * b },
	'|': func(a int, b int) int { return utils.ToInt(strconv.Itoa(a) + strconv.Itoa(b)) },
}

func generateOperationPermutationsPartOne(permutations []string, permutation string, n int) []string {
	if len(permutation) == n {
		return append(permutations, permutation)
	}

	permutations = generateOperationPermutationsPartOne(permutations, permutation + "+", n)
	return generateOperationPermutationsPartOne(permutations, permutation + "*", n)
}

func generateOperationPermutationsPartTwo(permutations []string, permutation string, n int) []string {
	if len(permutation) == n {
		return append(permutations, permutation)
	}

	permutations = generateOperationPermutationsPartTwo(permutations, permutation + "+", n)
	permutations = generateOperationPermutationsPartTwo(permutations, permutation + "*", n)
	return generateOperationPermutationsPartTwo(permutations, permutation + "|", n)
}

func (equation Equation) partOneCalibratable() bool {
	operationPermutations := generateOperationPermutationsPartOne([]string{}, "", len(equation.Operators) - 1)

	for _, permutation := range operationPermutations {
		currentResult := append([]int{}, equation.Operators...)
		for _, operation := range permutation {
			currentResult = append([]int{OperationsPartOne[operation](currentResult[0], currentResult[1])}, currentResult[2:]...)  
		}

		if currentResult[0] == equation.Value {
			return true
		}
	}

	return false
}

func (equation Equation) partTwoCalibratable() bool {
	operationPermutations := generateOperationPermutationsPartTwo([]string{}, "", len(equation.Operators) - 1)

	for _, permutation := range operationPermutations {
		currentResult := append([]int{}, equation.Operators...)
		for _, operation := range permutation {
			currentResult = append([]int{OperationsPartTwo[operation](currentResult[0], currentResult[1])}, currentResult[2:]...)  
		}

		if currentResult[0] == equation.Value {
			return true
		}
	}

	return false
}

func partOne(input []string) int {
	equations := []Equation{}
	for _, line := range input {
		lineSplit := strings.Split(line, ": ")

		operatorsSplit := strings.Split(lineSplit[1], " ")
		operators := []int{}
		for _, rawOperator := range operatorsSplit {
			operators = append(operators, utils.ToInt(rawOperator))
		}

		equations = append(equations, Equation{ Value: utils.ToInt(lineSplit[0]), Operators: operators })
	}

	sum := 0
	for _, equation := range equations {
		if equation.partOneCalibratable() {
			sum += equation.Value
		}
	}

	return sum
}

func partTwo(input []string) int {
	equations := []Equation{}
	for _, line := range input {
		lineSplit := strings.Split(line, ": ")

		operatorsSplit := strings.Split(lineSplit[1], " ")
		operators := []int{}
		for _, rawOperator := range operatorsSplit {
			operators = append(operators, utils.ToInt(rawOperator))
		}

		equations = append(equations, Equation{ Value: utils.ToInt(lineSplit[0]), Operators: operators })
	}

	sum := 0
	for _, equation := range equations {
		if equation.partTwoCalibratable() {
			sum += equation.Value
		}
	}

	return sum
}

func main() {
	exampleOne := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
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
