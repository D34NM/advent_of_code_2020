package main

import (
	"strconv"
	"strings"
)

func solutionOne(input *[]string) int {
	accumulator := 0
	seen := map[int]bool{}
	length := len(*input)
	for i := 0; i < length; {
		instruction := (*input)[i]
		if seen[i] {
			break
		} else {
			seen[i] = true
		}

		instructionParts := strings.Fields(instruction)

		switch instructionParts[0] {
		case "acc":
			n, _ := strconv.Atoi(instructionParts[1])
			accumulator = accumulator + n
			i++
			break
		case "jmp":
			n, _ := strconv.Atoi(instructionParts[1])
			i = (i + n) % length
			break
		default:
			i++
			break
		}
	}

	return accumulator
}
