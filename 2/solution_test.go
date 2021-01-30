package main

import (
	"testing"

	"github.com/DeanMilojevic/advent_of_code_2020/internal/reader"
)

func Test01(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

	result := solutionOne(&input)

	if result != 2 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	input := []string{"1-13 a: abcde"}

	result := solutionOne(&input)

	if result != 1 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	result := solutionOne(input)

	if result != 439 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	result := solutionOne(input)

	if result != 439 {
		t.Fail()
	}
}

func Test05(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	result := solutionTwo(input)

	if result != 584 {
		t.Fail()
	}
}
