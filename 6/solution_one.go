package main

func solutionOne(input *[]string) int {
	total := 0
	empty := ""

	answers := map[byte]bool{}
	groupTotal := 0
	for _, line := range *input {
		if line == empty {
			total = total + groupTotal
			groupTotal = 0
			answers = map[byte]bool{}
		}

		for i := 0; i < len(line); i++ {
			if answers[line[i]] {
				continue
			} else {
				answers[line[i]] = true
				groupTotal++
			}
		}
	}

	return total
}
