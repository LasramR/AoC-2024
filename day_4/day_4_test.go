package day_4_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/LasramR/AoC-2024/day_4"
)

func Get_input(t *testing.T) [][]rune {
	fd, err := os.Open("day_4_input.txt")
	if err != nil {
		t.Fatalf("could not open input file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	matrix := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	return matrix
}

func TestCount_word_frequency_in_matrix(t *testing.T) {
	matrix := Get_input(t)

	t.Log("XMAS Count=", day_4.Count_word_frequency_in_matrix(matrix, []rune("XMAS")))
}

func TestCount_X_word_frequency_in_matrix(t *testing.T) {
	matrix := Get_input(t)
	t.Log("X-MAS Count=", day_4.Count_X_word_frequency_in_matrix(matrix, []rune("MAS")))
}
