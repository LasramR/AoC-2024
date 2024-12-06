package day_6

const (
	VISITED_DISTRICT_CELL   rune = 'X'
	UNVISITED_DISTRICT_CELL rune = '.'
	OBSTRUCTED_CELL         rune = '#'
	GUARD_UP                rune = '^'
	GUARD_DOWN              rune = 'v'
	GUARD_RIGHT             rune = '>'
	GUARD_LEFT              rune = '<'
)

func get_guard_position(district_layout [][]rune) (int, int, rune) {
	for i := 0; i < len(district_layout); i++ {
		for j := 0; j < len(district_layout[i]); j++ {
			switch district_layout[i][j] {
			case GUARD_UP:
				return i, j, GUARD_UP
			case GUARD_DOWN:
				return i, j, GUARD_DOWN
			case GUARD_RIGHT:
				return i, j, GUARD_RIGHT
			case GUARD_LEFT:
				return i, j, GUARD_LEFT
			}
		}
	}
	return -1, -1, -1
}

func next_guard_move(district_layout [][]rune, guard_i, guard_j int, guard_facing rune) (int, int, rune) {
	if guard_facing == GUARD_UP {
		if district_layout[guard_i-1][guard_j] == OBSTRUCTED_CELL {
			return next_guard_move(district_layout, guard_i, guard_j, GUARD_RIGHT)
		}
		return guard_i - 1, guard_j, GUARD_UP
	}

	if guard_facing == GUARD_DOWN {
		if district_layout[guard_i+1][guard_j] == OBSTRUCTED_CELL {
			return next_guard_move(district_layout, guard_i, guard_j, GUARD_LEFT)
		}
		return guard_i + 1, guard_j, GUARD_DOWN
	}

	if guard_facing == GUARD_RIGHT {
		if district_layout[guard_i][guard_j+1] == OBSTRUCTED_CELL {
			return next_guard_move(district_layout, guard_i, guard_j, GUARD_DOWN)
		}
		return guard_i, guard_j + 1, GUARD_RIGHT
	}

	if district_layout[guard_i][guard_j-1] == OBSTRUCTED_CELL {
		return next_guard_move(district_layout, guard_i, guard_j, GUARD_UP)
	}
	return guard_i, guard_j - 1, GUARD_LEFT
}

func is_guard_looping(district_layout [][]rune, guard_i, guard_j int, guard_facing rune) bool {
	return district_layout[guard_i][guard_j] == guard_facing
}

func visit_next_district_cell(district_layout [][]rune, guard_i, guard_j int, guard_facing rune, only_visit bool) (int, int, rune, bool) {
	next_guard_i, next_guard_j, next_guard_facing := next_guard_move(district_layout, guard_i, guard_j, guard_facing)

	already_visited := false
	if only_visit {
		already_visited = district_layout[next_guard_i][next_guard_j] == VISITED_DISTRICT_CELL
		district_layout[guard_i][guard_j] = VISITED_DISTRICT_CELL
	} else {
		already_visited = is_guard_looping(district_layout, next_guard_i, next_guard_j, next_guard_facing)
	}

	district_layout[next_guard_i][next_guard_j] = next_guard_facing

	return next_guard_i, next_guard_j, next_guard_facing, already_visited
}

func is_guard_leaving_next_step(district_layout [][]rune, guard_i, guard_j int, guard_facing rune) bool {
	if guard_facing == GUARD_UP && guard_i == 0 {
		return true
	}
	if guard_facing == GUARD_DOWN && guard_i == len(district_layout)-1 {
		return true
	}
	if guard_facing == GUARD_LEFT && guard_j == 0 {
		return true
	}
	if guard_facing == GUARD_RIGHT && guard_j == len(district_layout[0])-1 {
		return true
	}
	return false
}

func Count_visited_district_position(district_layout [][]rune) int {
	guard_i, guard_j, guard_facing := get_guard_position(district_layout)
	already_visited := false
	step_count := 0

	for !is_guard_leaving_next_step(district_layout, guard_i, guard_j, guard_facing) {
		guard_i, guard_j, guard_facing, already_visited = visit_next_district_cell(district_layout, guard_i, guard_j, guard_facing, true)
		if !already_visited {
			step_count += 1
		}
	}
	district_layout[guard_i][guard_j] = VISITED_DISTRICT_CELL

	return step_count + 1
}

func copy_district(district_layout [][]rune) [][]rune {
	district_copy := make([][]rune, len(district_layout))
	for i := range district_layout {
		district_copy[i] = make([]rune, len(district_layout[i]))
		copy(district_copy[i], district_layout[i])
	}
	return district_copy
}

func Count_number_of_possible_loop_in_district(district_layout [][]rune) int {
	possible_loop_count := 0

	initial_guard_i, initial_guard_j, initial_guard_facing := get_guard_position(district_layout)

	total_steps := Count_visited_district_position(copy_district(district_layout))
	prev_guard_facing := initial_guard_facing

	for prev_guard_facing == initial_guard_facing {
		initial_guard_i, initial_guard_j, initial_guard_facing, _ = visit_next_district_cell(district_layout, initial_guard_i, initial_guard_j, initial_guard_facing, false)
		total_steps -= 1
	}

	for i := 0; i < total_steps; i++ {
		district_copy := copy_district(district_layout)

		new_obstruction_i, new_obstruction_j, _ := next_guard_move(district_copy, initial_guard_i, initial_guard_j, initial_guard_facing)
		district_copy[new_obstruction_i][new_obstruction_j] = OBSTRUCTED_CELL

		guard_i, guard_j, guard_facing := initial_guard_i, initial_guard_j, initial_guard_facing
		is_looping := false
		for !is_guard_leaving_next_step(district_copy, guard_i, guard_j, guard_facing) {

			guard_i, guard_j, guard_facing, is_looping = visit_next_district_cell(district_copy, guard_i, guard_j, guard_facing, false)
			if is_looping {
				possible_loop_count += 1
				break
			}
		}

		initial_guard_i, initial_guard_j, initial_guard_facing, _ = visit_next_district_cell(district_layout, initial_guard_i, initial_guard_j, initial_guard_facing, false)
	}

	return possible_loop_count
}
