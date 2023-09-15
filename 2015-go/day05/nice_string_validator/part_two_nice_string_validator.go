package nice_string_validator

func PartTwoCountNiceStrings(input []string) int {
	nice_string_count := 0
	for _, nice_string_candidate := range input {
		if has_repeating_non_overlapping_pairs(nice_string_candidate) &&
			contains_a_letter_sandwich(nice_string_candidate) {
			nice_string_count++
		}
	}

	return nice_string_count
}

func has_repeating_non_overlapping_pairs(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		pair := input[i : i+2]
		pair_repeat_count := 0

		for j := 0; j < len(input)-1; j++ {
			if input[j:j+2] == pair {
				pair_repeat_count++
				j += 1
			}
		}
		if pair_repeat_count >= 2 {
			return true
		}
	}

	return false
}

func contains_a_letter_sandwich(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		sequence := input[i : i+3]
		if sequence[0] == sequence[2] {
			return true
		}
	}

	return false
}
