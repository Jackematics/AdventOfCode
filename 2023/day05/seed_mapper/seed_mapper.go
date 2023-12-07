package seed_mapper

import (
	"slices"
	"strconv"
	"strings"
)

type Mapping struct {
	source_start int
	source_end   int
	dest_start   int
	dest_end     int
}

func FindLowestLocationNumber(raw_input []string, part string) int {
	raw_seeds := strings.Split(raw_input[0], ": ")[1]
	mappings := parseRawMappings(raw_input)

	var seed_results []int
	if part == "one" {
		seed_results = extractSeedsPartOne(raw_seeds)
	} else {
		seed_results = extractSeedsPartTwo(raw_seeds)
	}

	for _, mapping_set := range mappings {
		for index, seed_result := range seed_results {
			for _, mapping := range mapping_set {
				if mapping.source_start <= seed_result && seed_result <= mapping.source_end {
					source_index := seed_result - mapping.source_start
					seed_results[index] = source_index + mapping.dest_start
					break
				}
			}
		}
	}

	return slices.Min(seed_results)
}

func extractSeedsPartOne(raw_seeds string) []int {
	seeds_str := strings.Split(raw_seeds, " ")
	seeds := []int{}
	for _, seed_str := range seeds_str {
		seed, _ := strconv.Atoi(seed_str)
		seeds = append(seeds, seed)
	}

	return seeds
}

func extractSeedsPartTwo(raw_seeds string) []int {
	raw_values_arr := strings.Split(raw_seeds, " ")

	seed_starts := []int{}
	seed_ranges := []int{}

	for i, raw_value_str := range raw_values_arr {
		value, _ := strconv.Atoi(raw_value_str)
		if i%2 == 0 {
			seed_starts = append(seed_starts, value)
		} else {
			seed_ranges = append(seed_ranges, value)
		}
	}

	seeds := []int{}
	for i := range seed_starts {
		for j := seed_starts[i]; j < seed_ranges[i]+seed_starts[i]; j++ {
			seeds = append(seeds, j)
		}
	}

	return seeds
}

func parseRawMappings(raw_input []string) [][]Mapping {
	mappings := [][]Mapping{}
	current_mapping_index := -1
	for i := 2; i < len(raw_input); i++ {
		line := raw_input[i]
		if line == "" {
			continue
		}

		if strings.Contains(line, "map") {
			current_mapping_index += 1
			mappings = append(mappings, []Mapping{})
			continue
		}

		raw_mapping := strings.Split(line, " ")
		range_count, _ := strconv.Atoi(raw_mapping[2])

		source_start, _ := strconv.Atoi(raw_mapping[1])
		source_end := source_start + (range_count - 1)

		dest_start, _ := strconv.Atoi(raw_mapping[0])
		dest_end := dest_start + (range_count - 1)

		mappings[current_mapping_index] = append(mappings[current_mapping_index], Mapping{
			source_start: source_start,
			source_end:   source_end,
			dest_start:   dest_start,
			dest_end:     dest_end,
		})
	}

	return mappings
}
