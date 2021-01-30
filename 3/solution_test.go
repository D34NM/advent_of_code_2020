package main

import (
	"testing"

	"github.com/DeanMilojevic/advent_of_code_2020/internal/reader"
)

func Test01(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	result := solutionOne(&input)

	if result != 7 {
		t.Fail()
	}
}

func Test02(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	result := solutionOne(input)

	if result != 193 {
		t.Fail()
	}
}

func Test03(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	first := solutionTwo(&input, 1, 1)
	second := solutionTwo(&input, 3, 1)
	third := solutionTwo(&input, 5, 1)
	forth := solutionTwo(&input, 7, 1)
	fifth := solutionTwo(&input, 1, 2)

	result := first * second * third * forth * fifth

	if result != 336 {
		t.Fail()
	}
}

func Test04(t *testing.T) {
	input := reader.ReadAsStrings("input.txt")

	first := solutionTwo(input, 1, 1)
	second := solutionTwo(input, 3, 1)
	third := solutionTwo(input, 5, 1)
	forth := solutionTwo(input, 7, 1)
	fifth := solutionTwo(input, 1, 2)

	result := first * second * third * forth * fifth

	if result != 1355323200 {
		t.Fail()
	}
}
