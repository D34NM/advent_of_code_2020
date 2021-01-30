package main

import (
	"testing"

	"github.com/DeanMilojevic/advent_of_code_2020/internal/reader"
)

func Test01(t *testing.T) {
	input := []string{
		"abcx",
		"abcy",
		"abcz",
		"",
	}

	result := solutionOne(&input)

	if result != 6 {
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
		"",
	}

	result := solutionOne(&input)

	if result != 11 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	*input = append(*input, "")

	result := solutionOne(input)

	if result != 6947 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
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
		"",
	}

	result := solutionTwo(&input)

	if result != 6 {
		t.Fail()
	}
}

func Test05(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	*input = append(*input, "")

	result := solutionTwo(input)

	if result != 3398 {
		t.Fail()
	}
}
