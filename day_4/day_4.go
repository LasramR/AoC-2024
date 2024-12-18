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
	// Up Left
	if 0 <= j_center-distance && 0 <= i_center-distance {
		indexes[5] = (i_center-distance)*row_len + j_center - distance
	}
	// Up Right
	if j_center+distance < row_len && 0 <= i_center-distance {
		indexes[6] = (i_center-distance)*row_len + j_center + distance
	}
	// Down Left
	if 0 <= j_center-distance && i_center+distance < col_len {
		indexes[7] = (i_center+distance)*row_len + j_center - distance
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

// We will not check for cases when len(word)%2 == 0 :)
func check_X_in_flattened_matrix(flattened_matrix []rune, word []rune, row_len, col_len, i_center, j_center int) int {
	word_i_center := (len(word) - 1) / 2

	direction_is_valid := [8]bool{true, true, true, true, true, true, true, true}
	word_direction := [4]int{0, 0, 0, 0}

	if flattened_matrix[flattened_matrix_index(i_center, j_center, row_len)] != word[word_i_center] {
		return 0
	}

	for i := 1; i < word_i_center+1; i++ {
		still_match := false

		direction_indexes := flattened_matrix_all_direction_relative_index(i_center, j_center, i, row_len, col_len)

		for di := 0; di < 4; di++ {
			if direction_is_valid[di*2] && direction_is_valid[di*2+1] {
				if flattened_matrix[direction_indexes[di*2]] == word[word_i_center-i] && flattened_matrix[direction_indexes[di*2+1]] == word[word_i_center+i] && word_direction[di] <= 0 {
					word_direction[di] = -1
					still_match = true
				} else if flattened_matrix[direction_indexes[di*2+1]] == word[word_i_center-i] && flattened_matrix[direction_indexes[di*2]] == word[word_i_center+i] && word_direction[di] >= 0 {
					word_direction[di] = 1
					still_match = true
				} else {
					direction_is_valid[di*2] = false
					direction_is_valid[di*2+1] = false
				}
			}
		}

		if !still_match {
			return 0
		}
	}

	if direction_is_valid[4] && direction_is_valid[5] && direction_is_valid[6] && direction_is_valid[7] {
		return 1
	}

	return 0
}

func Count_X_word_frequency_in_matrix(matrix [][]rune, word []rune) int {
	count := 0
	flattened, row_len, col_len := flattened_matrix(matrix)
	word_i_center := (len(word) - 1) / 2
	for i := word_i_center; i < col_len-word_i_center; i++ {
		for j := word_i_center; j < row_len-word_i_center; j++ {
			count += check_X_in_flattened_matrix(flattened, word, row_len, col_len, i, j)
		}
	}

	return count
}
