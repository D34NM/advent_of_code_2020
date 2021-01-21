package main

import (
	"strconv"
)

func main() {}

func solution(input []string) int {
	valid := 0

	for _, line := range input {
		found := 0
		first, fEnd := parseNumber(line, 0, byte('-'))
		second, sEnd := parseNumber(line, fEnd+1, byte(' '))
		letter := line[sEnd+1]

		for i := sEnd + 4; i < len(line); i++ {
			if line[i] == letter {
				found++
			}
		}

		if found >= first && found <= second {
			valid++
		}
	}

	return valid
}

func parseNumber(input string, fromIndex int, character byte) (int, int) {
	endIndex := fromIndex
	for i := fromIndex; i < len(input); i++ {
		if input[i] == character {
			break
		}

		endIndex++
	}

	num, _ := strconv.Atoi(input[fromIndex:endIndex])
	return num, endIndex
}
