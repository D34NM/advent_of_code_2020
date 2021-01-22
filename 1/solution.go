package main

import (
	"sort"
)

func main() {}

func solutionPartOne(input []int) int {
	target := 2020

	m := make(map[int]bool)
	for _, num := range input {
		tmp := target - num

		if m[tmp] {
			return num * tmp
		}

		m[num] = true
	}

	return 0
}

func solutionPartTwo(input []int) int {
	sort.Ints(input)
	length := len(input)
	for i := 0; i < length-2; i++ {
		j, k := i+1, length-1
		for j < k {
			cur := input[i] + input[j] + input[k]
			if cur > 2020 {
				k--
			} else if cur < 2020 {
				j++
			} else {
				return input[i] * input[j] * input[k]
			}
		}
	}

	return 0
}
