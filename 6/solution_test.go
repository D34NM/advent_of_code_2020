package main

import (
	"bufio"
	"os"
	"testing"
)

func Test01(t *testing.T) {
	input := []string{
		"abcx",
		"abcy",
		"abcz",
	}

	result := solutionOne(input)

	if result != 0 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	result := solutionOne(input)

	if result != 0 {
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

	result := solutionOne(input)

	if result != 0 {
		t.Fail()
	}
}
