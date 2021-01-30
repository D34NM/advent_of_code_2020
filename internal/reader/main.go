package reader

import (
	"bufio"
	"os"
	"strconv"
)

func main() {}

// ReadAsStrings reads the input file from the file system as slice of strings
func ReadAsStrings(file string) *[]string {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return &input
}

// ReadAsInts reads the input file from the file system as slice of ints
func ReadAsInts(file string) *[]int {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		input = append(input, i)
	}

	return &input
}
