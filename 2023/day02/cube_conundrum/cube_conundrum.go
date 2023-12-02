package cube_conundrum

import (
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

var max_cube_set = CubeSet{
	red:   12,
	green: 13,
	blue:  14,
}

func SumValidIds(raw_records []string) int {
	sum := 0
	for _, raw_record := range raw_records {
		game_split := strings.Split(raw_record, ":")
		game_id, _ := strconv.Atoi(strings.Split(game_split[0], " ")[1])
		raw_set_split := strings.Split(game_split[1], ";")
		games_valid := gamesValid(raw_set_split)

		if games_valid {
			sum += game_id
		}
	}

	return sum
}

func SumSetPower(raw_records []string) int {
	sum := 0
	for _, raw_record := range raw_records {
		game_split := strings.Split(raw_record, ":")
		raw_set_split := strings.Split(game_split[1], ";")

		sets := []CubeSet{}
		for _, raw_set := range raw_set_split {
			sets = append(sets, getCubeSet(raw_set))
		}

		max_colours := getMaxColours(sets)

		sum += max_colours.red * max_colours.green * max_colours.blue
	}

	return sum
}

func gamesValid(raw_set_split []string) bool {
	validation_results := []bool{}
	for _, raw_set := range raw_set_split {
		set := getCubeSet(raw_set)

		red_valid := set.red <= max_cube_set.red
		green_valid := set.green <= max_cube_set.green
		blue_valid := set.blue <= max_cube_set.blue

		validation_results = append(validation_results, red_valid && blue_valid && green_valid)
	}

	for _, valid_result := range validation_results {
		if !valid_result {
			return false
		}
	}

	return true
}

func getCubeSet(raw_set string) CubeSet {
	red_value := getValue(raw_set, "red")
	green_value := getValue(raw_set, "green")
	blue_value := getValue(raw_set, "blue")

	return CubeSet{
		red:   red_value,
		green: green_value,
		blue:  blue_value,
	}
}

func getValue(raw_set string, colour string) int {
	colour_pairs := strings.Split(raw_set, ",")

	for _, colour_pair := range colour_pairs {
		if strings.Contains(colour_pair, colour) {
			split_by_whitespace := strings.Split(colour_pair, " ")
			colour_value_str := split_by_whitespace[1]
			value, _ := strconv.Atoi(colour_value_str)

			return value
		}
	}

	return 0
}

func getMaxColours(sets []CubeSet) CubeSet {
	max_red, max_green, max_blue := sets[0].red, sets[0].green, sets[0].blue

	for _, set := range sets {
		if max_red < set.red {
			max_red = set.red
		}

		if max_green < set.green {
			max_green = set.green
		}

		if max_blue < set.blue {
			max_blue = set.blue
		}
	}

	return CubeSet{red: max_red, green: max_green, blue: max_blue}
}
