package pipe_maze

import (
	"aoc_2023/day10/pipe_maze/pipe"
	"slices"
)

func FindFarthestDistanceFromLoopStart(raw_input []string) int {
	pipe_maze := parseInput(raw_input)
	start_index := getStartIndex(raw_input)

	start := pipe_maze[start_index.Row][start_index.Col]
	current_direction := start.ValidEntryDirections()[0]

	current_index_exit_data := start.Enter(current_direction)

	steps := 1
	for !indexesEqual(current_index_exit_data.Index, start_index) {
		current_pipe := pipe_maze[current_index_exit_data.Index.Row][current_index_exit_data.Index.Col]
		current_index_exit_data = current_pipe.Enter(current_index_exit_data.Direction)
		steps++
	}

	if steps%2 == 0 {
		return steps / 2
	} else {
		return (steps + 1) / 2
	}
}

func indexesEqual(indexA pipe.GridIndex, indexB pipe.GridIndex) bool {
	return indexA.Row == indexB.Row && indexA.Col == indexB.Col
}

func parseInput(raw_input []string) [][]pipe.Pipe {
	pipe_maze := [][]pipe.Pipe{}
	for row := range raw_input {
		pipe_maze = append(pipe_maze, []pipe.Pipe{})

		for col := range raw_input[0] {
			pipe_maze[row] = append(pipe_maze[row], setPipe(raw_input, row, col))
		}
	}

	return pipe_maze
}

func setPipe(raw_input []string, row, col int) pipe.Pipe {
	raw_pipe := raw_input[row][col]
	gridIndex := pipe.GridIndex{Row: row, Col: col}

	if raw_pipe == '|' {
		return pipe.VerticalPipe{GridIndex: gridIndex}
	}

	if raw_pipe == '-' {
		return pipe.HorizontalPipe{GridIndex: gridIndex}
	}

	if raw_pipe == 'L' {
		return pipe.NorthAndEastBend{GridIndex: gridIndex}
	}

	if raw_pipe == 'J' {
		return pipe.NorthAndWestBend{GridIndex: gridIndex}
	}

	if raw_pipe == '7' {
		return pipe.SouthAndWestBend{GridIndex: gridIndex}
	}

	if raw_pipe == 'F' {
		return pipe.SouthAndEastBend{GridIndex: gridIndex}
	}

	if raw_pipe == 'S' {
		pipe_type := determinePipe(raw_input, row, col)
		return pipe.StartingPipe{GridIndex: gridIndex, Pipe: pipe_type}
	}

	return pipe.Ground{GridIndex: gridIndex}
}

func determinePipe(raw_input []string, row, col int) pipe.Pipe {
	valid_directions := getValidDirections(raw_input, row, col)
	gridIndex := pipe.GridIndex{Row: row, Col: col}

	if slices.Contains(valid_directions, pipe.North) && slices.Contains(valid_directions, pipe.South) {
		return pipe.VerticalPipe{GridIndex: gridIndex}
	}

	if slices.Contains(valid_directions, pipe.East) && slices.Contains(valid_directions, pipe.West) {
		return pipe.HorizontalPipe{GridIndex: gridIndex}
	}

	if slices.Contains(valid_directions, pipe.North) && slices.Contains(valid_directions, pipe.East) {
		return pipe.NorthAndEastBend{GridIndex: gridIndex}
	}

	if slices.Contains(valid_directions, pipe.North) && slices.Contains(valid_directions, pipe.West) {
		return pipe.NorthAndWestBend{GridIndex: gridIndex}
	}

	if slices.Contains(valid_directions, pipe.South) && slices.Contains(valid_directions, pipe.West) {
		return pipe.SouthAndWestBend{GridIndex: gridIndex}
	}

	if slices.Contains(valid_directions, pipe.South) && slices.Contains(valid_directions, pipe.East) {
		return pipe.SouthAndEastBend{GridIndex: gridIndex}
	}

	return pipe.Ground{GridIndex: gridIndex}
}

func getValidDirections(raw_input []string, row, col int) []pipe.Direction {
	valid_north := []rune{'|', '7', 'F'}
	valid_east := []rune{'-', 'J', '7'}
	valid_south := []rune{'|', 'L', 'J'}
	valid_west := []rune{'-', 'L', 'F'}

	valid_directions := []pipe.Direction{}
	if slices.Contains(valid_north, rune(raw_input[row-1][col])) {
		valid_directions = append(valid_directions, pipe.North)
	}

	if slices.Contains(valid_east, rune(raw_input[row][col+1])) {
		valid_directions = append(valid_directions, pipe.East)
	}

	if slices.Contains(valid_south, rune(raw_input[row+1][col])) {
		valid_directions = append(valid_directions, pipe.South)
	}

	if slices.Contains(valid_west, rune(raw_input[row][col-1])) {
		valid_directions = append(valid_directions, pipe.West)
	}

	return valid_directions
}

func getStartIndex(raw_input []string) pipe.GridIndex {
	for row := range raw_input {
		for col := range raw_input[0] {
			if raw_input[row][col] == 'S' {
				return pipe.GridIndex{Row: row, Col: col}
			}
		}
	}

	return pipe.GridIndex{Row: -1, Col: -1}
}
