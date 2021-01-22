package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func Test01(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	result := solutionPartOne(input)

	if result != 514579 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	f, err := os.Open("input.txt")

	if err != nil {
		t.Fail()
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		input = append(input, i)
	}

	result := solutionPartOne(input)

	if result != 898299 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	f, err := os.Open("input.txt")

	if err != nil {
		t.Fail()
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		input = append(input, i)
	}

	result := solutionPartTwo(input)

	if result != 143933922 {
		t.Fail()
	}
}
