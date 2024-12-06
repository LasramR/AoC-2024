package day_2

func min_max(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func is_report_safe(report []int, min_gap int, max_gap int, max_unsafe_levels int) bool {
	check_without_level := func(i_level_to_ignore int) bool {
		report_without_level := append([]int{}, report[:i_level_to_ignore]...)
		report_without_level = append(report_without_level, report[i_level_to_ignore+1:]...)
		return is_report_safe(report_without_level, min_gap, max_gap, max_unsafe_levels-1)
	}

	if report[0] == report[1] {
		return 0 < max_unsafe_levels && check_without_level(0)
	}

	i_prev_level := 0
	initial_curve := report[0] < report[1]
	for i_level := 1; i_level < len(report); i_level++ {
		current_curve := report[i_prev_level] < report[i_level]
		if initial_curve != current_curve {
			return 0 < max_unsafe_levels && (check_without_level(i_prev_level) || check_without_level(i_level))
		}

		min, max := min_max(report[i_prev_level], report[i_level])
		gap := max - min
		if gap < min_gap || max_gap < gap {
			return 0 < max_unsafe_levels && (check_without_level(i_prev_level) || check_without_level(i_level))
		}

		i_prev_level = i_level
	}

	return true
}

func Count_safe_reports(reports [][]int, max_unsafe_levels int) int {
	safe_reports_count := 0

	for _, report := range reports {
		if is_report_safe(report, 1, 3, max_unsafe_levels) {
			safe_reports_count += 1
		}
	}

	return safe_reports_count
}
