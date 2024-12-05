package day_2_test

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/LasramR/AoC-2024/day_2"
)

func Get_input(t *testing.T) [][]int {
	fd, err := os.Open("day_2_input.txt")
	if err != nil {
		t.Fatalf("could not open input file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	reports := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		input_values := slices.DeleteFunc(strings.Split(strings.Trim(line, " "), " "), func(e string) bool { return e == "" })

		report := []int{}

		for _, raw_value := range input_values {
			level, err := strconv.Atoi(raw_value)

			if err != nil {
				t.Fatalf("could not parse input")
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports
}

func TestCount_safe_reports(t *testing.T) {
	reports := Get_input(t)

	t.Log("safe record count=", day_2.Count_safe_reports(reports))
}
