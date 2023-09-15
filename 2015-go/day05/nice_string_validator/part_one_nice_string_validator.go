package nice_string_validator

import (
	"slices"
	"strings"
)

func PartOneCountNiceStrings(input []string) int {
	nice_string_count := 0
	for _, nice_string_candidate := range input {
		if contains_three_vowels(nice_string_candidate) &&
			has_two_consecutive_letters(nice_string_candidate) &&
			!contains_forbidden_strings(nice_string_candidate) {
			nice_string_count++
		}
	}

	return nice_string_count
}

func contains_three_vowels(input string) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	vowel_count := 0

	for i := range input {
		if slices.Contains(vowels, input[i]) {
			vowel_count++
		}
	}

	return vowel_count >= 3
}

func has_two_consecutive_letters(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}

	return false
}

func contains_forbidden_strings(input string) bool {
	forbidden_strings := []string{"ab", "cd", "pq", "xy"}

	for _, forbidden_string := range forbidden_strings {
		if strings.Contains(input, forbidden_string) {
			return true
		}
	}

	return false
}
