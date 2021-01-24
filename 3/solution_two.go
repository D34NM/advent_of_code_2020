package main

func solutionTwo(input []string, right, down int) int {
	trees := 0
	tree := byte('#')

	length := len(input[0])
	index := right

	for i := down; i < len(input); i = i + down {
		row := input[i]
		tmp := index % length
		current := row[tmp]
		if current == tree {
			trees++
		}

		index = index + right
	}

	return trees
}
