package day_6_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/LasramR/AoC-2024/day_6"
)

func Get_input(t *testing.T) [][]rune {
	fd, err := os.Open("day_6_input.txt")
	if err != nil {
		t.Fatalf("could not open input file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	district_layout := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		district_layout = append(district_layout, []rune(line))
	}

	return district_layout
}

func TestCount_word_frequency_in_matrix(t *testing.T) {
	district_layout := Get_input(t)

	t.Log("Guard step count=", day_6.Count_visited_district_position(district_layout))
}

func TestCount_number_of_possible_loop_in_district(t *testing.T) {
	district_layout := Get_input(t)

	t.Log("Number of possible loop in district=", day_6.Count_number_of_possible_loop_in_district(district_layout))
}
