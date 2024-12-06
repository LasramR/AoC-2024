package day_3_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/LasramR/AoC-2024/day_3"
)

func Get_input(t *testing.T) string {
	fd, err := os.Open("day_3_input.txt")
	if err != nil {
		t.Fatalf("could not open input file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	memory := ""
	for scanner.Scan() {
		line := scanner.Text()
		memory += line
	}

	return memory
}

func TestEvaluate_sequence(t *testing.T) {
	memory := Get_input(t)

	t.Log("memory instructions result=", day_3.Evaluate_sequence(memory))
}
