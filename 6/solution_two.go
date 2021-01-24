package main

func solutionTwo(input []string) int {
	empty := ""
	total := 0

	groupSize := 0
	questions := map[byte]int{}
	answers := 0
	for _, line := range input {
		if line == empty {
			for _, val := range questions {
				if val == groupSize {
					total++
				}
			}

			groupSize = 0
			questions = map[byte]int{}
			answers = 0
			continue
		}

		groupSize++

		for i := 0; i < len(line); i++ {
			val, ok := questions[line[i]]
			if ok {
				questions[line[i]] = val + 1
			} else {
				questions[line[i]] = 1
				answers++
			}
		}
	}

	return total
}
