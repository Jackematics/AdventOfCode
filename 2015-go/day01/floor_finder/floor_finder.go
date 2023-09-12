package floor_finder

import (
	"errors"
)

func FindFloor(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			count += 1
		}

		if input[i] == ')' {
			count -= 1
		}
	}

	return count
}

func FindBasementEntryPosition(input string) (int, error) {
	santa_position := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			santa_position += 1
		}

		if input[i] == ')' {
			santa_position -= 1
		}

		if santa_position == -1 {
			return i + 1, nil
		}
	}

	return -1, errors.New("Santa never enters the basement")
}
