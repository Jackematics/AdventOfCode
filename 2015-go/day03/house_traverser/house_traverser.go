package house_traverser

type GridRef struct {
	row int
	col int
}

func TraverseHouses(move_orders string) int {
	houses := [][]int{{2}}

	santa_position := GridRef{row: 0, col: 0}
	robo_santa_position := GridRef{row: 0, col: 0}

	for i := range move_orders {

		move_order := move_orders[i]
		if i%2 == 0 {
			adjust_entity_on_house_expansion(move_order, &santa_position, &robo_santa_position)
			santa_position, houses = move_santa(move_order, santa_position, houses)
		}

		if i%2 == 1 {
			adjust_entity_on_house_expansion(move_order, &robo_santa_position, &santa_position)
			robo_santa_position, houses = move_santa(move_order, robo_santa_position, houses)
		}
	}
	return count_houses_with_presents(houses)
}

func adjust_entity_on_house_expansion(move_order byte, entity_a *GridRef, entity_b *GridRef) {
	if move_order == '^' && entity_a.row == 0 {
		entity_b.row += 1
	}
	if move_order == '<' && entity_a.col == 0 {
		entity_b.col += 1
	}
}

func move_santa(move_order byte, entity_position GridRef, houses [][]int) (GridRef, [][]int) {
	switch move_order {
	case '^':
		entity_position.row -= 1

		if entity_position.row == -1 {
			houses = expand_row_up(houses)
			entity_position.row += 1
		}

	case '>':
		entity_position.col += 1
		if entity_position.col == len(houses[0]) {
			houses = expand_col_right(houses)
		}

	case 'v':
		entity_position.row += 1

		if entity_position.row == len(houses) {
			houses = expand_row_down(houses)
		}

	case '<':
		entity_position.col -= 1

		if entity_position.col == -1 {
			houses = expand_col_left(houses)

			entity_position.col += 1
		}

	}
	houses[entity_position.row][entity_position.col] += 1

	return entity_position, houses
}

func count_houses_with_presents(houses [][]int) int {
	at_least_one_present_count := 0
	for row := range houses {
		for col := range houses[row] {
			if houses[row][col] > 0 {
				at_least_one_present_count += 1
			}
		}
	}

	return at_least_one_present_count
}

func expand_row_up(houses [][]int) [][]int {
	newRow := make([]int, len(houses[0]))
	return append([][]int{newRow}, houses...)
}

func expand_col_right(houses [][]int) [][]int {
	for i := range houses {
		houses[i] = append(houses[i], 0)
	}

	return houses
}

func expand_row_down(houses [][]int) [][]int {
	newRow := make([]int, len(houses[0]))
	return append(houses, newRow)
}

func expand_col_left(houses [][]int) [][]int {
	for i := range houses {
		houses[i] = append([]int{0}, houses[i]...)
	}

	return houses
}
