package present_wrapping_calculator

import (
	"slices"
	"strconv"
	"strings"
)

func CalculateAllRibbonNeeded(raw_dimensions_list []string) int {
	ribbon_count_ft := 0
	for i := range raw_dimensions_list {
		ribbon_count_ft += calculate_ribbon_needed(raw_dimensions_list[i])
	}
	return ribbon_count_ft
}

func calculate_ribbon_needed(raw_dimensions string) int {
	dimensions := get_dimensions(raw_dimensions)
	slices.Sort(dimensions)

	smallest_perimeter := 2*dimensions[0] + 2*dimensions[1]

	bow_length := dimensions[0] * dimensions[1] * dimensions[2]

	return smallest_perimeter + bow_length
}

func CalculateAllPaperNeeded(raw_dimensions_list []string) int {
	paper_count_ft := 0
	for i := range raw_dimensions_list {
		paper_count_ft += calculate_paper_needed(raw_dimensions_list[i])
	}

	return paper_count_ft
}

func calculate_paper_needed(raw_dimensions string) int {
	dimensions := get_dimensions(raw_dimensions)

	areas := []int{
		dimensions[0] * dimensions[1],
		dimensions[1] * dimensions[2],
		dimensions[0] * dimensions[2],
	}

	min := get_min_area(areas)

	return (2*areas[0] + 2*areas[1] + 2*areas[2]) + min
}

func get_dimensions(raw_dimensions string) []int {
	raw_dimensions_split := strings.Split(raw_dimensions, "x")

	var dimensions []int
	for i := range raw_dimensions_split {
		int_value, _ := strconv.Atoi(raw_dimensions_split[i])
		dimensions = append(dimensions, int_value)
	}

	return dimensions
}

func get_min_area(areas []int) int {
	min := areas[0]
	for i := range areas {
		if areas[i] < min {
			min = areas[i]
		}
	}

	return min
}
