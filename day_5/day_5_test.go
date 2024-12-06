package day_5_test

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/LasramR/AoC-2024/day_5"
)

func Get_input(t *testing.T) ([][2]int, [][]int) {
	fd, err := os.Open("day_5_input.txt")
	if err != nil {
		t.Fatalf("could not open input file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	scanning_ordering_rules := true
	ordering_rules_list := [][2]int{}
	updates := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanning_ordering_rules = false
			continue
		}

		if scanning_ordering_rules {
			input_values := slices.DeleteFunc(strings.Split(strings.Trim(line, " "), "|"), func(e string) bool { return e == "" })
			a, _ := strconv.Atoi(input_values[0])
			b, _ := strconv.Atoi(input_values[1])
			ordering_rules_list = append(ordering_rules_list, [2]int{a, b})
		} else {
			input_values := slices.DeleteFunc(strings.Split(strings.Trim(line, " "), ","), func(e string) bool { return e == "" })
			update := []int{}
			for _, v := range input_values {
				a, _ := strconv.Atoi(v)
				update = append(update, a)
			}
			updates = append(updates, update)
		}
	}

	return ordering_rules_list, updates
}

func TestSum_of_valid_updates_middle_number(t *testing.T) {
	ordering_rules_list, updates := Get_input(t)
	t.Log("Sum of the middle numbers of valid updates=", day_5.Sum_of_valid_updates_middle_number(updates, ordering_rules_list))
}

func TestSum_of_correct_updates_middle_number(t *testing.T) {
	ordering_rules_list, updates := Get_input(t)
	t.Log("Sum of the middle numbers of corrected updates=", day_5.Sum_of_correct_updates_middle_number(updates, ordering_rules_list))
}
