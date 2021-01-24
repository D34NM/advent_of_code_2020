package main

import (
	"bufio"
	"os"
	"testing"
)

func Test01(t *testing.T) {
	input := []string{
		"FBFBBFFRLR",
	}

	result := solutionOne(input)

	if result != 357 {
		t.Fail()
	}
}

func Test02One(t *testing.T) {
	input := []string{
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}

	result := solutionOne(input)

	if result != 820 {
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

	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	result := solutionOne(input)

	if result != 906 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	f, err := os.Open("input.txt")

	if err != nil {
		t.Fail()
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	result := solutionTwo(input)

	if result != 519 {
		t.Fail()
	}
}
