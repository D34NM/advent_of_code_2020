package main

import (
	"sort"
)

func solutionTwo(input *[]int) int {
	sort.Ints(*input)
	length := len(*input)
	for i := 0; i < length-2; i++ {
		j, k := i+1, length-1
		for j < k {
			cur := (*input)[i] + (*input)[j] + (*input)[k]
			if cur > 2020 {
				k--
			} else if cur < 2020 {
				j++
			} else {
				return (*input)[i] * (*input)[j] * (*input)[k]
			}
		}
	}

	return 0
}
