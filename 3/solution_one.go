package main

func solutionOne(input *[]string) int {
	trees := 0
	tree := byte('#')

	length := len((*input)[0])
	index := 3

	for i := 1; i < len(*input); i++ {
		row := (*input)[i]
		tmp := index % length
		current := row[tmp]
		if current == tree {
			trees++
		}

		index = index + 3
	}

	return trees
}
