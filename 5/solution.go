package main

func main() {}

func solution(input []string) int {
	highestID := 0
	front := byte('F')
	back := byte('B')
	left := byte('L')
	right := byte('R')

	row := 0
	column := 0

	for _, line := range input {
		lowR, highR := 0, 127
		lowC, highC := 0, 7
		for i := 0; i < len(line); i++ {
			if line[i] == front {
				highR = (highR + lowR) / 2
				row = highR
			}
			if line[i] == back {
				lowR = ((highR + lowR) / 2) + 1
				row = lowR
			}

			if line[i] == left {
				highC = (highC + lowC) / 2
				column = highC
			}
			if line[i] == right {
				lowC = ((highC + lowC) / 2) + 1
				column = lowC
			}
		}

		tmp := column + (row * 8)

		if tmp > highestID {
			highestID = tmp
		}

		column = 0
		row = 0
	}

	return highestID
}
