package main

import (
	"strconv"
	"strings"
)

func solutionOne(input []string) int {
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
