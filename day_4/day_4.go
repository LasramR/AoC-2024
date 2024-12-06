package day_4

// -1 if outside of matrix
func flattened_matrix_all_direction_relative_index(i_center, j_center, distance, row_len, col_len int) [8]int {
	indexes := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}

	// Right
	if j_center+distance < row_len {
		indexes[0] = i_center*row_len + j_center + distance
	}
	// Left
	if 0 <= j_center-distance {
		indexes[1] = i_center*row_len + j_center - distance
	}
	// Down
	if i_center+distance < col_len {
		indexes[2] = (i_center+distance)*row_len + j_center
	}
	// Up
	if 0 <= i_center-distance {
		indexes[3] = (i_center-distance)*row_len + j_center
	}
	// Down Right
	if j_center+distance < row_len && i_center+distance < col_len {
		indexes[4] = (i_center+distance)*row_len + j_center + distance
	}
	// Up Right
	if j_center+distance < row_len && 0 <= i_center-distance {
		indexes[5] = (i_center-distance)*row_len + j_center + distance
	}
	// Down Left
	if 0 <= j_center-distance && i_center+distance < col_len {
		indexes[6] = (i_center+distance)*row_len + j_center - distance
	}
	// Up Left
	if 0 <= j_center-distance && 0 <= i_center-distance {
		indexes[7] = (i_center-distance)*row_len + j_center - distance
	}

	return indexes
}

func flattened_matrix_index(i, j, row_len int) int {
	return i*row_len + j
}
func check_all_directions_in_flattened_matrix(flattened_matrix []rune, word []rune, row_len, col_len, i_center, j_center int) int {
	direction_is_valid := [8]bool{true, true, true, true, true, true, true, true}

	if flattened_matrix[flattened_matrix_index(i_center, j_center, row_len)] != word[0] {
		return 0
	}

	for i := 1; i < len(word); i++ {
		still_match := false

		for direction_index, next_rune_index := range flattened_matrix_all_direction_relative_index(i_center, j_center, i, row_len, col_len) {
			if next_rune_index == -1 {
				direction_is_valid[direction_index] = false
			}
			if direction_is_valid[direction_index] && flattened_matrix[next_rune_index] == word[i] {
				still_match = true
			} else {
				direction_is_valid[direction_index] = false
			}
		}

		if !still_match {
			return 0
		}
	}

	count := 0

	for _, direction_valid := range direction_is_valid {
		if direction_valid {
			count += 1
		}
	}

	return count
}

func flattened_matrix[T any](matrix [][]T) ([]T, int, int) {
	col_len := len(matrix)
	row_len := len(matrix[0])

	flattened := []T{}
	for _, row := range matrix {
		flattened = append(flattened, row...)
	}

	return flattened, row_len, col_len
}

func Count_word_frequency_in_matrix(matrix [][]rune, word []rune) int {
	count := 0

	flattened, row_len, col_len := flattened_matrix(matrix)

	for i := 0; i < col_len; i++ {
		for j := 0; j < row_len; j++ {
			count += check_all_directions_in_flattened_matrix(flattened, word, row_len, col_len, i, j)
		}
	}

	return count
}
