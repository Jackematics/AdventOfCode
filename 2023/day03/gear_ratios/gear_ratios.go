package gear_ratios

import (
	"slices"
	"strconv"
)

type PotentialPart struct {
	row       int
	col_begin int
	col_end   int
	value     int
}

type Symbol struct {
	row   int
	col   int
	value byte
}

type GridRef struct {
	row int
	col int
}

var integers = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func SumPartNumbers(engine_schematic []string) int {

	potential_parts := findPotentialParts(engine_schematic)
	symbols := findSymbols(engine_schematic)

	parts := findValidParts(engine_schematic, potential_parts, symbols)

	sum := 0
	for _, part := range parts {
		sum += part.value
	}

	return sum
}

func SumGearRatios(engine_schematic []string) int {
	potential_parts := findPotentialParts(engine_schematic)

	symbols := findSymbols(engine_schematic)
	potential_gears := findPotentialGears(symbols)

	gear_values := findGearValues(engine_schematic, potential_parts, potential_gears)

	sum := 0
	for _, value := range gear_values {
		sum += value
	}

	return sum
}

func findPotentialParts(engine_schematic []string) []PotentialPart {
	potential_part_indexes := []PotentialPart{}

	for row := range engine_schematic {
		potential_part := ""
		for col := range engine_schematic[0] {
			current_value := engine_schematic[row][col]

			if slices.Contains(integers, current_value) {
				potential_part = potential_part + string(current_value)
			}

			if potential_part == "" {
				continue
			}

			potential_value_ended := !slices.Contains(integers, current_value) || col == len(engine_schematic[row])-1

			if potential_value_ended {
				potential_part_value, _ := strconv.Atoi(potential_part)
				potential_part_indexes = append(potential_part_indexes, PotentialPart{
					row:       row,
					col_begin: col - len(potential_part),
					col_end:   col - 1,
					value:     potential_part_value,
				})
				potential_part = ""
			}
		}
	}

	return potential_part_indexes
}

func findSymbols(engine_schematic []string) []Symbol {
	symbols := []Symbol{}
	for row := range engine_schematic {
		for col := range engine_schematic[0] {
			current_value := engine_schematic[row][col]

			if !slices.Contains(integers, current_value) && current_value != '.' {
				symbols = append(symbols, Symbol{
					row:   row,
					col:   col,
					value: current_value,
				})
			}
		}
	}

	return symbols
}

func findValidParts(
	engine_schematic []string,
	potential_parts []PotentialPart,
	symbols []Symbol) []PotentialPart {

	valid_parts := []PotentialPart{}

	for _, potential_part := range potential_parts {
		for row := potential_part.row - 1; row <= potential_part.row+1; row++ {
			for col := potential_part.col_begin - 1; col <= potential_part.col_end+1; col++ {
				for _, symbol := range symbols {
					if symbol.row == row && symbol.col == col {
						valid_parts = append(valid_parts, potential_part)
					}
				}
			}
		}
	}

	return valid_parts
}

func findPotentialGears(symbols []Symbol) []Symbol {
	potential_gears := []Symbol{}
	for _, symbol := range symbols {
		if symbol.value == '*' {
			potential_gears = append(potential_gears, symbol)
		}
	}

	return potential_gears
}

func findGearValues(
	engine_schematic []string,
	potential_parts []PotentialPart,
	potential_gears []Symbol) []int {

	gear_values := []int{}

	for _, potential_gear := range potential_gears {
		adjacent_part_values := []int{}

		forbidden_part_indexes := []int{}
		for row := potential_gear.row - 1; row <= potential_gear.row+1; row++ {
			for col := potential_gear.col - 1; col <= potential_gear.col+1; col++ {
				for index, potential_part := range potential_parts {
					if slices.Contains(forbidden_part_indexes, index) {
						continue
					}

					if potential_part.row == row && potential_part.col_begin <= col && col <= potential_part.col_end {
						adjacent_part_values = append(adjacent_part_values, potential_part.value)
						forbidden_part_indexes = append(forbidden_part_indexes, index)
						break
					}
				}
			}
		}

		if len(adjacent_part_values) == 2 {
			gear_values = append(gear_values, adjacent_part_values[0]*adjacent_part_values[1])
		}

		adjacent_part_values = adjacent_part_values[:0]
		forbidden_part_indexes = forbidden_part_indexes[:0]
	}

	return gear_values
}
