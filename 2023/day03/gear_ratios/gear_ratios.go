package gear_ratios

import "slices"

type GridRef struct {
	row int
	col int
}

var integers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func SumPartNumbers(engine_schematic [][]string) int {
	sum := 0
	for row := 0; row < len(engine_schematic); row++ {
		potential_part_number := ""
		for col := 0; col < len(engine_schematic[0]); col++ {
			current_value := engine_schematic[row][col]

			if slices.Contains(integers, current_value) {
				potential_part_number = potential_part_number + current_value
			}

			if !slices.Contains(integers, current_value) || col == len(engine_schematic[row]) {
				if potential_part_number == "" {
					continue
				}

			}
		}
	}

	return sum
}
