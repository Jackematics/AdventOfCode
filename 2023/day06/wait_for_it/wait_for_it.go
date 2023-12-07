package wait_for_it

import (
	"strconv"
	"strings"
)

func MultiplyWinningStrategyCount(raw_input []string) int {
	times := extractValuesPartOne(raw_input[0])
	record_distances := extractValuesPartOne(raw_input[1])

	products := calculateWinningStrategyCount(times, record_distances)

	sum := 1
	for _, product := range products {
		sum *= product
	}

	return sum
}

func TotalWinningStrategyCount(raw_input []string) int {
	time := extractValuePartTwo(raw_input[0])
	record_distance := extractValuePartTwo(raw_input[1])

	win_sum := 0
	for j := 0; j <= time; j++ {
		if j*(time-j) > record_distance {
			win_sum++
		}
	}

	return win_sum
}

func extractValuesPartOne(raw_str string) []int {
	values := []int{}
	split_values_by_whitespace := strings.Split(strings.Split(raw_str, ":")[1], " ")
	for _, str := range split_values_by_whitespace {
		if str == "" {
			continue
		}

		value, _ := strconv.Atoi(str)
		values = append(values, value)
	}

	return values
}

func extractValuePartTwo(raw_str string) int {
	split_values_by_whitespace := strings.Split(strings.Split(raw_str, ":")[1], " ")
	value_str := ""
	for _, str := range split_values_by_whitespace {
		if str == "" {
			continue
		}

		value_str += str
	}

	value, _ := strconv.Atoi(value_str)

	return value
}

func calculateWinningStrategyCount(times []int, record_distances []int) []int {
	products := []int{}
	for i := range times {
		win_sum := 0
		for j := 0; j <= times[i]; j++ {
			if j*(times[i]-j) > record_distances[i] {
				win_sum++
			}
		}

		products = append(products, win_sum)
	}

	return products
}
