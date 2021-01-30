package main

import (
	"strconv"
	"strings"
)

func solutionTwo(input *[]string) int {
	valid := 0

	for _, line := range *input {
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
