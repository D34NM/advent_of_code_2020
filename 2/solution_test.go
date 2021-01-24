package main

import (
	"bufio"
	"os"
	"testing"
)

func Test01(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

	result := solutionOne(input)

	if result != 2 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	input := []string{"1-13 a: abcde"}

	result := solutionOne(input)

	if result != 1 {
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

	if result != 439 {
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

	if result != 439 {
		t.Fail()
	}
}

func Test05(t *testing.T) {
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

	if result != 584 {
		t.Fail()
	}
}
