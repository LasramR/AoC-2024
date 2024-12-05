package day_1

import (
	"errors"
	"sort"
)

func sort_list(l []int) []int {
	sorted := append([]int{}, l...)
	sort.Ints(sorted)
	return sorted
}

func min_max(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func Total_distance(l1, l2 []int) (int, error) {
	if len(l1) != len(l2) {
		return -1, errors.New("total_distance: input lists must have the same length")
	}

	item_count := len(l1)
	sorted_l1 := sort_list(l1)
	sorted_l2 := sort_list(l2)

	distance := 0
	for i := 0; i < item_count; i++ {
		min, max := min_max(sorted_l1[i], sorted_l2[i])
		distance += max - min
	}

	return distance, nil
}
