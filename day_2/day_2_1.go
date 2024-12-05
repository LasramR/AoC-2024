package day_2

func min_max(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// It wasn't specified, but I assume that each report contains unique level values
func is_report_safe(report []int, is_increasing bool, min_gap int, max_gap int) bool {
	prev_level := report[0]

	for i_level := 1; i_level < len(report); i_level++ {
		if (is_increasing && report[i_level] < prev_level) || (!is_increasing && prev_level < report[i_level]) {
			return false
		}

		min, max := min_max(prev_level, report[i_level])
		gap := max - min

		if gap < min_gap || max_gap < gap {
			return false
		}

		prev_level = report[i_level]
	}

	return true
}

func Count_safe_reports(reports [][]int) int {
	safe_reports_count := 0

	for _, report := range reports {
		if len(report) <= 1 || is_report_safe(report, report[0] < report[1], 1, 3) {
			safe_reports_count += 1
		}
	}

	return safe_reports_count
}
