package day_1_test

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/LasramR/AoC-2024/day_1"
)

func Get_input(t *testing.T) ([]int, []int) {
	fd, err := os.Open("day_1_input.txt")
	if err != nil {
		t.Fatalf("could not open input file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	l1 := []int{}
	l2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		input_values := slices.DeleteFunc(strings.Split(strings.Trim(line, " "), " "), func(e string) bool { return e == "" })
		l1_value, l1_err := strconv.Atoi(input_values[0])
		l2_value, l2_err := strconv.Atoi(input_values[1])

		if l1_err != nil || l2_err != nil {
			t.Fatalf("could not parse input")
		}
		l1 = append(l1, l1_value)
		l2 = append(l2, l2_value)
	}

	return l1, l2
}

func TestTotal_Distance(t *testing.T) {
	l1, l2 := Get_input(t)

	distance, err := day_1.Total_distance(l1, l2)

	if err != nil {
		t.Fatalf("Total_distance should not have failed")
	}

	t.Log("Total_distance=", distance)
}

func TestTotal_Similarity(t *testing.T) {
	l1, l2 := Get_input(t)

	similarity, err := day_1.Total_similarity(l1, l2)

	if err != nil {
		t.Fatalf("Total_similarity should not have failed")
	}

	t.Log("Total_similarity=", similarity)
}
