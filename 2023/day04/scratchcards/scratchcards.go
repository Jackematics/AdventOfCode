package scratchcards

import (
	"slices"
	"strconv"
	"strings"
)

func SumScratchardsPoints(raw_cards []string) int {
	sum := 0
	for _, raw_card := range raw_cards {
		card_number_removed := strings.Split(raw_card, ": ")[1]
		split_drawn_and_winners := strings.Split(card_number_removed, " | ")

		numbers_drawn := formatRawNumbers(split_drawn_and_winners[0])
		winners := formatRawNumbers(split_drawn_and_winners[1])

		matches := findMatches(numbers_drawn, winners)

		sum += calculateCardValue(matches)
	}

	return sum
}

func FindTotalScratchcards(raw_cards []string) int {
	copy_counts := map[int]int{
		1: 1,
	}

	for i := 1; i <= len(raw_cards); i++ {
		_, key_exists := copy_counts[i]
		if !key_exists {
			copy_counts[i] = 1
		}

		for copy_number := 0; copy_number < copy_counts[i]; copy_number++ {
			split_by_colon := strings.Split(raw_cards[i-1], ": ")

			split_drawn_and_winners := strings.Split(split_by_colon[1], " | ")

			numbers_drawn := formatRawNumbers(split_drawn_and_winners[0])
			winners := formatRawNumbers(split_drawn_and_winners[1])
			matches := findMatches(numbers_drawn, winners)

			for match_number := 1; match_number <= len(matches); match_number++ {
				_, key_exists := copy_counts[i+match_number]
				if !key_exists {
					copy_counts[i+match_number] = 1
				}

				copy_counts[i+match_number]++
			}
		}
	}

	sum := 0
	for _, v := range copy_counts {
		sum += v
	}

	return sum
}

func formatRawNumbers(raw_numbers string) []int {
	split_by_whitespace := strings.Split(raw_numbers, " ")

	formatted := []int{}
	for _, element := range split_by_whitespace {
		if element == "" {
			continue
		}

		parsed, _ := strconv.Atoi(element)

		formatted = append(formatted, parsed)
	}

	return formatted
}

func findMatches(numbers_drawn []int, winners []int) []int {
	matches := []int{}
	for _, number := range numbers_drawn {
		if slices.Contains(winners, number) {
			matches = append(matches, number)
		}
	}

	return matches
}

func calculateCardValue(matches []int) int {
	if len(matches) == 0 {
		return 0
	}

	product := 1
	for i := 0; i < len(matches)-1; i++ {
		product *= 2
	}

	return product
}
