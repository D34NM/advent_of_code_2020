package main

import (
	"testing"

	"github.com/DeanMilojevic/advent_of_code_2020/internal/reader"
)

func Test01(t *testing.T) {
	input := []string{
		"FBFBBFFRLR",
	}

	result := solutionOne(&input)

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

	result := solutionOne(&input)

	if result != 820 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	result := solutionOne(input)

	if result != 906 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	result := solutionTwo(input)

	if result != 519 {
		t.Fail()
	}
}
