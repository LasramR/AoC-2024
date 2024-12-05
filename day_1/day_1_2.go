package day_1

import (
	"errors"
)

func list_frequency_map(l []int) map[int]int {
	frequency_map := make(map[int]int)

	for i := 0; i < len(l); i++ {
		frequency, is_in_map := frequency_map[l[i]]

		if is_in_map {
			frequency_map[l[i]] = frequency + 1
		} else {
			frequency_map[l[i]] = 1
		}
	}

	return frequency_map
}

func Total_similarity(l1, l2 []int) (int, error) {
	if len(l1) != len(l2) {
		return -1, errors.New("Total_similarity: input lists must have the same length")
	}

	l2_frenquency_map := list_frequency_map(l2)

	similarity := 0
	for i := 0; i < len(l1); i++ {
		if frequency_in_l2, is_in_l2 := l2_frenquency_map[l1[i]]; is_in_l2 {
			similarity += l1[i] * frequency_in_l2
		}
	}

	return similarity, nil
}
