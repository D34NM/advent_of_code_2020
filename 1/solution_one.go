package main

func solutionOne(input *[]int) int {
	target := 2020

	m := make(map[int]bool)
	for _, num := range *input {
		tmp := target - num

		if m[tmp] {
			return num * tmp
		}

		m[num] = true
	}

	return 0
}
