package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type OrderingRule struct {
	Before int
	After int
}

func (rule OrderingRule) valid(a int, aIndex int, b int, bIndex int) bool {
	if rule.Before == a {
		return rule.After != b || aIndex < bIndex 
	}	

	if rule.Before == b {
		return rule.After != a || bIndex < aIndex
	}

	return true
}

type Update []int

func (update Update) correct(rules []OrderingRule) bool {
	correct := true
	for i := range update {
		for j := i + 1; j < len(update); j++ {
			for _, rule := range rules {
				if !rule.valid(update[i], i, update[j], j) {
					correct = false
				}
			}
		}
	}

	return correct
}

func (update *Update) fix(rules []OrderingRule) {
	for !update.correct(rules) {
		for i := range *update {
			for j := i + 1; j < len(*update); j++ {
				for _, rule := range rules {
					if !rule.valid((*update)[i], i, (*update)[j], j) {
						tempI := (*update)[i]
						(*update)[i] = (*update)[j]
						(*update)[j] = tempI
					}
				}
			}
		}
	}
}


func partOneAndTwo(input []string) (int, int) {
	orderingRules := []OrderingRule{}
	updates := []Update{}

	for _, line := range input {
		if strings.Contains(line, "|") {
			ruleSplit := strings.Split(line, "|")

			orderingRules = append(orderingRules, OrderingRule{Before: utils.ToInt(ruleSplit[0]), After: utils.ToInt(ruleSplit[1])})
		}

		if strings.Contains(line, ",") {
			pageNumberSplit := strings.Split(line, ",")

			update := []int{} 
			for _, pageNumber := range pageNumberSplit {
				update = append(update, utils.ToInt(pageNumber))
			}

			updates = append(updates, update)
		}
	}

	partOneSolution := 0
	partTwoSolution := 0
	for _, update := range updates {
		if update.correct(orderingRules) {
			partOneSolution += update[utils.MiddleIndex(update)]
		} else {
			update.fix(orderingRules)
			partTwoSolution += update[utils.MiddleIndex(update)]
		}
	}

	return partOneSolution, partTwoSolution
}

func main() {
	exampleOne := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	examplePartOne, examplePartTwo := partOneAndTwo(exampleOne)
	fmt.Println("Example Result Part One: " + strconv.Itoa(examplePartOne))
	fmt.Println("Example Result Part Two: " + strconv.Itoa(examplePartTwo))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult, partTwoResult := partOneAndTwo(input)

	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
