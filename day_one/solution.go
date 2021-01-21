package main

func main() {}

func solution(input []int) int {
	target := 2020

	m := make(map[int]bool)
	for _, num := range input {
		tmp := target - num

		if m[tmp] {
			return num * tmp
		} else {
			m[num] = true
		}
	}

	return 0
}
