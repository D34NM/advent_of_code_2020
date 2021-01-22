package main

import (
	"strconv"
	"strings"
)

func main() {}

func solutionPartOne(input []string) int {
	valid := 0

	for _, line := range input {
		found := 0
		fields := strings.Fields(line)
		split := strings.Index(fields[0], "-")
		min, _ := strconv.Atoi(fields[0][:split])
		max, _ := strconv.Atoi(fields[0][split+1:])
		letter := fields[1][:1]

		for i := split + 4; i < len(line); i++ {
			if line[i:i+1] == letter {
				found++
			}
		}

		if found >= min && found <= max {
			valid++
		}
	}

	return valid
}

func solutionPartTwo(input []string) int {
	valid := 0

	for _, line := range input {
		fields := strings.Fields(line)
		split := strings.Index(fields[0], "-")

		first, _ := strconv.Atoi(fields[0][:split])
		second, _ := strconv.Atoi(fields[0][split+1:])

		found := 0
		if fields[2][first-1] == fields[1][0] {
			found++
		}
		if fields[2][second-1] == fields[1][0] {
			found++
		}
		if found == 1 {
			valid++
		}

	}

	return valid
}
