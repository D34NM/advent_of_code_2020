package reader

import (
	"bufio"
	"os"
)

func main() {}

// ReadInputFile reads the input file from the file system
func ReadInputFile(file string) *[]string {
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
