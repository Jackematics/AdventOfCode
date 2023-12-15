package mirage_maintenance

import (
	"strconv"
	"strings"
)

func SumNextHistoryValues(raw_input []string) int {
	sum := 0
	for _, raw_history := range raw_input {
		history := extractHistory(raw_history)
		sequences := getSequences(history)

		for i := len(sequences) - 1; i > 0; i-- {
			current_last_el := sequences[i][len(sequences[i])-1]
			next_last_el := sequences[i-1][len(sequences[i-1])-1]
			sequences[i-1] = append(sequences[i-1], current_last_el+next_last_el)
		}

		sum += sequences[0][len(sequences[0])-1]
	}

	return sum
}

func SumPrevHistoryValues(raw_input []string) int {
	sum := 0
	for _, raw_history := range raw_input {
		history := extractHistory(raw_history)
		sequences := getSequences(history)

		for i := len(sequences) - 1; i > 0; i-- {
			current_first_el := sequences[i][0]
			next_first_el := sequences[i-1][0]
			sequences[i-1] = append([]int{next_first_el - current_first_el}, sequences[i-1]...)
		}

		sum += sequences[0][0]
	}

	return sum
}

func extractHistory(raw_history string) []int {
	split_whitespace := strings.Split(raw_history, " ")
	history := []int{}

	for _, char := range split_whitespace {
		as_int, _ := strconv.Atoi(char)
		history = append(history, as_int)
	}

	return history
}

func getSequences(initial_sequence []int) [][]int {
	sequences := [][]int{initial_sequence}
	for sequenceNotAllZeroes(sequences[len(sequences)-1]) {
		sequences = append(sequences, generateNextSequence(sequences[len(sequences)-1]))
	}

	return sequences
}

func sequenceNotAllZeroes(sequence []int) bool {
	for _, element := range sequence {
		if element != 0 {
			return true
		}
	}

	return false
}

func generateNextSequence(current_sequence []int) []int {
	next_sequence := []int{}
	for i := 0; i < len(current_sequence)-1; i++ {
		next_sequence = append(next_sequence, current_sequence[i+1]-current_sequence[i])
	}

	return next_sequence
}
