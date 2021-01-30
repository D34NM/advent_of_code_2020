package main

import (
	"testing"

	"github.com/DeanMilojevic/advent_of_code_2020/internal/reader"
)

func Test01(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	result := solutionOne(&input)

	if result != 514579 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	input := reader.ReadAsInts("input.txt")

	result := solutionOne(input)

	if result != 898299 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	input := reader.ReadAsInts("input.txt")
	result := solutionTwo(input)

	if result != 143933922 {
		t.Fail()
	}
}
