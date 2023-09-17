package part_2_light_configurer

import (
	"strconv"
	"strings"
)

func CalculateTotalBrightness(input []string) int {
	light_grid := initialise_light_grid(1000, 1000)

	for _, instruction := range input {
		instruction_split := strings.Split(instruction, " ")

		operation := determine_light_operation(instruction_split)
		apply_instruction_to_grid(instruction_split, operation, light_grid)

	}

	return count_lights_on(light_grid)
}

func initialise_light_grid(numRows int, numCols int) [][]*Light {
	light_grid := make([][]*Light, 1000)

	for row := 0; row < numRows; row++ {
		light_grid[row] = make([]*Light, numCols)
		for col := 0; col < numCols; col++ {
			light_grid[row][col] = NewLight()
		}
	}

	return light_grid
}

func determine_light_operation(instruction_split []string) string {
	var operation = "toggle"
	if instruction_split[1] == "on" {
		operation = "turn_on"
	}
	if instruction_split[1] == "off" {
		operation = "turn_off"
	}

	return operation
}

func apply_instruction_to_grid(instruction_split []string, operation string, light_grid [][]*Light) {
	initialRowCol := strings.Split(instruction_split[len(instruction_split)-3], ",")
	endRowCol := strings.Split(instruction_split[len(instruction_split)-1], ",")

	rowRangeStart, _ := strconv.Atoi(initialRowCol[0])
	rowRangeEnd, _ := strconv.Atoi(endRowCol[0])
	colRangeStart, _ := strconv.Atoi(initialRowCol[1])
	colRangeEnd, _ := strconv.Atoi(endRowCol[1])

	for row := rowRangeStart; row <= rowRangeEnd; row++ {
		for col := colRangeStart; col <= colRangeEnd; col++ {
			if operation == "toggle" {
				light_grid[row][col].Toggle()
			}
			if operation == "turn_on" {
				light_grid[row][col].TurnOn()
			}
			if operation == "turn_off" {
				light_grid[row][col].TurnOff()
			}
		}
	}
}

func count_lights_on(light_grid [][]*Light) int {
	total_brightness := 0
	for row := range light_grid {
		for col := range light_grid[0] {
			total_brightness += light_grid[row][col].GetBrightness()
		}
	}

	return total_brightness
}
