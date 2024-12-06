package day_5

func create_ordering_rules_map(ordering_rules_list [][2]int) map[int][]int {
	ordering_rules_map := make(map[int][]int)
	for i := 0; i < len(ordering_rules_list); i++ {
		if ordering_rules, is_in_map := ordering_rules_map[ordering_rules_list[i][0]]; is_in_map {
			ordering_rules_map[ordering_rules_list[i][0]] = append(ordering_rules, ordering_rules_list[i][1])
		} else {
			ordering_rules_map[ordering_rules_list[i][0]] = append([]int{}, ordering_rules_list[i][1])
		}
	}
	return ordering_rules_map
}

func is_update_valid(update []int, ordering_rules_map map[int][]int) bool {
	forbidden_updates := make(map[int]bool)

	for i := len(update) - 1; 0 <= i; i-- {
		if ordering_rules, is_in_map := ordering_rules_map[update[i]]; is_in_map {
			for _, v := range ordering_rules {
				forbidden_updates[v] = true
			}
		}

		if is_forbidden, is_in_map := forbidden_updates[update[i]]; is_in_map && is_forbidden {
			return false
		}
	}

	return true
}

func update_middle_number(update []int) int {
	i_center := (len(update) - 1) / 2
	return update[i_center]
}

func Sum_of_valid_updates_middle_number(updates [][]int, ordering_rules_list [][2]int) int {
	ordering_rules_map := create_ordering_rules_map(ordering_rules_list)

	middle_numbers_sum := 0

	for _, update := range updates {
		if is_update_valid(update, ordering_rules_map) {
			middle_numbers_sum += update_middle_number(update)
		}
	}

	return middle_numbers_sum
}

func correct_update(update []int, ordering_rules_map map[int][]int) []int {
	corrected_update := append([]int{}, update...)
	forbidden_updates := make(map[int]bool)

	for i := len(corrected_update) - 1; 0 <= i; {
		if ordering_rules, is_in_map := ordering_rules_map[corrected_update[i]]; is_in_map {
			for _, v := range ordering_rules {
				forbidden_updates[v] = true
			}
		}

		if is_forbidden, is_in_map := forbidden_updates[corrected_update[i]]; is_in_map && is_forbidden {
			forbidden_updates = make(map[int]bool)
			swap := corrected_update[i]
			corrected_update[i] = corrected_update[i+1]
			corrected_update[i+1] = swap
			i = len(corrected_update) - 1
		} else {
			i -= 1
		}
	}

	return corrected_update
}

func Sum_of_correct_updates_middle_number(updates [][]int, ordering_rules_list [][2]int) int {
	ordering_rules_map := create_ordering_rules_map(ordering_rules_list)

	middle_numbers_sum := 0

	for _, update := range updates {
		if !is_update_valid(update, ordering_rules_map) {
			corrected := correct_update(update, ordering_rules_map)
			middle_numbers_sum += update_middle_number(corrected)
		}
	}

	return middle_numbers_sum
}
