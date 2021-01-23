package main

import (
	"bufio"
	"os"
	"testing"
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

	result := solution(input, 3, 1)

	if result != 7 {
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

	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	result := solution(input, 3, 1)

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

	first := solution(input, 1, 1)
	second := solution(input, 3, 1)
	third := solution(input, 5, 1)
	forth := solution(input, 7, 1)
	fifth := solution(input, 1, 2)

	result := first * second * third * forth * fifth

	if result != 336 {
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

	first := solution(input, 1, 1)
	second := solution(input, 3, 1)
	third := solution(input, 5, 1)
	forth := solution(input, 7, 1)
	fifth := solution(input, 1, 2)

	result := first * second * third * forth * fifth

	if result != 1355323200 {
		t.Fail()
	}
}
